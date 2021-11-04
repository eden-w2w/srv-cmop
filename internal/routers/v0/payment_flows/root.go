package payment_flows

import "github.com/eden-framework/courier"

var Router = courier.NewRouter(PaymentFlowsRouter{})

type PaymentFlowsRouter struct {
	courier.EmptyOperator
}

func (PaymentFlowsRouter) Path() string {
	return "/payment_flows"
}
