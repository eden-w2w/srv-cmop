package notify

import "github.com/eden-framework/courier"

var Router = courier.NewRouter(NotifyRouter{})

type NotifyRouter struct {
	courier.EmptyOperator
}

func (NotifyRouter) Path() string {
	return "/notify"
}
