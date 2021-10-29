package users

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/lib-modules/modules/user"
)

func init() {
	Router.Register(courier.NewRouter(DeleteShippingAddress{}))
}

// DeleteShippingAddress 删除收货地址
type DeleteShippingAddress struct {
	httpx.MethodDelete
	// 用户ID
	UserID uint64 `in:"path" name:"userID,string"`
	// 业务ID
	ShippingID uint64 `in:"path" name:"shippingID,string"`
}

func (req DeleteShippingAddress) Path() string {
	return "/:userID/address/:shippingID"
}

func (req DeleteShippingAddress) Output(ctx context.Context) (result interface{}, err error) {
	err = user.GetController().DeleteShippingAddress(req.ShippingID, 0, nil)
	return
}
