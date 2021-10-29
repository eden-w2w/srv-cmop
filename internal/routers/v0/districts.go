package v0

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/lib-modules/clients/gaode"
	"github.com/eden-w2w/srv-cmop/internal/global"
)

func init() {
	Router.Register(courier.NewRouter(Districts{}))
}

// Districts 获取省市区行政划分
type Districts struct {
	httpx.MethodGet

	Keyword string `in:"query" name:"keyword" default:""`
}

func (req Districts) Path() string {
	return "/districts"
}

type DistrictResponse struct {
	Data []gaode.DistrictItem
}

func (req Districts) Output(ctx context.Context) (result interface{}, err error) {
	request := gaode.DistrictRequest{
		Keywords:    req.Keyword,
		SubDistrict: 1,
		Extensions:  "base",
		Page:        1,
		Offset:      20,
	}
	resp, err := global.Config.ClientGaode.District(request)
	if err != nil {
		return
	}

	return resp.Districts, nil
}
