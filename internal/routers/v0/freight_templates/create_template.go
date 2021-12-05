package freight_templates

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/lib-modules/modules/freight_template"
)

func init() {
	Router.Register(courier.NewRouter(CreateTemplate{}))
}

// CreateTemplate 创建运费模板
type CreateTemplate struct {
	httpx.MethodPost
	Data freight_template.CreateTemplateParams `in:"body"`
}

func (req CreateTemplate) Path() string {
	return ""
}

func (req CreateTemplate) Output(ctx context.Context) (result interface{}, err error) {
	_, result, err = freight_template.GetController().CreateTemplate(req.Data, nil, nil)
	return
}
