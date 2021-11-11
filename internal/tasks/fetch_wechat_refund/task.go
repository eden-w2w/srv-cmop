package fetch_wechat_refund

import (
	"fmt"
	"github.com/eden-framework/sqlx"
	"github.com/eden-w2w/lib-modules/constants/enums"
	"github.com/eden-w2w/lib-modules/modules"
	"github.com/eden-w2w/lib-modules/modules/refund_flow"
	"github.com/eden-w2w/lib-modules/modules/wechat"
	"github.com/eden-w2w/srv-cmop/internal/global"
	wechat2 "github.com/eden-w2w/srv-cmop/internal/modules/wechat"
	"github.com/eden-w2w/wechatpay-go/core"
	"github.com/eden-w2w/wechatpay-go/services/refunddomestic"
	"github.com/sirupsen/logrus"
)

func TaskFetchWechatRefundStatus() {
	// 查询所有未达终态的退款单
	flows, _, err := refund_flow.GetController().GetRefundFlows(
		refund_flow.GetRefundFlowsRequest{
			Status: enums.REFUND_STATUS__PROCESSING,
			Pagination: modules.Pagination{
				Size: -1,
			},
		}, false,
	)
	if err != nil {
		logrus.Errorf("[TaskFetchWechatRefundStatus] refund_flow.GetController().GetRefundFlows err: %v", err)
		return
	}

	for _, flow := range flows {
		req := refunddomestic.QueryByOutRefundNoRequest{
			OutRefundNo: core.String(fmt.Sprintf("%d", flow.FlowID)),
			SubMchid:    nil,
		}
		refund, err := wechat.GetController().QueryRefundByOutRefundID(req)
		if err != nil {
			continue
		}
		status, err := enums.ParseRefundStatusFromString(string(*refund.Status))
		if err != nil {
			logrus.Errorf(
				"[TaskFetchWechatRefundStatus] enums.ParseRefundStatusFromString err: %v, RefundStatus: %s",
				err,
				*refund.Status,
			)
			continue
		}
		if status == flow.Status {
			continue
		}

		tx := sqlx.NewTasks(global.Config.MasterDB)
		tx = tx.With(
			func(db sqlx.DBExecutor) error {
				return wechat2.UpdateRefundByWechat(&flow, refund, db)
			},
		)

		err = tx.Do()
		if err != nil {
			logrus.Errorf("[TaskFetchWechatRefundStatus] tx.Do() err: %v, refund: %+v", err, refund)
		}
	}
}
