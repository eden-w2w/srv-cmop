package freight_templates

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/lib-modules/modules/freight_template"
)

func init() {
	Router.Register(courier.NewRouter(GetTemplateByID{}))
}

// GetTemplateByID 根据ID获取模板
type GetTemplateByID struct {
	httpx.MethodGet
	TemplateID uint64 `in:"path" name:"templateID,string"`
}

func (req GetTemplateByID) Path() string {
	return "/:templateID"
}

func (req GetTemplateByID) Output(ctx context.Context) (result interface{}, err error) {
	return freight_template.GetController().GetTemplateByID(req.TemplateID, nil, false)
}
