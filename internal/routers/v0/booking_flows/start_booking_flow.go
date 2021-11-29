package booking_flows

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-framework/sqlx"
	"github.com/eden-w2w/lib-modules/constants/enums"
	"github.com/eden-w2w/lib-modules/modules/booking_flow"
	"github.com/eden-w2w/srv-cmop/internal/contants/errors"
	"github.com/eden-w2w/srv-cmop/internal/global"
)

func init() {
	Router.Register(courier.NewRouter(StartBookingFlow{}))
}

// StartBookingFlow 启动预售
type StartBookingFlow struct {
	httpx.MethodPatch
	FlowID uint64 `in:"path" name:"flowID,string"`
}

func (req StartBookingFlow) Path() string {
	return "/:flowID/start"
}

func (req StartBookingFlow) Output(ctx context.Context) (result interface{}, err error) {
	tx := sqlx.NewTasks(global.Config.MasterDB)
	tx = tx.With(
		func(db sqlx.DBExecutor) error {
			model, err := booking_flow.GetController().GetBookingFlowByID(req.FlowID, db, true)
			if err != nil {
				return err
			}

			if model.Status != enums.BOOKING_STATUS__READY {
				return errors.BookingStatusForbidStart
			}

			return booking_flow.GetController().UpdateBookingFlow(
				model, booking_flow.UpdateBookingFlowParams{
					Status: enums.BOOKING_STATUS__PROCESS,
				}, db,
			)
		},
	)
	err = tx.Do()
	return
}
