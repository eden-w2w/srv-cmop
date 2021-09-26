package v0

import (
	"github.com/eden-framework/courier"
)

var Router = courier.NewRouter(V0Router{})

type V0Router struct {
	courier.EmptyOperator
}

func (V0Router) Path() string {
	return "/v0"
}
