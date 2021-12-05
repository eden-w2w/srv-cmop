package freight_templates

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/lib-modules/databases"
	"github.com/eden-w2w/lib-modules/modules/freight_template"
)

func init() {
	Router.Register(courier.NewRouter(GetTemplates{}))
}

// GetTemplates 获取运费模板
type GetTemplates struct {
	httpx.MethodGet
	freight_template.GetTemplatesParams
}

func (req GetTemplates) Path() string {
	return ""
}

type GetTemplatesResponse struct {
	Data  []databases.FreightTemplate `json:"data"`
	Total int                         `json:"total"`
}

func (req GetTemplates) Output(ctx context.Context) (result interface{}, err error) {
	list, count, err := freight_template.GetController().GetTemplates(req.GetTemplatesParams, true)
	if err != nil {
		return
	}

	return &GetTemplatesResponse{
		Data:  list,
		Total: count,
	}, nil
}
