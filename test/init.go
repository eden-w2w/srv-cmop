package test

import (
	"github.com/eden-framework/context"
	"github.com/eden-framework/eden-framework/pkg/application"
	"github.com/eden-w2w/lib-modules/databases"
	"github.com/eden-w2w/lib-modules/modules/admins"
	"github.com/eden-w2w/lib-modules/modules/events"
	"github.com/eden-w2w/lib-modules/modules/goods"
	"github.com/eden-w2w/lib-modules/modules/id_generator"
	"github.com/eden-w2w/lib-modules/modules/order"
	"github.com/eden-w2w/lib-modules/modules/user"
	"github.com/eden-w2w/srv-cmop/internal/global"
	"github.com/sirupsen/logrus"
)

func init() {
	app := application.NewApplication(runner, true,
		application.WithConfig(&global.Config),
		application.WithConfig(&databases.Config))

	app.Start()
}

func runner(ctx *context.WaitStopContext) error {
	logrus.SetLevel(global.Config.LogLevel)
	id_generator.GetGenerator().Init(global.Config.SnowflakeConfig)
	admins.GetController().Init(global.Config.MasterDB, global.Config.PasswordSalt, global.Config.TokenExpired)
	user.GetController().Init(global.Config.MasterDB)
	goods.GetController().Init(global.Config.MasterDB)
	order.GetController().Init(global.Config.MasterDB, global.Config.OrderExpireIn, events.NewOrderEvent())

	return nil
}
