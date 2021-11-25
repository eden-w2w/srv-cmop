package discounts

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-framework/sqlx"
	"github.com/eden-w2w/lib-modules/constants/enums"
	"github.com/eden-w2w/lib-modules/constants/general_errors"
	"github.com/eden-w2w/lib-modules/modules/discounts"
	"github.com/eden-w2w/srv-cmop/internal/global"
)

func init() {
	Router.Register(courier.NewRouter(Stop{}))
}

// Stop 停止活动
type Stop struct {
	httpx.MethodPost
	// ID
	DiscountID uint64 `in:"path" name:"discountID,string"`
}

func (req Stop) Path() string {
	return "/:discountID/stop"
}

func (req Stop) Output(ctx context.Context) (result interface{}, err error) {
	tx := sqlx.NewTasks(global.Config.MasterDB)
	tx = tx.With(
		func(db sqlx.DBExecutor) error {
			model, err := discounts.GetController().GetDiscountByID(req.DiscountID, db, true)
			if err != nil {
				return err
			}
			if model.Status == enums.DISCOUNT_STATUS__STOP {
				return nil
			}
			if model.Status == enums.DISCOUNT_STATUS__READY {
				return general_errors.BadRequest.StatusError().WithErrTalk().WithMsg("当前活动状态为就绪时，无需停止")
			}
			err = discounts.GetController().UpdateDiscount(
				model, discounts.UpdateDiscountParams{
					Status: enums.DISCOUNT_STATUS__STOP,
				}, db,
			)
			return nil
		},
	)
	err = tx.Do()
	return
}
