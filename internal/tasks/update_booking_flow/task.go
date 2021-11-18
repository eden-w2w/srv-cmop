package update_booking_flow

import (
	"github.com/eden-framework/sqlx"
	"github.com/eden-framework/sqlx/datatypes"
	"github.com/eden-w2w/lib-modules/constants/enums"
	"github.com/eden-w2w/lib-modules/modules"
	"github.com/eden-w2w/lib-modules/modules/booking_flow"
	"github.com/eden-w2w/lib-modules/modules/goods"
	"github.com/eden-w2w/srv-cmop/internal/global"
	"github.com/sirupsen/logrus"
	"time"
)

func TaskUpdateBookingFlow() {
	goodsList, err := goods.GetController().GetGoods(
		goods.GetGoodsParams{
			Pagination: modules.Pagination{
				Size:   -1,
				Offset: 0,
			},
		},
	)
	if err != nil {
		logrus.Errorf("[TaskUpdateBookingFlow] goods.GetController().GetGoods err: %v", err)
		return
	}
	for _, g := range goodsList {
		if g.Inventory == 0 && g.IsAllowBooking == datatypes.BOOL_TRUE {
			flows, err := booking_flow.GetController().GetBookingFlowByGoodsIDAndStatus(
				g.GoodsID,
				enums.BOOKING_STATUS__PROCESS,
			)
			if err != nil {
				continue
			}
			if len(flows) > 0 {
				continue
			}

			tx := sqlx.NewTasks(global.Config.MasterDB)
			tx.With(
				func(db sqlx.DBExecutor) error {
					flow, err := booking_flow.GetController().CreateBookingFlow(
						booking_flow.CreateBookingFlowParams{
							GoodsID:   g.GoodsID,
							Limit:     0,
							Type:      enums.BOOKING_TYPE__AUTO,
							StartTime: datatypes.MySQLTimestamp(time.Now()),
						}, db,
					)
					if err != nil {
						return err
					}
					return booking_flow.GetController().UpdateBookingFlow(
						flow, booking_flow.UpdateBookingFlowParams{
							Status: enums.BOOKING_STATUS__PROCESS,
						}, db,
					)
				},
			)
			_ = tx.Do()
		}
	}
}
