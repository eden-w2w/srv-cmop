package notify

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-framework/courier/transport_http"
	"github.com/eden-framework/sqlx"
	errors "github.com/eden-w2w/lib-modules/constants/general_errors"
	"github.com/eden-w2w/lib-modules/modules/wechat"
	"github.com/eden-w2w/srv-cmop/internal/global"
	wechatModule "github.com/eden-w2w/srv-cmop/internal/modules/wechat"
)

func init() {
	Router.Register(courier.NewRouter(RefundNotify{}))
}

// RefundNotify 微信支付回调
type RefundNotify struct {
	httpx.MethodPost
}

func (req RefundNotify) Path() string {
	return "/refund"
}

func (req RefundNotify) Output(ctx context.Context) (result interface{}, err error) {
	request := transport_http.GetRequest(ctx)
	_, refund, err := wechat.GetController().ParseWechatRefundNotify(ctx, request)
	if err != nil {
		return nil, err
	}

	tx := sqlx.NewTasks(global.Config.MasterDB)
	tx = tx.With(
		func(db sqlx.DBExecutor) error {
			return wechatModule.UpdateRefundByWechatNotify(refund, db)
		})

	err = tx.Do()
	if err != nil {
		return nil, errors.InternalError
	}
	return wechat.WechatNotifyResponse{
		Code:    "SUCCESS",
		Message: "",
	}, nil
}
