package v0

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/lib-modules/clients/gaode"
	"github.com/eden-w2w/srv-cmop/internal/global"
)

func init() {
	Router.Register(courier.NewRouter(DistrictsAll{}))
}

// DistrictsAll 获取所有行政区划
type DistrictsAll struct {
	httpx.MethodGet
}

func (req DistrictsAll) Path() string {
	return "/districts_all"
}

type DistrictsAllResponse struct {
	Data []gaode.DistrictItem
}

func (req DistrictsAll) Output(ctx context.Context) (result interface{}, err error) {
	request := gaode.DistrictRequest{
		SubDistrict: 3,
		Extensions:  "base",
		Page:        1,
		Offset:      1000,
	}
	resp, err := global.Config.ClientGaode.District(request)
	if err != nil {
		return
	}

	return resp.Districts, nil
}
