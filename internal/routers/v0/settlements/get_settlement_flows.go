package settlements

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/lib-modules/databases"
	"github.com/eden-w2w/lib-modules/modules/settlement_flow"
)

func init() {
	Router.Register(courier.NewRouter(GetSettlementFlows{}))
}

// GetSettlementFlows 获取结算流水单
type GetSettlementFlows struct {
	httpx.MethodGet
	settlement_flow.GetSettlementFlowsParams
}

func (req GetSettlementFlows) Path() string {
	return ""
}

type GetSettlementFlowsResponse struct {
	Data  []databases.SettlementFlow `json:"data"`
	Total int                        `json:"total"`
}

func (req GetSettlementFlows) Output(ctx context.Context) (result interface{}, err error) {
	data, total, err := settlement_flow.GetController().GetSettlementFlows(req.GetSettlementFlowsParams, true)
	if err != nil {
		return nil, err
	}
	response := &GetSettlementFlowsResponse{
		Data:  data,
		Total: total,
	}
	return response, nil
}
