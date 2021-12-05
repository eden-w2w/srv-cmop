package freight_templates

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/lib-modules/modules/freight_template"
)

func init() {
	Router.Register(courier.NewRouter(UpdateTemplate{}))
}

// UpdateTemplate 更新运费模板
type UpdateTemplate struct {
	httpx.MethodPatch
	TemplateID uint64                                `in:"path" name:"templateID,string"`
	Data       freight_template.UpdateTemplateParams `in:"body"`
}

func (req UpdateTemplate) Path() string {
	return "/:templateID"
}

func (req UpdateTemplate) Output(ctx context.Context) (result interface{}, err error) {
	_, err = freight_template.GetController().UpdateTemplate(req.TemplateID, req.Data, nil, nil)
	return
}
