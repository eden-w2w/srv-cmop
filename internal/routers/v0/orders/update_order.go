package orders

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/lib-modules/modules/order"
)

func init() {
	Router.Register(courier.NewRouter(UpdateOrder{}))
}

// UpdateOrder 更新订单
type UpdateOrder struct {
	httpx.MethodPatch
	// 订单号
	OrderID uint64                  `in:"path" name:"orderID,string"`
	Body    order.UpdateOrderParams `in:"body"`
}

func (req UpdateOrder) Path() string {
	return "/:orderID"
}

func (req UpdateOrder) Output(ctx context.Context) (result interface{}, err error) {
	panic("implement me")
}
