package refund_flows

import "github.com/eden-framework/courier"

var Router = courier.NewRouter(RefundFlowRouter{})

type RefundFlowRouter struct {
	courier.EmptyOperator
}

func (RefundFlowRouter) Path() string {
	return "/refund_flows"
}
