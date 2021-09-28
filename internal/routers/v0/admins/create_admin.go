package admins

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/lib-modules/modules/admins"
)

func init() {
	Router.Register(courier.NewRouter(CreateAdmin{}))
}

// CreateAdmin 创建管理员
type CreateAdmin struct {
	httpx.MethodPost
	Body admins.LoginParams `in:"body"`
}

func (req CreateAdmin) Path() string {
	return ""
}

func (req CreateAdmin) Output(ctx context.Context) (result interface{}, err error) {
	return admins.GetController().CreateAdmin(req.Body)
}
