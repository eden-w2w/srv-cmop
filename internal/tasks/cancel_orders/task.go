package cancel_orders

import (
	"github.com/eden-framework/sqlx/builder"
	"github.com/eden-framework/sqlx/datatypes"
	"github.com/eden-w2w/lib-modules/constants/enums"
	"github.com/eden-w2w/lib-modules/databases"
	"github.com/eden-w2w/lib-modules/modules/order"
	"github.com/eden-w2w/srv-cmop/internal/global"
	"github.com/sirupsen/logrus"
	"time"
)

func TaskCancelOrders(unlocker order.InventoryUnlock) func() {
	return func() {
		currentTime := datatypes.MySQLTimestamp(time.Now())
		logrus.Infof("[TaskCancelExpiredOrders] start cancel expired orders for %s", currentTime.String())
		defer logrus.Infof("[TaskCancelExpiredOrders] complete cancel expired orders for %s", currentTime.String())

		model := &databases.Order{}
		condition := model.FieldStatus().Eq(enums.ORDER_STATUS__CREATED)
		condition = builder.And(condition, model.FieldExpiredAt().Lte(currentTime))
		orders, err := model.List(global.Config.MasterDB, condition, builder.OrderBy(builder.AscOrder(model.FieldExpiredAt())))
		if err != nil {
			logrus.Errorf("[TaskCancelExpiredOrders] model.List err: %v", err)
			return
		}

		for _, o := range orders {
			err := order.GetController().CancelOrder(o.OrderID, 0)
			if err != nil {
				logrus.Errorf("[TaskCancelExpiredOrders] c.CancelOrder err: %v, orderID: %d", err, o.OrderID)
			}
		}
	}
}
