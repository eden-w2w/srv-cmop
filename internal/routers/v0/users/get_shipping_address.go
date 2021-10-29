package users

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/lib-modules/modules/user"
)

func init() {
	Router.Register(courier.NewRouter(GetShippingAddress{}))
}

// GetShippingAddress 获取收货地址
type GetShippingAddress struct {
	httpx.MethodGet
	// 用户ID
	UserID uint64 `in:"path" name:"userID"`
}

func (req GetShippingAddress) Path() string {
	return "/:userID/address"
}

func (req GetShippingAddress) Output(ctx context.Context) (result interface{}, err error) {
	return user.GetController().GetShippingAddressByUserID(req.UserID)
}
