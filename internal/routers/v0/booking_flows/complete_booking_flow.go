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
	Router.Register(courier.NewRouter(CompleteBookingFlow{}))
}

// CompleteBookingFlow 完成预售
type CompleteBookingFlow struct {
	httpx.MethodPatch
	FlowID uint64 `in:"path" name:"flowID,string"`
}

func (req CompleteBookingFlow) Path() string {
	return "/:flowID/complete"
}

func (req CompleteBookingFlow) Output(ctx context.Context) (result interface{}, err error) {
	tx := sqlx.NewTasks(global.Config.MasterDB)
	tx = tx.With(
		func(db sqlx.DBExecutor) error {
			model, err := booking_flow.GetController().GetBookingFlowByID(req.FlowID, db, true)
			if err != nil {
				return err
			}

			if model.Status != enums.BOOKING_STATUS__PROCESS {
				return errors.BookingStatusForbidComplete
			}

			return booking_flow.GetController().UpdateBookingFlow(
				model, booking_flow.UpdateBookingFlowParams{
					Status: enums.BOOKING_STATUS__COMPLETE,
				}, db,
			)
		},
	)
	err = tx.Do()
	return
}
