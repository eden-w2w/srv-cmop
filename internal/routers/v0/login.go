package v0

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/srv-cmop/internal/modules/admins"
)

func init() {
	Router.Register(courier.NewRouter(Login{}))
}

// Login 登录接口
type Login struct {
	httpx.MethodPost
	Body admins.LoginParams `in:"body"`
}

func (req Login) Path() string {
	return "/login"
}

func (req Login) Output(ctx context.Context) (result interface{}, err error) {
	admin, err := admins.GetController().LoginCheck(req.Body)
	if err != nil {
		return nil, err
	}

	return admins.GetController().RefreshToken(admin.AdministratorsID)
}
