package settlements

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/lib-modules/modules/settlement_flow"
)

func init() {
	Router.Register(courier.NewRouter(TaskSettlements{}))
}

// TaskSettlements 触发结算任务
type TaskSettlements struct {
	httpx.MethodPost
	Data struct {
		Name string `in:"body" json:"name"`
	} `in:"body"`
}

func (req TaskSettlements) Path() string {
	return ""
}

func (req TaskSettlements) Output(ctx context.Context) (result interface{}, err error) {
	err = settlement_flow.GetController().RunTaskSettlement(req.Data.Name)
	return
}
