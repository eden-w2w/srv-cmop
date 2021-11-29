package booking_flows

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-framework/sqlx"
	"github.com/eden-w2w/lib-modules/modules/booking_flow"
	"github.com/eden-w2w/srv-cmop/internal/global"
)

func init() {
	Router.Register(courier.NewRouter(UpdateBookingFlow{}))
}

// UpdateBookingFlow 更新预售单
type UpdateBookingFlow struct {
	httpx.MethodPatch
	FlowID uint64                               `in:"path" name:"flowID,string"`
	Data   booking_flow.UpdateBookingFlowParams `in:"body"`
}

func (req UpdateBookingFlow) Path() string {
	return "/:flowID"
}

func (req UpdateBookingFlow) Output(ctx context.Context) (result interface{}, err error) {
	tx := sqlx.NewTasks(global.Config.MasterDB)
	tx = tx.With(
		func(db sqlx.DBExecutor) error {
			model, err := booking_flow.GetController().GetBookingFlowByID(req.FlowID, db, true)
			if err != nil {
				return err
			}
			return booking_flow.GetController().UpdateBookingFlow(model, req.Data, db)
		},
	)
	err = tx.Do()
	return
}
