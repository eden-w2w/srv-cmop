package booking_flows

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/lib-modules/modules/booking_flow"
)

func init() {
	Router.Register(courier.NewRouter(CreateBookingFlow{}))
}

// CreateBookingFlow 创建预售单
type CreateBookingFlow struct {
	httpx.MethodPost

	Data booking_flow.CreateBookingFlowParams `in:"body"`
}

func (req CreateBookingFlow) Path() string {
	return ""
}

func (req CreateBookingFlow) Output(ctx context.Context) (result interface{}, err error) {
	return booking_flow.GetController().CreateBookingFlow(req.Data, nil)
}
