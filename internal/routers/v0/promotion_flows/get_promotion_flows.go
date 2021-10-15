package promotion_flows

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/lib-modules/databases"
	"github.com/eden-w2w/lib-modules/modules/promotion_flow"
)

func init() {
	Router.Register(courier.NewRouter(GetPromotionFlows{}))
}

// GetPromotionFlows 获取佣金流水列表
type GetPromotionFlows struct {
	httpx.MethodGet
	promotion_flow.GetPromotionFlowParams
}

type GetPromotionFlowsResponse struct {
	Data  []databases.PromotionFlow `json:"data"`
	Total int                       `json:"total"`
}

func (req GetPromotionFlows) Path() string {
	return ""
}

func (req GetPromotionFlows) Output(ctx context.Context) (result interface{}, err error) {
	data, total, err := promotion_flow.GetController().GetPromotionFlows(req.GetPromotionFlowParams, true)
	if err != nil {
		return nil, err
	}

	response := &GetPromotionFlowsResponse{
		Data:  data,
		Total: total,
	}
	return response, nil
}
