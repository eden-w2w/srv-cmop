package refund_flows

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/lib-modules/databases"
	"github.com/eden-w2w/lib-modules/modules/refund_flow"
)

func init() {
	Router.Register(courier.NewRouter(GetRefundFlows{}))
}

// GetRefundFlows 获取退款单列表
type GetRefundFlows struct {
	httpx.MethodGet
	refund_flow.GetRefundFlowsRequest `in:"body"`
}

func (req GetRefundFlows) Path() string {
	return ""
}

type GetRefundFlowsResponse struct {
	Data  []databases.RefundFlow `json:"data"`
	Total int                    `json:"total"`
}

func (req GetRefundFlows) Output(ctx context.Context) (result interface{}, err error) {
	data, total, err := refund_flow.GetController().GetRefundFlows(req.GetRefundFlowsRequest, true)
	if err != nil {
		return
	}

	return GetRefundFlowsResponse{
		Data:  data,
		Total: total,
	}, nil
}
