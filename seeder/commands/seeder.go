package commands

import (
	"context"
	"fmt"
	"github.com/cortezaproject/corteza-server/compose/types"
	"github.com/cortezaproject/corteza-server/seeder"
	"strconv"

	"github.com/cortezaproject/corteza-server/pkg/cli"
	"github.com/spf13/cobra"
)

func Seeder(app serviceInitializer) *cobra.Command {
	// Fake data generation commands.
	cmd := &cobra.Command{
		Use:   "faker",
		Short: "Fake data generation",
	}

	// Create users
	createUserCmd := &cobra.Command{
		Use:   "createUser",
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
				"                     Created    %d    user\n",
				len(userIDs),
			)
		},
	}

	createUserCmd.Flags().IntP("limit", "l", 1, "How many users to be created")

	// Clear users
	deleteAllUserCmd := &cobra.Command{
		Use:   "deleteAllUser",
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
				"                     Deleted    all    user\n",
			)
		},
	}

	// Create records
	createRecordCmd := &cobra.Command{
		Use:   "createRecord",
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

			recordIDs, err := dataGen.CreateRecord(&types.Module{}, seeder.Params{Limit: limit})
			cli.HandleError(err)

			fmt.Fprintf(
				cmd.OutOrStdout(),
				"                     Created    %d    record\n",
				len(recordIDs),
			)
		},
	}

	createRecordCmd.Flags().IntP("limit", "l", 1, "How many records to be created")

	// Delete records
	deleteAllRecordCmd := &cobra.Command{
		Use:   "deleteAllRecord",
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
				"                     Deleted    all    record\n",
			)
		},
	}

	// Clear users
	deleteAllCmd := &cobra.Command{
		Use:   "deleteAll",
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
		createUserCmd,
		deleteAllUserCmd,
		createRecordCmd,
		deleteAllRecordCmd,
		deleteAllCmd,
	)

	return cmd
}
