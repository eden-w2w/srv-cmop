package goods

import "github.com/eden-framework/courier"

var Router = courier.NewRouter(GoodsRouter{})

type GoodsRouter struct {
	courier.EmptyOperator
}

func (GoodsRouter) Path() string {
	return "/goods"
}
