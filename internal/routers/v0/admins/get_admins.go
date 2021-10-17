package admins

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/lib-modules/modules/admins"
)

func init() {
	Router.Register(courier.NewRouter(GetAdmins{}))
}

// GetAdmins 获取管理员列表
type GetAdmins struct {
	httpx.MethodGet
}

func (req GetAdmins) Path() string {
	return ""
}

func (req GetAdmins) Output(ctx context.Context) (result interface{}, err error) {
	return admins.GetController().GetAdmins()
}
