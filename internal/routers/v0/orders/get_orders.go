package orders

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/lib-modules/databases"
	"github.com/eden-w2w/lib-modules/modules/order"
)

func init() {
	Router.Register(courier.NewRouter(GetOrders{}))
}

// GetOrders 获取订单列表
type GetOrders struct {
	httpx.MethodGet
	order.GetOrdersParams
}

func (req GetOrders) Path() string {
	return ""
}

type GetOrdersResponse struct {
	Data  []databases.Order `json:"data"`
	Total int               `json:"total"`
}

func (req GetOrders) Output(ctx context.Context) (result interface{}, err error) {
	list, count, err := order.GetController().GetOrders(req.GetOrdersParams, true)
	if err != nil {
		return nil, err
	}
	resp := &GetOrdersResponse{
		Data:  list,
		Total: count,
	}
	return resp, nil
}
