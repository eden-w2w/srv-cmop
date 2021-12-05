package freight_templates

import "github.com/eden-framework/courier"

var Router = courier.NewRouter(FreightTemplateRouters{})

type FreightTemplateRouters struct {
	courier.EmptyOperator
}

func (FreightTemplateRouters) Path() string {
	return "/freight_templates"
}
