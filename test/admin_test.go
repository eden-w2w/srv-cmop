package test

import (
	"context"
	"github.com/eden-w2w/lib-modules/databases"
	admins2 "github.com/eden-w2w/lib-modules/modules/admins"
	"github.com/eden-w2w/srv-cmop/internal/routers/v0/admins"
	"github.com/stretchr/testify/require"
	"testing"
)

func testCreateAdmin(t *testing.T) {
	ctx := context.Background()
	request := &admins.CreateAdmin{
		Body: admins2.LoginParams{
			UserName: "admin",
			Password: "123456",
		},
	}
	resp, err := request.Output(ctx)
	require.Nil(t, err)

	adminModel = resp.(*databases.Administrators)
	require.Equal(t, "admin", adminModel.UserName)
	require.NotEmpty(t, adminModel.Token)
}
