package promotion_flows

import "github.com/eden-framework/courier"

var Router = courier.NewRouter(PromotionFlowRouter{})

type PromotionFlowRouter struct {
	courier.EmptyOperator
}

func (PromotionFlowRouter) Path() string {
	return "/promotion_flows"
}
