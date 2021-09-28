package middleware

import (
	"context"
	"github.com/eden-w2w/lib-modules/databases"
)

const AuthContextKey = "Authorization"

func GetUserByContext(ctx context.Context) *databases.Administrators {
	val := ctx.Value(AuthContextKey)
	if user, ok := val.(*databases.Administrators); ok {
		return user
	}
	return nil
}
