package v0

import (
	"github.com/eden-framework/courier"
	"github.com/eden-w2w/srv-cmop/internal/routers/middleware"
	"github.com/eden-w2w/srv-cmop/internal/routers/v0/admins"
	"github.com/eden-w2w/srv-cmop/internal/routers/v0/goods"
	"github.com/eden-w2w/srv-cmop/internal/routers/v0/orders"
)

var Router = courier.NewRouter(V0Router{})
var AuthRouter = courier.NewRouter(middleware.Authorization{})

type V0Router struct {
	courier.EmptyOperator
}

func (V0Router) Path() string {
	return "/v0"
}

func init() {
	Router.Register(AuthRouter)
	AuthRouter.Register(admins.Router)
	AuthRouter.Register(goods.Router)
	AuthRouter.Register(orders.Router)
}
