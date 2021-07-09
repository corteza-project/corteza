package app

import (
	"context"
	"fmt"

	authCommands "github.com/cortezaproject/corteza-server/auth/commands"
	federationCommands "github.com/cortezaproject/corteza-server/federation/commands"
	"github.com/cortezaproject/corteza-server/pkg/cli"
	"github.com/cortezaproject/corteza-server/pkg/logger"
	"github.com/cortezaproject/corteza-server/pkg/options"
	"github.com/cortezaproject/corteza-server/pkg/rbac"
	fakerCommands "github.com/cortezaproject/corteza-server/seeder/commands"
	"github.com/cortezaproject/corteza-server/store"
	systemCommands "github.com/cortezaproject/corteza-server/system/commands"
)

// InitCLI function initializes basic Corteza subsystems
// and sets-up the command line interface
func (app *CortezaApp) InitCLI() {
	var (
		ctx = cli.Context()

		// path to all environmental files (or locations with .env file)
		// filled from flag values
		envs []string
	)

	app.Command = cli.RootCommand(func() error {
		if len(envs) == 0 {
			envs = []string{"."}
		}

		if err := cli.LoadEnv(envs...); err != nil {
			return fmt.Errorf("failed to load environmental variables: %w", err)
		}

		logger.Init()

		// Environmental variables (from the env, files, see cli.LoadEnv) MUST be
		// loaded at this point!
		app.Opt = options.Init()

		return nil
	})

	app.Command.Flags().StringSliceVar(&envs, "env-file", nil,
		"Load environmental variables from files and directories containing .env file.\n"+
			"Values from loaded files DO NOT override existing variables from the environment.\n"+
			"This flag can be used multiple times, values are loaded from all provided locations.\n"+
			"If no paths are provided, corteza loads .env file from the current directory (equivalent to --env-file .)")

	serveCmd := cli.ServeCommand(func() (err error) {
		if err = app.Activate(ctx); err != nil {
			return
		}

		return app.Serve(ctx)
	})

	upgradeCmd := cli.UpgradeCommand(func() (err error) {
		if err = app.InitStore(ctx); err != nil {
			return
		}

		return
	})

	provisionCmd := cli.ProvisionCommand(func() (err error) {
		if err = app.Provision(ctx); err != nil {
			return
		}

		// Provisioning doesn't automatically reload rbac rules, so this is required
		rbac.Global().Reload(ctx)
		return
	})

	storeInit := func(ctx context.Context) (store.Storer, error) {
		err := app.InitStore(ctx)
		return app.Store, err
	}

	app.Command.AddCommand(
		systemCommands.Users(ctx, app),
		systemCommands.Roles(ctx, app),
		systemCommands.RBAC(ctx, storeInit),
		systemCommands.Sink(ctx, app),
		systemCommands.Settings(ctx, app),
		systemCommands.Import(ctx, storeInit),
		systemCommands.Export(ctx, storeInit),
		serveCmd,
		upgradeCmd,
		provisionCmd,
		authCommands.Command(ctx, app, storeInit),
		federationCommands.Sync(ctx, app),
		cli.EnvCommand(),
		cli.VersionCommand(),
		fakerCommands.Seeder(ctx, app),
	)

}

func (app *CortezaApp) Execute() error {
	return app.Command.Execute()
}
