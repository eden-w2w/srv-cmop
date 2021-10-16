package goods

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/lib-modules/modules/goods"
)

func init() {
	Router.Register(courier.NewRouter(CreateGoods{}))
}

// CreateGoods 创建商品
type CreateGoods struct {
	httpx.MethodPost
	Data goods.CreateGoodsParams `in:"body"`
}

func (req CreateGoods) Path() string {
	return ""
}

func (req CreateGoods) Output(ctx context.Context) (result interface{}, err error) {
	return goods.GetController().CreateGoods(req.Data)
}
