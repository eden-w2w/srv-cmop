package discounts

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-framework/sqlx"
	"github.com/eden-w2w/lib-modules/constants/enums"
	"github.com/eden-w2w/lib-modules/modules/discounts"
	"github.com/eden-w2w/srv-cmop/internal/global"
)

func init() {
	Router.Register(courier.NewRouter(Start{}))
}

// Start 开启活动
type Start struct {
	httpx.MethodPost
	// ID
	DiscountID uint64 `in:"path" name:"discountID,string"`
}

func (req Start) Path() string {
	return "/:discountID/start"
}

func (req Start) Output(ctx context.Context) (result interface{}, err error) {
	tx := sqlx.NewTasks(global.Config.MasterDB)
	tx = tx.With(
		func(db sqlx.DBExecutor) error {
			model, err := discounts.GetController().GetDiscountByID(req.DiscountID, db, true)
			if err != nil {
				return err
			}
			if model.Status == enums.DISCOUNT_STATUS__PROCESS {
				return nil
			}
			err = discounts.GetController().UpdateDiscount(
				model, discounts.UpdateDiscountParams{
					Status: enums.DISCOUNT_STATUS__PROCESS,
				}, db,
			)
			return nil
		},
	)
	err = tx.Do()
	return
}
