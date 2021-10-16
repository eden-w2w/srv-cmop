package v0

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/lib-modules/constants/enums"
)

func init() {
	AuthRouter.Register(courier.NewRouter(Enums{}))
}

// Enums 获取系统枚举配置项
type Enums struct {
	httpx.MethodGet
}

func (req Enums) Path() string {
	return "/enums"
}

type EnumsResponse struct {
	OrderStatus      []KVItem `json:"orderStatus"`
	PaymentMethod    []KVItem `json:"paymentMethod"`
	SettlementStatus []KVItem `json:"settlementStatus"`
}

type KVItem struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

func (req Enums) Output(ctx context.Context) (result interface{}, err error) {
	response := &EnumsResponse{
		OrderStatus: make([]KVItem, 0),
	}

	orderStatus := (enums.OrderStatus(0)).Enums()
	for _, i := range orderStatus {
		response.OrderStatus = append(response.OrderStatus, KVItem{
			Value: i[0],
			Label: i[1],
		})
	}
	paymentMethod := (enums.PaymentMethod(0)).Enums()
	for _, i := range paymentMethod {
		response.PaymentMethod = append(response.PaymentMethod, KVItem{
			Value: i[0],
			Label: i[1],
		})
	}
	settlementStatus := (enums.SettlementStatus(0)).Enums()
	for _, i := range settlementStatus {
		response.SettlementStatus = append(response.SettlementStatus, KVItem{
			Value: i[0],
			Label: i[1],
		})
	}

	return response, nil
}