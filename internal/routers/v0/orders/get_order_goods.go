package orders

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/lib-modules/modules/order"
)

func init() {
	Router.Register(courier.NewRouter(GetOrderGoods{}))
}

// GetOrderGoods 获取订单商品列表
type GetOrderGoods struct {
	httpx.MethodGet

	// 订单号
	OrderID uint64 `in:"path" name:"orderID,string"`
}

func (req GetOrderGoods) Path() string {
	return "/:orderID/goods"
}

func (req GetOrderGoods) Output(ctx context.Context) (result interface{}, err error) {
	return order.GetController().GetOrderGoods(req.OrderID, nil)
}
