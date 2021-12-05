package freight_templates

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/lib-modules/modules/freight_template"
)

func init() {
	Router.Register(courier.NewRouter(GetRuleByID{}))
}

// GetRuleByID 根据ID获取模板规则
type GetRuleByID struct {
	httpx.MethodGet
	RuleID uint64 `in:"path" name:"ruleID,string"`
}

func (req GetRuleByID) Path() string {
	return "/:templateID/rules/:ruleID"
}

func (req GetRuleByID) Output(ctx context.Context) (result interface{}, err error) {
	return freight_template.GetController().GetTemplateRuleByID(req.RuleID, nil, false)
}
