package task_flows

import "github.com/eden-framework/courier"

var Router = courier.NewRouter(TaskFlowsRouter{})

type TaskFlowsRouter struct {
	courier.EmptyOperator
}

func (TaskFlowsRouter) Path() string {
	return "/task_flows"
}
