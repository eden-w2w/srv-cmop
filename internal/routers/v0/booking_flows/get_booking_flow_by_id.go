package booking_flows

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/lib-modules/modules/booking_flow"
)

func init() {
	Router.Register(courier.NewRouter(GetBookingFlowByID{}))
}

// GetBookingFlowByID 获取预售单
type GetBookingFlowByID struct {
	httpx.MethodGet
	FlowID uint64 `in:"path" name:"flowID,string"`
}

func (req GetBookingFlowByID) Path() string {
	return "/:flowID"
}

func (req GetBookingFlowByID) Output(ctx context.Context) (result interface{}, err error) {
	return booking_flow.GetController().GetBookingFlowByID(req.FlowID, nil, false)
}
