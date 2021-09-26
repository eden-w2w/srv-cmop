package admins

import "github.com/eden-framework/courier"

var Router = courier.NewRouter(AdminRouter{})

type AdminRouter struct {
	courier.EmptyOperator
}

func (AdminRouter) Path() string {
	return "/admins"
}
