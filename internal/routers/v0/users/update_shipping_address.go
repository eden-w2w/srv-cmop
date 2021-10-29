package users

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/lib-modules/modules/user"
)

func init() {
	Router.Register(courier.NewRouter(UpdateShippingAddress{}))
}

// UpdateShippingAddress 更新收货地址
type UpdateShippingAddress struct {
	httpx.MethodPatch
	// 用户ID
	UserID uint64 `in:"path" name:"userID,string"`
	// 业务ID
	ShippingID uint64                           `in:"path" name:"shippingID,string"`
	Data       user.UpdateShippingAddressParams `in:"body"`
}

func (req UpdateShippingAddress) Path() string {
	return "/:userID/address/:shippingID"
}

func (req UpdateShippingAddress) Output(ctx context.Context) (result interface{}, err error) {
	req.Data.ShippingID = req.ShippingID
	err = user.GetController().UpdateShippingAddress(req.Data, 0, nil)
	return
}
