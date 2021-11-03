package orders

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/lib-modules/constants/enums"
	"github.com/eden-w2w/lib-modules/modules/payment_flow"
	"github.com/eden-w2w/srv-cmop/internal/global"
)

func init() {
	Router.Register(courier.NewRouter(GetOrderPayment{}))
}

// GetOrderPayment 获取支付成功的支付单
type GetOrderPayment struct {
	httpx.MethodGet

	// 订单号
	OrderID uint64 `in:"path" name:"orderID,string"`
}

func (req GetOrderPayment) Path() string {
	return "/:orderID/payments/complete"
}

func (req GetOrderPayment) Output(ctx context.Context) (result interface{}, err error) {
	flows, err := payment_flow.GetController().GetFlowByOrderIDAndStatus(
		req.OrderID,
		0,
		[]enums.PaymentStatus{enums.PAYMENT_STATUS__SUCCESS, enums.PAYMENT_STATUS__REFUND},
		global.Config.MasterDB,
	)
	if err != nil {
		return
	}
	if flows != nil && len(flows) > 0 {
		return flows[0], nil
	}
	return
}
