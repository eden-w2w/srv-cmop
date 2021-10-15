package orders

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/lib-modules/modules/goods"
	"github.com/eden-w2w/lib-modules/modules/order"
)

func init() {
	Router.Register(courier.NewRouter(CancelOrder{}))
}

// CancelOrder 关闭订单
type CancelOrder struct {
	httpx.MethodDelete
	// 订单ID
	OrderID uint64 `in:"path" name:"orderID,string"`
}

func (req CancelOrder) Path() string {
	return "/:orderID"
}

func (req CancelOrder) Output(ctx context.Context) (result interface{}, err error) {
	err = order.GetController().CancelOrder(req.OrderID, 0, goods.GetController().UnlockInventory)
	return
}
