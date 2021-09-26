package goods

import (
	"github.com/eden-framework/sqlx"
	"github.com/eden-framework/sqlx/builder"
	"github.com/eden-w2w/srv-cmop/internal/contants/types"

	"github.com/eden-w2w/srv-cmop/internal/modules"
)

type GetGoodsParams struct {
	modules.Pagination
}

func (p GetGoodsParams) Conditions(db sqlx.DBExecutor) builder.SqlCondition {
	return nil
}

func (p GetGoodsParams) Additions() []builder.Addition {
	var additions = make([]builder.Addition, 0)

	if p.Size != 0 {
		limit := builder.Limit(int64(p.Size))
		if p.Offset != 0 {
			limit.Offset(int64(p.Offset))
		}
		additions = append(additions, limit)
	}

	return additions
}

type UpdateGoodsParams struct {
	// 名称
	Name string `json:"name" default:"" in:"body"`
	// 描述
	Comment string `json:"comment" default:"" in:"body"`
	// 发货地
	DispatchAddr string `json:"dispatchAddr" default:"" in:"body"`
	// 销量
	Sales int `json:"sales" default:"" in:"body"`
	// 标题图片
	MainPicture string `json:"mainPicture" default:"" in:"body"`
	// 所有展示图片
	Pictures types.GoodsPictures `json:"pictures" default:"" in:"body"`
	// 规格
	Specifications []string `json:"specifications" default:"" in:"body"`
	// 活动
	Activities []string `json:"activities" default:"" in:"body"`
	// 物流政策
	LogisticPolicy string `json:"logisticPolicy" default:"" in:"body"`
	// 价格
	Price uint64 `json:"price" default:"" in:"body"`
	// 库存
	Inventory uint64 `json:"inventory" default:"" in:"body"`
	// 详细介绍
	Detail string `json:"detail" default:"" in:"body"`
}
