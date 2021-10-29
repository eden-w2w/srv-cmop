package users

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/lib-modules/modules/user"
)

func init() {
	Router.Register(courier.NewRouter(CreateShippingAddress{}))
}

// CreateShippingAddress 创建收货地址
type CreateShippingAddress struct {
	httpx.MethodPost
	// 用户ID
	UserID uint64                           `in:"path" name:"userID"`
	Data   user.CreateShippingAddressParams `in:"body"`
}

func (req CreateShippingAddress) Path() string {
	return "/:userID/address"
}

func (req CreateShippingAddress) Output(ctx context.Context) (result interface{}, err error) {
	req.Data.UserID = req.UserID
	return user.GetController().CreateShippingAddress(req.Data, nil)
}
