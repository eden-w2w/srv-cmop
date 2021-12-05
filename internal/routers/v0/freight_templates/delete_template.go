package freight_templates

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/lib-modules/modules/freight_template"
	"github.com/eden-w2w/lib-modules/modules/goods"
	"github.com/eden-w2w/srv-cmop/internal/contants/errors"
)

func init() {
	Router.Register(courier.NewRouter(DeleteTemplate{}))
}

// DeleteTemplate 删除运费模板
type DeleteTemplate struct {
	httpx.MethodDelete
	TemplateID uint64 `in:"path" name:"templateID,string"`
}

func (req DeleteTemplate) Path() string {
	return "/:templateID"
}

func (req DeleteTemplate) Output(ctx context.Context) (result interface{}, err error) {
	list, err := goods.GetController().GetGoods(goods.GetGoodsParams{
		FreightTemplateID: req.TemplateID,
	})
	if err != nil {
		return
	}
	if len(list) > 0 {
		err = errors.TemplateForbidDelete
		return
	}
	_, err = freight_template.GetController().DeleteTemplate(req.TemplateID, nil, nil)
	return
}
