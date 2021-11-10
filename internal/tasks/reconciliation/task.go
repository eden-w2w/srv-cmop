package reconciliation

import (
	"fmt"
	"github.com/eden-framework/sqlx/datatypes"
	"github.com/eden-w2w/lib-modules/constants/enums"
	"github.com/eden-w2w/lib-modules/modules/payment_flow"
	"github.com/eden-w2w/lib-modules/modules/task_flow"
	"github.com/eden-w2w/lib-modules/modules/wechat"
	"github.com/eden-w2w/wechatpay-go/core"
	"github.com/eden-w2w/wechatpay-go/services/payments/jsapi"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

func TaskReconciliation() {
	currentTime := datatypes.MySQLTimestamp(time.Now())
	logrus.Infof("[TaskReconciliation] start at %s", currentTime.String())
	task, _ := task_flow.GetController().CreateTaskFlow(
		task_flow.CreateTaskFlowParams{
			Name: currentTime.Format("2006-01-02") + "对账任务",
			Type: enums.TASK_TYPE__RECONCILIATION,
		}, nil,
	)
	var errs = make([]string, 0)
	defer func() {
		if task != nil {
			if len(errs) > 0 {
				_ = task_flow.GetController().UpdateTaskFlow(
					task.FlowID, task_flow.UpdateTaskParams{
						EndedAt: datatypes.MySQLTimestamp(time.Now()),
						Status:  enums.TASK_PROCESS_STATUS__FAIL,
						Message: strings.Join(errs, ", "),
					},
				)
			} else {
				_ = task_flow.GetController().UpdateTaskFlow(
					task.FlowID, task_flow.UpdateTaskParams{
						EndedAt: datatypes.MySQLTimestamp(time.Now()),
						Status:  enums.TASK_PROCESS_STATUS__COMPLETE,
					},
				)
			}
		}
		logrus.Infof("[TaskReconciliation] complete at %s", datatypes.MySQLTimestamp(time.Now()).String())
	}()
	lastDate := datatypes.MySQLTimestamp(time.Time(currentTime).AddDate(0, 0, -1))
	bill, err := wechat.GetController().GetTradeBill(
		jsapi.TradeBillRequest{
			BillDate: core.String(lastDate.Format("2006-01-02")),
		},
	)
	if err != nil {
		errs = append(errs, err.Error())
		return
	}

	data, err := wechat.GetController().DownloadURL(*bill)
	if err != nil {
		errs = append(errs, err.Error())
		return
	}

	parser := NewTradeBillParser(data)
	err = parser.Iterator(checker)
	if err != nil {
		errs = append(errs, err.Error())
	}
	errs = append(errs, parser.Errors())
}

var refundList = make(map[uint64]*TradeBill, 0)

func checker(bill *TradeBill, last bool) (err error) {
	if last {
		if len(refundList) > 0 {
			errList := make([]uint64, 0)
			for flowID := range refundList {
				errList = append(errList, flowID)
			}
			err = fmt.Errorf("存在交易单已退款，微信支付单未退款的记录[%v]", errList)
		}
		return
	}
	flow, err := payment_flow.GetController().GetPaymentFlowByID(bill.PaymentFlowID, nil, false)
	if err != nil {
		return
	}

	if flow.Status == enums.PAYMENT_STATUS__SUCCESS {
		if bill.Status != enums.WECHAT_TRADE_STATE__SUCCESS {
			return fmt.Errorf(
				"交易单状态不一致，交易单ID[%d]，交易单状态[%s]，微信支付单状态[%s]",
				flow.FlowID,
				flow.Status.String(),
				bill.Status.String(),
			)
		}
		if flow.Amount != uint64(bill.OrderPrice) {
			return fmt.Errorf(
				"交易单金额不一致，交易单ID[%d]，交易单金额[%d]，微信支付单金额[%d]",
				flow.FlowID,
				flow.Amount,
				bill.OrderPrice,
			)
		}
	} else if flow.Status == enums.PAYMENT_STATUS__REFUND {
		if bill.Status == enums.WECHAT_TRADE_STATE__SUCCESS {
			if _, ok := refundList[flow.FlowID]; ok {
				return fmt.Errorf("交易单退款状态下，不能存在两笔支付成功的微信支付单，交易单ID[%d]", flow.FlowID)
			}
			refundList[flow.FlowID] = bill
		} else if bill.Status == enums.WECHAT_TRADE_STATE__REFUND {
			if _, ok := refundList[flow.FlowID]; !ok {
				return fmt.Errorf("交易单退款状态下，未找到支付成功的微信支付单，交易单ID[%d]", flow.FlowID)
			}
			delete(refundList, flow.FlowID)
			if flow.Amount != uint64(bill.ActualRefundPrice) {
				return fmt.Errorf(
					"退款金额不一致，交易单ID[%d]，交易单金额[%d]，微信支付单退款金额[%d]",
					flow.FlowID,
					flow.Amount,
					bill.ActualRefundPrice,
				)
			}
		}
	}

	return nil
}
