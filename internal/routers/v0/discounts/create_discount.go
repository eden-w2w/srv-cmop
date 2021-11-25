package discounts

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/lib-modules/modules/discounts"
)

func init() {
	Router.Register(courier.NewRouter(CreateDiscount{}))
}

// CreateDiscount 创建优惠
type CreateDiscount struct {
	httpx.MethodPost
	Data discounts.CreateDiscountParams `in:"body"`
}

func (req CreateDiscount) Path() string {
	return ""
}

func (req CreateDiscount) Output(ctx context.Context) (result interface{}, err error) {
	model, err := discounts.GetController().CreateDiscount(req.Data, nil)
	if err != nil {
		return
	}
	return model, nil
}
