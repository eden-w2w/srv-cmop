package admins

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/lib-modules/modules/admins"
)

func init() {
	Router.Register(courier.NewRouter(ResetPassword{}))
}

// ResetPassword 重置管理员密码
type ResetPassword struct {
	httpx.MethodPut
	// 管理员ID
	AdministratorID uint64                     `in:"path" name:"adminID,string"`
	Body            admins.ResetPasswordParams `in:"body"`
}

func (req ResetPassword) Path() string {
	return "/:adminID/password"
}

func (req ResetPassword) Output(ctx context.Context) (result interface{}, err error) {
	err = admins.GetController().ResetPassword(req.AdministratorID, req.Body)
	return
}
