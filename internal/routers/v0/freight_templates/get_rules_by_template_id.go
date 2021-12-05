package freight_templates

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/lib-modules/modules/freight_template"
)

func init() {
	Router.Register(courier.NewRouter(GetRulesByTemplateID{}))
}

// GetRulesByTemplateID 根据模板ID获取模板规则
type GetRulesByTemplateID struct {
	httpx.MethodGet
	TemplateID uint64 `in:"path" name:"templateID,string"`
	freight_template.GetTemplateRuleParams
}

func (req GetRulesByTemplateID) Path() string {
	return "/:templateID/rules"
}

func (req GetRulesByTemplateID) Output(ctx context.Context) (result interface{}, err error) {
	return freight_template.GetController().GetTemplateRules(req.TemplateID, req.GetTemplateRuleParams)
}
