package goods

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/srv-cmop/internal/modules/goods"
)

func init() {
	Router.Register(courier.NewRouter(UpdateGoodsByID{}))
}

// UpdateGoodsByID 根据ID更新商品
type UpdateGoodsByID struct {
	httpx.MethodPut

	// 商品ID
	GoodsID uint64                  `in:"path" name:"goodsID,string"`
	Body    goods.UpdateGoodsParams `in:"body"`
}

func (req UpdateGoodsByID) Path() string {
	return "/:goodsID"
}

func (req UpdateGoodsByID) Output(ctx context.Context) (result interface{}, err error) {
	err = goods.GetController().UpdateGoods(req.GoodsID, req.Body)
	return
}
