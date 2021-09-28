package orders

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/lib-modules/modules/order"
)

func init() {
	Router.Register(courier.NewRouter(GetOrderByID{}))
}

// GetOrderByID 通过订单号获取订单
type GetOrderByID struct {
	httpx.MethodGet

	// 订单号
	OrderID uint64 `in:"path" name:"orderID,string"`
}

func (req GetOrderByID) Path() string {
	return "/:orderID"
}

func (req GetOrderByID) Output(ctx context.Context) (result interface{}, err error) {
	return order.GetController().GetOrder(req.OrderID, 0, nil, false)
}
