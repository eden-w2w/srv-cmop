package goods

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/lib-modules/modules/goods"
)

func init() {
	Router.Register(courier.NewRouter(GetGoods{}))
}

// GetGoods 获取商品列表
type GetGoods struct {
	httpx.MethodGet
	goods.GetGoodsParams
}

func (req GetGoods) Path() string {
	return ""
}

func (req GetGoods) Output(ctx context.Context) (result interface{}, err error) {
	return goods.GetController().GetGoods(req.GetGoodsParams)
}
