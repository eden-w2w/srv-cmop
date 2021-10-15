package v0

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/lib-modules/modules/user"
)

func init() {
	Router.Register(courier.NewRouter(GetUserByKeyword{}))
}

// GetUserByKeyword 根据关键词搜索用户
type GetUserByKeyword struct {
	httpx.MethodGet
	user.GetUserByNameOrOpenIDParams
}

func (req GetUserByKeyword) Path() string {
	return "/userKeyword"
}

func (req GetUserByKeyword) Output(ctx context.Context) (result interface{}, err error) {
	return user.GetController().GetUserByNameOrOpenID(req.GetUserByNameOrOpenIDParams)
}
