package freight_templates

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/lib-modules/modules/freight_template"
)

func init() {
	Router.Register(courier.NewRouter(DeleteTemplateRule{}))
}

// DeleteTemplateRule 删除模板规则
type DeleteTemplateRule struct {
	httpx.MethodDelete
	RuleID uint64 `in:"path" name:"ruleID,string"`
}

func (req DeleteTemplateRule) Path() string {
	return "/:templateID/rules/:ruleID"
}

func (req DeleteTemplateRule) Output(ctx context.Context) (result interface{}, err error) {
	_, err = freight_template.GetController().DeleteTemplateRuleByID(req.RuleID, nil, nil)
	return
}
