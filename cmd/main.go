package main

import (
	"github.com/eden-framework/context"
	"github.com/eden-framework/eden-framework/pkg/application"
	"github.com/eden-framework/sqlx/migration"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/eden-w2w/srv-cmop/internal/databases"
	"github.com/eden-w2w/srv-cmop/internal/global"
	"github.com/eden-w2w/srv-cmop/internal/routers"
)

var cmdMigrationDryRun bool

func main() {
	app := application.NewApplication(runner, false,
		application.WithConfig(&global.Config),
		application.WithConfig(&databases.Config))

	cmdMigrate := &cobra.Command{
		Use: "migrate",
		Run: func(cmd *cobra.Command, args []string) {
			migrate(&migration.MigrationOpts{
				DryRun: cmdMigrationDryRun,
			})
		},
	}
	cmdMigrate.Flags().BoolVarP(&cmdMigrationDryRun, "dry", "d", false, "migrate --dry")
	app.AddCommand(cmdMigrate)

	app.Start()
}

func runner(ctx *context.WaitStopContext) error {
	logrus.SetLevel(global.Config.LogLevel)
	go global.Config.GRPCServer.Serve(ctx, routers.Router)
	return global.Config.HTTPServer.Serve(ctx, routers.Router)
}

func migrate(opts *migration.MigrationOpts) {
	if err := migration.Migrate(global.Config.MasterDB, opts); err != nil {
		panic(err)
	}
}
