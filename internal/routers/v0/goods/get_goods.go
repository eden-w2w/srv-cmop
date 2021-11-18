package goods

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/lib-modules/constants/enums"
	"github.com/eden-w2w/lib-modules/databases"
	"github.com/eden-w2w/lib-modules/modules/booking_flow"
	"github.com/eden-w2w/lib-modules/modules/goods"
)

func init() {
	Router.Register(courier.NewRouter(GetGoods{}))
}

// GetGoods 获取商品列表
type GetGoods struct {
	httpx.MethodGet
	goods.GetGoodsParams
}

func (req GetGoods) Path() string {
	return ""
}

type GetGoodsResponse struct {
	databases.Goods
	// 预售单ID
	BookingFlowID *uint64 `json:"bookingFlowID,string" default:""`
	// 预售销量
	BookingSales *uint32 `json:"bookingSales" default:""`
	// 预售模式
	BookingType *enums.BookingType `json:"bookingType" default:""`
	// 预售状态
	BookingStatus *enums.BookingStatus `json:"bookingStatus" default:""`
}

func (req GetGoods) Output(ctx context.Context) (result interface{}, err error) {
	data, err := goods.GetController().GetGoods(req.GetGoodsParams)
	if err != nil {
		return
	}

	resp := make([]GetGoodsResponse, 0)
	for _, gModel := range data {
		item := GetGoodsResponse{
			Goods: gModel,
		}
		flows, err := booking_flow.GetController().GetBookingFlowByGoodsIDAndStatus(
			gModel.GoodsID,
			enums.BOOKING_STATUS__PROCESS,
		)
		if err != nil {
			return nil, err
		}
		if len(flows) > 0 {
			item.BookingFlowID = &flows[0].FlowID
			item.BookingSales = &flows[0].Sales
			item.BookingType = &flows[0].Type
			item.BookingStatus = &flows[0].Status
		}
		resp = append(resp, item)
	}
	return resp, nil
}
