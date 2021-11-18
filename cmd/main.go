package main

import (
	"github.com/eden-framework/context"
	"github.com/eden-framework/eden-framework/pkg/application"
	"github.com/eden-framework/sqlx/migration"
	"github.com/eden-w2w/lib-modules/modules/admins"
	"github.com/eden-w2w/lib-modules/modules/booking_flow"
	"github.com/eden-w2w/lib-modules/modules/events"
	"github.com/eden-w2w/lib-modules/modules/goods"
	"github.com/eden-w2w/lib-modules/modules/id_generator"
	"github.com/eden-w2w/lib-modules/modules/order"
	"github.com/eden-w2w/lib-modules/modules/payment_flow"
	"github.com/eden-w2w/lib-modules/modules/promotion_flow"
	"github.com/eden-w2w/lib-modules/modules/refund_flow"
	"github.com/eden-w2w/lib-modules/modules/settings"
	"github.com/eden-w2w/lib-modules/modules/settlement_flow"
	"github.com/eden-w2w/lib-modules/modules/task_flow"
	"github.com/eden-w2w/lib-modules/modules/user"
	"github.com/eden-w2w/lib-modules/modules/wechat"
	"github.com/eden-w2w/lib-modules/pkg/cron"
	"github.com/eden-w2w/srv-cmop/internal/tasks/cancel_orders"
	"github.com/eden-w2w/srv-cmop/internal/tasks/fetch_wechat_payment"
	"github.com/eden-w2w/srv-cmop/internal/tasks/fetch_wechat_refund"
	"github.com/eden-w2w/srv-cmop/internal/tasks/reconciliation"
	"github.com/eden-w2w/srv-cmop/internal/tasks/settlement"
	"github.com/eden-w2w/srv-cmop/internal/tasks/update_booking_flow"
	"github.com/eden-w2w/srv-cmop/pkg/uploader"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/eden-w2w/lib-modules/databases"
	"github.com/eden-w2w/srv-cmop/internal/global"
	"github.com/eden-w2w/srv-cmop/internal/routers"
)

var cmdMigrationDryRun bool

func main() {
	app := application.NewApplication(
		runner, false,
		application.WithConfig(&global.Config),
		application.WithConfig(&databases.Config),
	)

	cmdMigrate := &cobra.Command{
		Use: "migrate",
		Run: func(cmd *cobra.Command, args []string) {
			migrate(
				&migration.MigrationOpts{
					DryRun: cmdMigrationDryRun,
				},
			)
		},
	}
	cmdMigrate.Flags().BoolVarP(&cmdMigrationDryRun, "dry", "d", false, "migrate --dry")
	app.AddCommand(cmdMigrate)

	app.Start()
}

func runner(ctx *context.WaitStopContext) error {
	initModules()
	initTask()

	cron.GetManager().Start()
	go global.Config.GRPCServer.Serve(ctx, routers.Router)
	return global.Config.HTTPServer.Serve(ctx, routers.Router)
}

func migrate(opts *migration.MigrationOpts) {
	if err := migration.Migrate(global.Config.MasterDB, opts); err != nil {
		panic(err)
	}
}

func initModules() {
	logrus.SetLevel(global.Config.LogLevel)
	id_generator.GetGenerator().Init(global.Config.SnowflakeConfig)
	admins.GetController().Init(global.Config.MasterDB, global.Config.PasswordSalt, global.Config.TokenExpired)
	user.GetController().Init(global.Config.MasterDB)
	goods.GetController().Init(global.Config.MasterDB)
	order.GetController().Init(
		global.Config.MasterDB,
		global.Config.OrderExpireIn,
		events.NewOrderEvent(
			global.Config.Wechat,
			goods.GetController().LockInventory,
			goods.GetController().UnlockInventory,
		),
	)
	payment_flow.GetController().Init(global.Config.MasterDB, 0)
	promotion_flow.GetController().Init(global.Config.MasterDB)
	uploader.GetManager().Init(
		global.Config.Uploader.Type,
		global.Config.Uploader.Endpoint,
		global.Config.Uploader.AccessKey,
		global.Config.Uploader.AccessSecret,
		global.Config.Uploader.BucketName,
	)
	settlement_flow.GetController().Init(global.Config.MasterDB, global.Config.SettlementConfig)
	task_flow.GetController().Init(global.Config.MasterDB)
	wechat.GetController().Init(global.Config.Wechat)
	refund_flow.GetController().Init(global.Config.MasterDB)
	settings.GetController().Init(global.Config.MasterDB)
	booking_flow.GetController().Init(global.Config.MasterDB)
}

func initTask() {
	if _, err := cron.GetManager().AddFunc(
		global.Config.SettlementConfig.ToSettlementCronRule(),
		settlement.TaskSettlement,
	); err != nil {
		panic(err)
	}
	if _, err := cron.GetManager().AddFunc(
		global.Config.CancelExpiredOrderTask,
		cancel_orders.TaskCancelOrders(goods.GetController().UnlockInventory),
	); err != nil {
		panic(err)
	}
	if _, err := cron.GetManager().AddFunc(
		global.Config.Wechat.FetchWechatPaymentStatusTask,
		fetch_wechat_payment.TaskFetchWechatPaymentStatus,
	); err != nil {
		panic(err)
	}
	if _, err := cron.GetManager().AddFunc(
		global.Config.ReconciliationTask,
		reconciliation.TaskReconciliation,
	); err != nil {
		panic(err)
	}
	if _, err := cron.GetManager().AddFunc(
		global.Config.Wechat.FetchWechatRefundTask,
		fetch_wechat_refund.TaskFetchWechatRefundStatus,
	); err != nil {
		panic(err)
	}
	if _, err := cron.GetManager().AddFunc(
		global.Config.BookingFlowsTask,
		update_booking_flow.TaskUpdateBookingFlow,
	); err != nil {
		panic(err)
	}
}
