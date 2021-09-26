package orders

import (
	"github.com/eden-framework/sqlx"
	"github.com/eden-w2w/srv-cmop/internal/global"
)

var controller *Controller

func GetController() *Controller {
	if controller == nil {
		controller = newController(global.Config.MasterDB)
	}
	return controller
}

type Controller struct {
	db sqlx.DBExecutor
}

func newController(db sqlx.DBExecutor) *Controller {
	return &Controller{db: db}
}
