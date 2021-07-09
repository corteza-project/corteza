package commands

import (
	"context"
	"fmt"
	"github.com/cortezaproject/corteza-server/compose/types"
	"github.com/cortezaproject/corteza-server/pkg/cli"
	"github.com/cortezaproject/corteza-server/seeder"

	"github.com/spf13/cobra"
)

func Seeder(ctx context.Context, app serviceInitializer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "seeder",
		Short: "Seeds fake data",
		Long:  "Generates fake data for user and records",
	}

	cmd.AddCommand(
		users(ctx, app),
		records(ctx, app),
		deleteAll(ctx, app),
	)

	return cmd
}

func users(ctx context.Context, app serviceInitializer) (cmd *cobra.Command) {
	var (
		limit int
	)
	cmd = &cobra.Command{
		Use:     "users [action]",
		Short:   "Seed users",
		Args:    cobra.MinimumNArgs(1),
		PreRunE: commandPreRunInitService(app),
		Run: func(cmd *cobra.Command, args []string) {
			var (
				seed = seeder.Seeder(ctx, seeder.DefaultStore, seeder.Faker())
				err  error
			)

			action := args[0]
			switch action {
			case "create":
				userIDs, err := seed.CreateUser(seeder.Params{Limit: limit})
				cli.HandleError(err)

				fmt.Fprintf(
					cmd.OutOrStdout(),
					"                     Created    %d    users\n",
					len(userIDs),
				)

				break
			case "delete":
				err = seed.DeleteAllUser()
				cli.HandleError(err)

				fmt.Fprintf(
					cmd.OutOrStdout(),
					"                     Deleted    all    users\n",
				)

				break
			}
		},
	}

	cmd.Flags().IntVarP(&limit, "limit", "l", 1, "How many users to be created")

	return cmd
}

func records(ctx context.Context, app serviceInitializer) (cmd *cobra.Command) {
	var (
		namespaceID int
		moduleID    int

		params = seeder.RecordParams{}
	)
	cmd = &cobra.Command{
		Use:     "records [action]",
		Short:   "Seed records",
		Args:    cobra.MinimumNArgs(1),
		PreRunE: commandPreRunInitService(app),
		Run: func(cmd *cobra.Command, args []string) {
			var (
				seed = seeder.Seeder(ctx, seeder.DefaultStore, seeder.Faker())
				err  error
			)

			params.NamespaceID = uint64(namespaceID)
			params.ModuleID = uint64(moduleID)

			switch args[0] {
			case "create":
				recordIDs, err := seed.CreateRecord(params)
				cli.HandleError(err)

				fmt.Fprintf(
					cmd.OutOrStdout(),
					"                     Created    %d    records\n",
					len(recordIDs),
				)
				break
			case "delete":
				err = seed.DeleteAllRecord(&types.Module{})
				cli.HandleError(err)

				fmt.Fprintf(
					cmd.OutOrStdout(),
					"                     Deleted    all    records\n",
				)

				break
			}
		},
	}

	cmd.Flags().IntVarP(&params.Limit, "limit", "l", 1, "How many users to be created")
	cmd.Flags().IntVarP(&namespaceID, "namespace-id", "n", 0, "namespace id for recode creation")
	cmd.Flags().StringVarP(&params.NamespaceHandle, "namespace-handle", "h", "", "namespace handle for recode creation")
	cmd.Flags().StringVarP(&params.ModuleHandle, "module-handle", "m", "", "module handle for recode creation")
	cmd.Flags().IntVarP(&moduleID, "module-id", "j", 0, "module id for recode creation")

	return cmd
}

func deleteAll(ctx context.Context, app serviceInitializer) (cmd *cobra.Command) {
	var (
		namespaceID int
		moduleID    int

		params = seeder.RecordParams{}
	)
	cmd = &cobra.Command{
		Use:     "delete",
		Short:   "delete all data",
		PreRunE: commandPreRunInitService(app),
		Run: func(cmd *cobra.Command, args []string) {
			var (
				seed = seeder.Seeder(ctx, seeder.DefaultStore, seeder.Faker())
				err  error
			)

			params.NamespaceID = uint64(namespaceID)
			params.ModuleID = uint64(moduleID)

			err = seed.DeleteAll(&params)
			cli.HandleError(err)

			fmt.Fprintf(
				cmd.OutOrStdout(),
				"                     Deleted    all    fake    data\n",
			)
		},
	}

	cmd.Flags().IntVarP(&params.Limit, "limit", "l", 1, "How many users to be created")
	cmd.Flags().IntVarP(&namespaceID, "namespace-id", "n", 0, "namespace id for recode creation")
	cmd.Flags().StringVarP(&params.NamespaceHandle, "namespace-handle", "h", "", "namespace handle for recode creation")
	cmd.Flags().IntVarP(&moduleID, "module-id", "m", 0, "module id for recode creation")
	cmd.Flags().StringVarP(&params.ModuleHandle, "module-handle", "j", "", "module handle for recode creation")

	return cmd
}
