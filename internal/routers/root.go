package routers

import (
	"github.com/eden-framework/context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/swagger"

	v0 "github.com/eden-w2w/srv-cmop/internal/routers/v0"
)

var Router = courier.NewRouter(RootRouter{})

func init() {
	Router.Register(v0.Router)
	if !context.IsOnline() {
		Router.Register(swagger.SwaggerRouter)
	}
}

type RootRouter struct {
	courier.EmptyOperator
}

func (RootRouter) Path() string {
	return "/cmop"
}
