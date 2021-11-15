package settings

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/lib-modules/modules/settings"
)

func init() {
	Router.Register(courier.NewRouter(UpdateSetting{}))
}

// UpdateSetting 更新系统设置
type UpdateSetting struct {
	httpx.MethodPatch

	Data settings.UpdateSettingParams `in:"body"`
}

func (req UpdateSetting) Path() string {
	return ""
}

func (req UpdateSetting) Output(ctx context.Context) (result interface{}, err error) {
	err = settings.GetController().UpdateSetting(req.Data, nil)
	return
}
