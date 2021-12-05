package freight_templates

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/lib-modules/modules/freight_template"
)

func init() {
	Router.Register(courier.NewRouter(CreateTemplateRule{}))
}

// CreateTemplateRule 创建模板规则
type CreateTemplateRule struct {
	httpx.MethodPost
	TemplateID uint64                                    `in:"path" name:"templateID,string"`
	Data       freight_template.CreateTemplateRuleParams `in:"body"`
}

func (req CreateTemplateRule) Path() string {
	return "/:templateID/rules"
}

func (req CreateTemplateRule) Output(ctx context.Context) (result interface{}, err error) {
	_, result, err = freight_template.GetController().CreateTemplateRule(req.TemplateID, req.Data, nil, nil)
	return
}
