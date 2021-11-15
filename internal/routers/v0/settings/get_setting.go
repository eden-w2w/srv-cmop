package settings

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/lib-modules/modules/settings"
)

func init() {
	Router.Register(courier.NewRouter(GetSetting{}))
}

// GetSetting 获取系统设置
type GetSetting struct {
	httpx.MethodGet
}

func (req GetSetting) Path() string {
	return ""
}

func (req GetSetting) Output(ctx context.Context) (result interface{}, err error) {
	setting := settings.GetController().GetSetting()
	return setting, nil
}
