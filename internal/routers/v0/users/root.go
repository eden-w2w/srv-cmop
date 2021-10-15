package users

import "github.com/eden-framework/courier"

var Router = courier.NewRouter(UsersRouter{})

type UsersRouter struct {
	courier.EmptyOperator
}

func (UsersRouter) Path() string {
	return "/users"
}
