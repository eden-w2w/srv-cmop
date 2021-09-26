package middleware

import (
	"context"
	"github.com/eden-w2w/srv-cmop/internal/databases"
)

const AuthContextKey = "Authorization"

func GetUserByContext(ctx context.Context) *databases.Administrators {
	val := ctx.Value(AuthContextKey)
	if user, ok := val.(*databases.Administrators); ok {
		return user
	}
	return nil
}
