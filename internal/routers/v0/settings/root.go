package settings

import "github.com/eden-framework/courier"

var Router = courier.NewRouter(SettingsRouter{})

type SettingsRouter struct {
	courier.EmptyOperator
}

func (SettingsRouter) Path() string {
	return "/settings"
}
