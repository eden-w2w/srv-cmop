package payment_flows

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/lib-modules/databases"
	"github.com/eden-w2w/lib-modules/modules/payment_flow"
)

func init() {
	Router.Register(courier.NewRouter(GetPaymentFlows{}))
}

// GetPaymentFlows 获取支付单列表
type GetPaymentFlows struct {
	httpx.MethodGet
	payment_flow.GetPaymentFlowsParams
}

func (req GetPaymentFlows) Path() string {
	return ""
}

type GetPaymentFlowsResponse struct {
	Data  []databases.PaymentFlow `json:"data"`
	Total int                     `json:"total"`
}

func (req GetPaymentFlows) Output(ctx context.Context) (result interface{}, err error) {
	data, total, err := payment_flow.GetController().GetPaymentFlows(req.GetPaymentFlowsParams, true)
	if err != nil {
		return
	}

	return GetPaymentFlowsResponse{
		Data:  data,
		Total: total,
	}, nil
}
