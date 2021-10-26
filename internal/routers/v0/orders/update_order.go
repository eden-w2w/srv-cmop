package orders

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-framework/sqlx"
	"github.com/eden-w2w/lib-modules/modules/goods"
	"github.com/eden-w2w/lib-modules/modules/order"
	"github.com/eden-w2w/srv-cmop/internal/global"
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
	tx := sqlx.NewTasks(global.Config.MasterDB)
	tx = tx.With(func(db sqlx.DBExecutor) error {
		orderModel, logistics, err := order.GetController().GetOrder(req.OrderID, 0, db, true)
		if err != nil {
			return err
		}
		orderGoods, err := order.GetController().GetOrderGoods(req.OrderID, db)
		if err != nil {
			return err
		}
		return order.GetController().UpdateOrder(orderModel, logistics, orderGoods, req.Body, goods.GetController().LockInventory, goods.GetController().UnlockInventory, db)
	})

	err = tx.Do()
	return
}
