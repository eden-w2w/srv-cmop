package booking_flows

import "github.com/eden-framework/courier"

var Router = courier.NewRouter(BookingFlowsRouter{})

type BookingFlowsRouter struct {
	courier.EmptyOperator
}

func (BookingFlowsRouter) Path() string {
	return "/booking_flows"
}
