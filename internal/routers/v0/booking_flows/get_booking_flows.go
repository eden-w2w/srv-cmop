package booking_flows

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/lib-modules/databases"
	"github.com/eden-w2w/lib-modules/modules/booking_flow"
)

func init() {
	Router.Register(courier.NewRouter(GetBookingFlows{}))
}

// GetBookingFlows 获取预售单列表
type GetBookingFlows struct {
	httpx.MethodGet
	booking_flow.GetBookingFlowParams
}

func (req GetBookingFlows) Path() string {
	return ""
}

type GetBookingFlowsResponse struct {
	Data  []databases.BookingFlow `json:"data"`
	Total int                     `json:"total"`
}

func (req GetBookingFlows) Output(ctx context.Context) (result interface{}, err error) {
	data, total, err := booking_flow.GetController().GetBookingFlows(req.GetBookingFlowParams, true)
	if err != nil {
		return
	}
	return &GetBookingFlowsResponse{
		Data:  data,
		Total: total,
	}, nil
}
