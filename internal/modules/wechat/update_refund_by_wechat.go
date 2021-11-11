package wechat

import (
	"github.com/eden-framework/sqlx"
	"github.com/eden-w2w/lib-modules/constants/enums"
	"github.com/eden-w2w/lib-modules/constants/general_errors"
	"github.com/eden-w2w/lib-modules/databases"
	"github.com/eden-w2w/lib-modules/modules/refund_flow"
	"github.com/eden-w2w/lib-modules/modules/wechat"
	"github.com/eden-w2w/wechatpay-go/services/refunddomestic"
	"github.com/sirupsen/logrus"
	"strconv"
)

func UpdateRefundByWechatNotify(notify *refunddomestic.RefundNotify, db sqlx.DBExecutor) error {
	status, err := enums.ParseRefundStatusFromString(*notify.RefundStatus)
	if err != nil {
		logrus.Errorf(
			"[RefundNotify] enums.ParseRefundStatusFromString err: %v, RefundStatus: %s",
			err,
			*notify.RefundStatus,
		)
		return err
	}

	flowID, err := strconv.ParseUint(*notify.OutRefundNo, 10, 64)
	if err != nil {
		logrus.Errorf("[RefundNotify] strconv.ParseUint err: %v, OutRefundNo: %s", err, *notify.OutRefundNo)
		return general_errors.InternalError
	}

	var refundFlow *databases.RefundFlow
	refundFlow, err = refund_flow.GetController().GetRefundFlowByFlowID(flowID, db, true)
	if err != nil {
		return err
	}

	if status == refundFlow.Status {
		return nil
	}

	refund, err := wechat.GetController().QueryRefundByOutRefundID(
		refunddomestic.QueryByOutRefundNoRequest{
			OutRefundNo: notify.OutRefundNo,
			SubMchid:    nil,
		},
	)
	if err != nil {
		return err
	}

	return UpdateRefundByWechat(refundFlow, refund, db)
}

func UpdateRefundByWechat(flow *databases.RefundFlow, refund *refunddomestic.Refund, db sqlx.DBExecutor) error {
	status, err := enums.ParseRefundStatusFromString(string(*refund.Status))
	if err != nil {
		logrus.Errorf(
			"[RefundNotify] enums.ParseRefundStatusFromString err: %v, RefundStatus: %s",
			err,
			*refund.Status,
		)
		return err
	}
	channel, err := enums.ParseRefundChannelFromString(string(*refund.Channel))
	if err != nil {
		logrus.Errorf(
			"[RefundNotify] enums.ParseRefundChannelFromString err: %v, RefundChannel: %s",
			err,
			*refund.Channel,
		)
		return err
	}
	updateParams := refund_flow.UpdateRefundFlowRequest{
		Status:  status,
		Channel: channel,
		Account: *refund.UserReceivedAccount,
	}
	return refund_flow.GetController().UpdateRefundFlow(flow.FlowID, updateParams, db)
}
