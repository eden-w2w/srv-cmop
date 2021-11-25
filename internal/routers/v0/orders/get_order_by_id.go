package orders

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/lib-modules/modules/order"
)

func init() {
	Router.Register(courier.NewRouter(GetOrderByID{}))
}

// GetOrderByID 通过订单号获取订单
type GetOrderByID struct {
	httpx.MethodGet

	// 订单号
	OrderID uint64 `in:"path" name:"orderID,string"`
}

func (req GetOrderByID) Path() string {
	return "/:orderID"
}

func (req GetOrderByID) Output(ctx context.Context) (result interface{}, err error) {
	o, l, err := order.GetController().GetOrder(req.OrderID, 0, nil, false)
	if err != nil {
		return nil, err
	}
	response := &order.GetOrderByIDResponse{
		OrderID:        o.OrderID,
		UserID:         o.UserID,
		NickName:       o.NickName,
		OpenID:         o.UserOpenID,
		TotalPrice:     o.TotalPrice,
		DiscountAmount: o.DiscountAmount,
		ActualAmount:   o.ActualAmount,
		PaymentMethod:  o.PaymentMethod,
		Remark:         o.Remark,
		Recipients:     l.Recipients,
		ShippingAddr:   l.ShippingAddr,
		Mobile:         l.Mobile,
		CourierCompany: l.CourierCompany,
		CourierNumber:  l.CourierNumber,
		Status:         o.Status,
		ExpiredAt:      o.ExpiredAt,
		CreatedAt:      o.CreatedAt,
		UpdatedAt:      o.UpdatedAt,
		Goods:          make([]order.GoodsListResponse, 0),
	}
	goods, err := order.GetController().GetOrderGoods(o.OrderID, nil)
	if err != nil {
		return nil, err
	}
	for _, g := range goods {
		response.Goods = append(response.Goods, order.GoodsListResponse{
			GoodsID:        g.GoodsID,
			Name:           g.Name,
			MainPicture:    g.MainPicture,
			Specifications: g.Specifications,
			Price:          g.Price,
			Amount:         g.Amount,
		})
	}
	return response, nil
}
