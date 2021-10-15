package settlements

import "github.com/eden-framework/courier"

var Router = courier.NewRouter(SettlementsRouter{})

type SettlementsRouter struct {
	courier.EmptyOperator
}

func (SettlementsRouter) Path() string {
	return "/settlements"
}
