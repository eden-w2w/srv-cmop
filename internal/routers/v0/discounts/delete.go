package discounts

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/lib-modules/modules/discounts"
)

func init() {
	Router.Register(courier.NewRouter(Delete{}))
}

// Delete 删除活动
type Delete struct {
	httpx.MethodDelete
	// ID
	DiscountID uint64 `in:"path" name:"discountID,string"`
}

func (req Delete) Path() string {
	return "/:discountID"
}

func (req Delete) Output(ctx context.Context) (result interface{}, err error) {
	err = discounts.GetController().DeleteDiscount(req.DiscountID, nil)
	return
}
