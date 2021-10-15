package users

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/lib-modules/databases"
	"github.com/eden-w2w/lib-modules/modules/user"
)

func init() {
	Router.Register(courier.NewRouter(GetUsers{}))
}

// GetUsers 获取用户列表
type GetUsers struct {
	httpx.MethodGet

	user.GetUsersParams
}

func (req GetUsers) Path() string {
	return ""
}

type GetUsersResponse struct {
	Data  []databases.User `json:"data"`
	Total int              `json:"total"`
}

func (req GetUsers) Output(ctx context.Context) (result interface{}, err error) {
	list, count, err := user.GetController().GetUsers(req.GetUsersParams, true)
	if err != nil {
		return nil, err
	}
	resp := &GetUsersResponse{
		Data:  list,
		Total: count,
	}
	return resp, nil
}
