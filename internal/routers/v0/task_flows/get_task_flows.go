package task_flows

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/lib-modules/databases"
	"github.com/eden-w2w/lib-modules/modules/task_flow"
)

func init() {
	Router.Register(courier.NewRouter(GetTaskFlows{}))
}

// GetTaskFlows 获取任务调度日志
type GetTaskFlows struct {
	httpx.MethodGet
	task_flow.GetTaskParams
}

func (req GetTaskFlows) Path() string {
	return ""
}

type GetTaskFlowsResponse struct {
	Data  []databases.TaskFlow `json:"data"`
	Total int                  `json:"total"`
}

func (req GetTaskFlows) Output(ctx context.Context) (result interface{}, err error) {
	list, count, err := task_flow.GetController().GetTaskFlows(req.GetTaskParams, true)
	if err != nil {
		return nil, err
	}
	resp := &GetTaskFlowsResponse{
		Data:  list,
		Total: count,
	}
	return resp, nil
}
