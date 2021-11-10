package fetch_wechat_payment

import (
	"fmt"
	"github.com/eden-framework/sqlx"
	"github.com/eden-w2w/lib-modules/constants/enums"
	"github.com/eden-w2w/lib-modules/constants/general_errors"
	"github.com/eden-w2w/lib-modules/databases"
	"github.com/eden-w2w/lib-modules/modules/goods"
	"github.com/eden-w2w/lib-modules/modules/order"
	"github.com/eden-w2w/lib-modules/modules/payment_flow"
	"github.com/eden-w2w/lib-modules/modules/wechat"
	"github.com/eden-w2w/srv-cmop/internal/global"
	"github.com/eden-w2w/wechatpay-go/core"
	"github.com/eden-w2w/wechatpay-go/services/payments/jsapi"
	"github.com/sirupsen/logrus"
)

func TaskFetchWechatPaymentStatus() {
	// 查询所有未达终态的支付单
	flows, err := payment_flow.GetController().GetFlowByOrderIDAndStatus(
		0, 0,
		[]enums.PaymentStatus{
			enums.PAYMENT_STATUS__CREATED,
			enums.PAYMENT_STATUS__PROCESS,
		}, nil,
	)
	if err != nil {
		logrus.Errorf(
			"[TaskFetchWechatPaymentStatus] payment_flow.GetController().GetFlowByOrderIDAndStatus err: %v",
			err,
		)
		return
	}

	for _, flow := range flows {
		req := jsapi.QueryOrderByOutTradeNoRequest{
			OutTradeNo: core.String(fmt.Sprintf("%d", flow.FlowID)),
			Mchid:      core.String(global.Config.Wechat.MerchantID),
		}
		tran, err := wechat.GetController().QueryOrderByOutTradeNo(req)
		if err != nil {
			continue
		}
		tradeState, err := enums.ParseWechatTradeStateFromString(*tran.TradeState)
		if err != nil {
			logrus.Warningf(
				"[TaskFetchWechatPaymentStatus] enums.ParseWechatTradeStateFromString err: %v, TradeState: %s",
				err,
				*tran.TradeState,
			)
			continue
		}

		tx := sqlx.NewTasks(global.Config.MasterDB)
		var paymentFlow *databases.PaymentFlow
		tx = tx.With(
			func(db sqlx.DBExecutor) error {
				paymentFlow, err = payment_flow.GetController().GetPaymentFlowByID(flow.FlowID, db, true)
				if err != nil {
					return err
				}

				if !tradeState.IsFail() {
					amount := uint64(*tran.Amount.Total)
					if paymentFlow.Amount != amount {
						return general_errors.FlowAmountIncorrect
					}
				}
				return nil
			},
		)

		tx = tx.With(
			func(db sqlx.DBExecutor) (err error) {
				if tradeState.IsEqual(paymentFlow.Status) {
					return nil
				}
				if tradeState.IsSuccess() {
					// 检查代金券信息
					var discountAmount, actualAmount = uint64(0), uint64(0)
					if len(tran.PromotionDetail) > 0 {
						for _, detail := range tran.PromotionDetail {
							discountAmount += uint64(*detail.Amount)
						}
						actualAmount = paymentFlow.Amount - discountAmount
						err = payment_flow.GetController().UpdatePaymentFlowAmount(
							paymentFlow.FlowID,
							discountAmount,
							actualAmount,
							db,
						)
						if err != nil {
							return err
						}
					}

					err = payment_flow.GetController().UpdatePaymentFlowStatus(
						paymentFlow,
						enums.PAYMENT_STATUS__SUCCESS,
						tran,
						db,
					)
					if err != nil {
						return
					}
					// 联动订单
					var orderModel *databases.Order
					var logistics *databases.OrderLogistics
					orderModel, logistics, err = order.GetController().GetOrder(
						paymentFlow.OrderID,
						paymentFlow.UserID,
						db,
						true,
					)
					if err != nil {
						return err
					}
					orderGoods, err := order.GetController().GetOrderGoods(paymentFlow.OrderID, db)
					if err != nil {
						return err
					}
					updateParams := order.UpdateOrderParams{
						Status: enums.ORDER_STATUS__PAID,
					}
					if len(tran.PromotionDetail) > 0 {
						updateParams.DiscountAmount = discountAmount
					}
					err = order.GetController().UpdateOrder(
						orderModel,
						logistics,
						orderGoods,
						updateParams,
						goods.GetController().LockInventory,
						goods.GetController().UnlockInventory,
						db,
					)
				} else if tradeState.IsFail() {
					err = payment_flow.GetController().UpdatePaymentFlowStatus(
						paymentFlow,
						tradeState.ToPaymentStatus(),
						tran,
						db,
					)
				}
				return payment_flow.GetController().UpdatePaymentFlowRemoteID(
					paymentFlow.FlowID,
					*tran.TransactionId,
					db,
				)
			},
		)

		err = tx.Do()
		if err != nil {
			logrus.Errorf("[TaskFetchWechatPaymentStatus] tx.Do() err: %v, tran: %+v", err, tran)
		}
	}
}
