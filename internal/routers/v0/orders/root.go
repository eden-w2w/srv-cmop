package orders

import "github.com/eden-framework/courier"

var Router = courier.NewRouter(OrdersRouter{})

type OrdersRouter struct {
	courier.EmptyOperator
}

func (OrdersRouter) Path() string {
	return "/orders"
}
