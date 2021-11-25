package discounts

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/lib-modules/modules/discounts"
)

func init() {
	Router.Register(courier.NewRouter(UpdateDiscount{}))
}

// UpdateDiscount 更新优惠活动信息
type UpdateDiscount struct {
	httpx.MethodPatch
	// ID
	DiscountID uint64                         `in:"path" name:"discountID,string"`
	Data       discounts.UpdateDiscountParams `in:"body"`
}

func (req UpdateDiscount) Path() string {
	return "/:discountID"
}

func (req UpdateDiscount) Output(ctx context.Context) (result interface{}, err error) {
	model, err := discounts.GetController().GetDiscountByID(req.DiscountID, nil, false)
	if err != nil {
		return
	}

	err = discounts.GetController().UpdateDiscount(model, req.Data, nil)
	return
}
