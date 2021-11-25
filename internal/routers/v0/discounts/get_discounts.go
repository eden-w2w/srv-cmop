package discounts

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/lib-modules/databases"
	"github.com/eden-w2w/lib-modules/modules/discounts"
)

func init() {
	Router.Register(courier.NewRouter(GetDiscounts{}))
}

// GetDiscounts 获取优惠列表
type GetDiscounts struct {
	httpx.MethodGet
	discounts.GetDiscountsParams
}

func (req GetDiscounts) Path() string {
	return ""
}

type GetDiscountsResponse struct {
	Data  []databases.MarketingDiscount `json:"data"`
	Total int                           `json:"total"`
}

func (req GetDiscounts) Output(ctx context.Context) (result interface{}, err error) {
	data, count, err := discounts.GetController().GetDiscounts(req.GetDiscountsParams, true)
	if err != nil {
		return
	}
	return &GetDiscountsResponse{
		Data:  data,
		Total: count,
	}, nil
}
