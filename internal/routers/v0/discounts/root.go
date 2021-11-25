package discounts

import "github.com/eden-framework/courier"

var Router = courier.NewRouter(DiscountsRouter{})

type DiscountsRouter struct {
	courier.EmptyOperator
}

func (DiscountsRouter) Path() string {
	return "/discounts"
}
