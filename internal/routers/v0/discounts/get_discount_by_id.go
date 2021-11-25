package discounts

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/lib-modules/modules/discounts"
)

func init() {
	Router.Register(courier.NewRouter(GetDiscountByID{}))
}

// GetDiscountByID 获取优惠
type GetDiscountByID struct {
	httpx.MethodGet
	// ID
	DiscountID uint64 `in:"path" name:"discountID,string"`
}

func (req GetDiscountByID) Path() string {
	return "/:discountID"
}

func (req GetDiscountByID) Output(ctx context.Context) (result interface{}, err error) {
	return discounts.GetController().GetDiscountByID(req.DiscountID, nil, false)
}
