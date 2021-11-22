package v0

import (
	"github.com/eden-framework/courier"
	"github.com/eden-w2w/srv-cmop/internal/routers/middleware"
	"github.com/eden-w2w/srv-cmop/internal/routers/v0/admins"
	"github.com/eden-w2w/srv-cmop/internal/routers/v0/booking_flows"
	"github.com/eden-w2w/srv-cmop/internal/routers/v0/goods"
	"github.com/eden-w2w/srv-cmop/internal/routers/v0/notify"
	"github.com/eden-w2w/srv-cmop/internal/routers/v0/orders"
	"github.com/eden-w2w/srv-cmop/internal/routers/v0/payment_flows"
	"github.com/eden-w2w/srv-cmop/internal/routers/v0/promotion_flows"
	"github.com/eden-w2w/srv-cmop/internal/routers/v0/refund_flows"
	"github.com/eden-w2w/srv-cmop/internal/routers/v0/settings"
	"github.com/eden-w2w/srv-cmop/internal/routers/v0/settlements"
	"github.com/eden-w2w/srv-cmop/internal/routers/v0/task_flows"
	"github.com/eden-w2w/srv-cmop/internal/routers/v0/users"
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
	Router.Register(notify.Router)
	AuthRouter.Register(admins.Router)
	AuthRouter.Register(goods.Router)
	AuthRouter.Register(orders.Router)
	AuthRouter.Register(users.Router)
	AuthRouter.Register(settlements.Router)
	AuthRouter.Register(promotion_flows.Router)
	AuthRouter.Register(payment_flows.Router)
	AuthRouter.Register(task_flows.Router)
	AuthRouter.Register(refund_flows.Router)
	AuthRouter.Register(settings.Router)
	AuthRouter.Register(booking_flows.Router)
}
