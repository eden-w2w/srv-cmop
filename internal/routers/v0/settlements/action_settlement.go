package settlements

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-framework/sqlx"
	"github.com/eden-w2w/lib-modules/constants/enums"
	"github.com/eden-w2w/lib-modules/databases"
	"github.com/eden-w2w/lib-modules/modules/settlement_flow"
	"github.com/eden-w2w/srv-cmop/internal/global"
)

func init() {
	Router.Register(courier.NewRouter(ActionSettlement{}))
}

// ActionSettlement 对指定单据执行结算
type ActionSettlement struct {
	httpx.MethodPost
	// 结算单ID
	SettlementID uint64 `in:"path" name:"settlementID"`
}

func (req ActionSettlement) Path() string {
	return "/:settlementID/action"
}

func (req ActionSettlement) Output(ctx context.Context) (result interface{}, err error) {
	tx := sqlx.NewTasks(global.Config.MasterDB)
	var model *databases.SettlementFlow

	tx = tx.With(func(db sqlx.DBExecutor) error {
		model, err = settlement_flow.GetController().GetSettlementByID(req.SettlementID, db, true)
		return err
	})

	tx = tx.With(func(db sqlx.DBExecutor) error {
		err := settlement_flow.GetController().UpdateSettlement(model, settlement_flow.UpdateSettlementParams{
			Status: enums.SETTLEMENT_STATUS__COMPLETE,
		}, db)
		return err
	})

	err = tx.Do()
	return model, err
}
