package commands

import (
	"context"
	"fmt"
	"github.com/cortezaproject/corteza-server/compose/types"
	"github.com/cortezaproject/corteza-server/pkg/id"
	"github.com/cortezaproject/corteza-server/seeder"
	"strconv"

	"github.com/cortezaproject/corteza-server/pkg/cli"
	"github.com/spf13/cobra"
)

func Seeder(app serviceInitializer) *cobra.Command {
	// Fake data generation commands.
	cmd := &cobra.Command{
		Use:   "seeder",
		Short: "Seeds fake data",
	}

	// Create users
	createUsersCmd := &cobra.Command{
		Use:   "users create",
		Short: "Create user",

		PreRunE: commandPreRunInitService(app),
		Run: func(cmd *cobra.Command, args []string) {
			var (
				ctx = context.Background()

				limitFlag = cmd.Flags().Lookup("limit").Value.String()

				limit int
				err   error
			)

			limit, err = strconv.Atoi(limitFlag)
			cli.HandleError(err)

			dataGen := seeder.Seeder(ctx, seeder.DefaultStore, seeder.Faker())

			userIDs, err := dataGen.CreateUser(seeder.Params{Limit: limit})
			cli.HandleError(err)

			fmt.Fprintf(
				cmd.OutOrStdout(),
				"                     Created    %d    users\n",
				len(userIDs),
			)
		},
	}

	createUsersCmd.Flags().IntP("limit", "l", 1, "How many users to be created")

	// Clear users
	deleteAllUsersCmd := &cobra.Command{
		Use:   "users delete all",
		Short: "Delete all user",

		PreRunE: commandPreRunInitService(app),
		Run: func(cmd *cobra.Command, args []string) {
			var (
				ctx = context.Background()
				err error
			)

			dataGen := seeder.Seeder(ctx, seeder.DefaultStore, seeder.Faker())

			err = dataGen.DeleteAllUser()
			cli.HandleError(err)

			fmt.Fprintf(
				cmd.OutOrStdout(),
				"                     Deleted    all    users\n",
			)
		},
	}

	// Create records
	createRecordsCmd := &cobra.Command{
		Use:   "records create",
		Short: "Create record",

		PreRunE: commandPreRunInitService(app),
		Run: func(cmd *cobra.Command, args []string) {
			var (
				ctx = context.Background()

				limitFlag = cmd.Flags().Lookup("limit").Value.String()

				limit int
				err   error
			)

			limit, err = strconv.Atoi(limitFlag)
			cli.HandleError(err)

			dataGen := seeder.Seeder(ctx, seeder.DefaultStore, seeder.Faker())

			recordIDs, err := dataGen.CreateRecord(id.Next(), seeder.Params{Limit: limit})
			cli.HandleError(err)

			fmt.Fprintf(
				cmd.OutOrStdout(),
				"                     Created    %d    records\n",
				len(recordIDs),
			)
		},
	}

	createRecordsCmd.Flags().IntP("limit", "l", 1, "How many records to be created")

	// Delete records
	deleteAllRecordsCmd := &cobra.Command{
		Use:   "record delete all",
		Short: "Delete all record",

		PreRunE: commandPreRunInitService(app),
		Run: func(cmd *cobra.Command, args []string) {
			var (
				ctx = context.Background()
				err error
			)

			dataGen := seeder.Seeder(ctx, seeder.DefaultStore, seeder.Faker())

			err = dataGen.DeleteAllRecord(&types.Module{})
			cli.HandleError(err)

			fmt.Fprintf(
				cmd.OutOrStdout(),
				"                     Deleted    all    records\n",
			)
		},
	}

	// Clear users
	deleteAllCmd := &cobra.Command{
		Use:   "delete all",
		Short: "Delete all fake data",

		PreRunE: commandPreRunInitService(app),
		Run: func(cmd *cobra.Command, args []string) {
			var (
				ctx = context.Background()
				err error
			)

			dataGen := seeder.Seeder(ctx, seeder.DefaultStore, seeder.Faker())

			err = dataGen.DeleteAll(&types.Module{})
			cli.HandleError(err)

			fmt.Fprintf(
				cmd.OutOrStdout(),
				"                     Deleted    all    fake    data\n",
			)
		},
	}

	cmd.AddCommand(
		createUsersCmd,
		deleteAllUsersCmd,
		createRecordsCmd,
		deleteAllRecordsCmd,
		deleteAllCmd,
	)

	return cmd
}
