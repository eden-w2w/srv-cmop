package goods

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/lib-modules/modules/goods"
)

func init() {
	Router.Register(courier.NewRouter(GetGoodsByID{}))
}

// GetGoodsByID 根据ID获取商品
type GetGoodsByID struct {
	httpx.MethodGet
	// 商品ID
	GoodsID uint64 `in:"path" name:"goodsID"`
}

func (req GetGoodsByID) Path() string {
	return "/:goodsID"
}

func (req GetGoodsByID) Output(ctx context.Context) (result interface{}, err error) {
	return goods.GetController().GetGoodsByID(req.GoodsID)
}
