package orders

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/lib-modules/modules/order"
)

func init() {
	Router.Register(courier.NewRouter(GetOrderLogistics{}))
}

// GetOrderLogistics 获取订单物流信息
type GetOrderLogistics struct {
	httpx.MethodGet

	// 订单号
	OrderID uint64 `in:"path" name:"orderID,string"`
}

func (req GetOrderLogistics) Path() string {
	return "/:orderID/logistics"
}

func (req GetOrderLogistics) Output(ctx context.Context) (result interface{}, err error) {
	return order.GetController().GetOrderLogistics(req.OrderID)
}
