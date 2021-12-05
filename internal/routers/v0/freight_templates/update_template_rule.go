package freight_templates

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/lib-modules/modules/freight_template"
)

func init() {
	Router.Register(courier.NewRouter(UpdateTemplateRule{}))
}

// UpdateTemplateRule 更新模板规则
type UpdateTemplateRule struct {
	httpx.MethodPatch
	RuleID uint64                                    `in:"path" name:"ruleID,string"`
	Data   freight_template.UpdateTemplateRuleParams `in:"body"`
}

func (req UpdateTemplateRule) Path() string {
	return "/:templateID/rules/:ruleID"
}

func (req UpdateTemplateRule) Output(ctx context.Context) (result interface{}, err error) {
	_, err = freight_template.GetController().UpdateTemplateRule(req.RuleID, req.Data, nil, nil)
	return
}
