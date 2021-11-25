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
	PaymentStatus    []KVItem `json:"paymentStatus"`
	RefundStatus     []KVItem `json:"refundStatus"`
	BookingType      []KVItem `json:"bookingType"`
	BookingStatus    []KVItem `json:"bookingStatus"`
	DiscountType     []KVItem `json:"discountType"`
	DiscountStatus   []KVItem `json:"discountStatus"`
	DiscountCal      []KVItem `json:"discountCal"`
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
		response.OrderStatus = append(
			response.OrderStatus, KVItem{
				Value: i[0],
				Label: i[1],
			},
		)
	}
	paymentMethod := (enums.PaymentMethod(0)).Enums()
	for _, i := range paymentMethod {
		response.PaymentMethod = append(
			response.PaymentMethod, KVItem{
				Value: i[0],
				Label: i[1],
			},
		)
	}
	settlementStatus := (enums.SettlementStatus(0)).Enums()
	for _, i := range settlementStatus {
		response.SettlementStatus = append(
			response.SettlementStatus, KVItem{
				Value: i[0],
				Label: i[1],
			},
		)
	}
	paymentStatus := (enums.PaymentStatus(0)).Enums()
	for _, i := range paymentStatus {
		response.PaymentStatus = append(
			response.PaymentStatus, KVItem{
				Value: i[0],
				Label: i[1],
			},
		)
	}
	refundStatus := (enums.RefundStatus(0)).Enums()
	for _, i := range refundStatus {
		response.RefundStatus = append(
			response.RefundStatus, KVItem{
				Value: i[0],
				Label: i[1],
			},
		)
	}
	bookingType := (enums.BookingType(0)).Enums()
	for _, i := range bookingType {
		response.BookingType = append(
			response.BookingType, KVItem{
				Value: i[0],
				Label: i[1],
			},
		)
	}
	bookingStatus := (enums.BookingStatus(0)).Enums()
	for _, i := range bookingStatus {
		response.BookingStatus = append(
			response.BookingStatus, KVItem{
				Value: i[0],
				Label: i[1],
			},
		)
	}
	discountType := (enums.DiscountType(0)).Enums()
	for _, i := range discountType {
		response.DiscountType = append(
			response.DiscountType, KVItem{
				Value: i[0],
				Label: i[1],
			},
		)
	}
	discountStatus := (enums.DiscountStatus(0)).Enums()
	for _, i := range discountStatus {
		response.DiscountStatus = append(
			response.DiscountStatus, KVItem{
				Value: i[0],
				Label: i[1],
			},
		)
	}
	discountCal := (enums.DiscountCal(0)).Enums()
	for _, i := range discountCal {
		response.DiscountCal = append(
			response.DiscountCal, KVItem{
				Value: i[0],
				Label: i[1],
			},
		)
	}

	return response, nil
}
