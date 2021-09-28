package middleware

import (
	"context"
	"github.com/eden-w2w/lib-modules/modules/admins"
)

// Authorization 认证中间件
type Authorization struct {
	Authorization string `name:"Authorization" in:"header" validate:"@string[0,256]"`
}

func (req Authorization) ContextKey() string {
	return AuthContextKey
}

func (req Authorization) Output(ctx context.Context) (result interface{}, err error) {
	return admins.GetController().GetAdminByToken(req.Authorization)
}
