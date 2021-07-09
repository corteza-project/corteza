package commands

import (
	"context"
	"fmt"
	"github.com/cortezaproject/corteza-server/seeder"
	"strconv"

	"github.com/cortezaproject/corteza-server/pkg/cli"
	"github.com/spf13/cobra"
)

func Faker(app serviceInitializer) *cobra.Command {
	// Fake data generation commands.
	cmd := &cobra.Command{
		Use:   "faker",
		Short: "Fake data generation",
	}

	// Create users
	createUserCmd := &cobra.Command{
		Use:   "addUser",
		Short: "Create users",

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

			dataGen := seeder.DataGen(ctx, seeder.DefaultStore, seeder.Faker())

			userIDs, err := dataGen.MakeMeSomeFakeUserPlease(seeder.GenOption{Limit: limit})
			cli.HandleError(err)

			fmt.Fprintf(
				cmd.OutOrStdout(),
				"                     Created    %d    users\n",
				len(userIDs),
			)
		},
	}

	createUserCmd.Flags().IntP("limit", "l", 1, "How many users to be created")

	// Clear users
	clearUserCmd := &cobra.Command{
		Use:   "clearUser",
		Short: "Clear fake users",

		PreRunE: commandPreRunInitService(app),
		Run: func(cmd *cobra.Command, args []string) {
			var (
				ctx = context.Background()
				err error
			)

			dataGen := seeder.DataGen(ctx, seeder.DefaultStore, seeder.Faker())

			err = dataGen.ClearFakeUsers()
			cli.HandleError(err)

			fmt.Fprintf(
				cmd.OutOrStdout(),
				"                     Cleared    all    fake    users\n",
			)
		},
	}

	// Create records
	createRecordCmd := &cobra.Command{
		Use:   "addRecords",
		Short: "Create users",

		PreRunE: commandPreRunInitService(app),
		Run: func(cmd *cobra.Command, args []string) {
			// var (
			// 	ctx = context.Background()
			//
			// 	limitFlag = cmd.Flags().Lookup("limit").Value.String()
			//
			// 	limit int
			// 	err   error
			// )
			//
			// limit, err = strconv.Atoi(limitFlag)
			// cli.HandleError(err)
			//
			// dataGen := seeder.DataGen(ctx, seeder.DefaultStore, seeder.Faker())
			//
			// // _, err := dataGen.MakeMeSomeFakeRecordPlease(seeder.GenOption{Limit: limit})
			// // cli.HandleError(err)
			//
			// fmt.Fprintf(
			// 	cmd.OutOrStdout(),
			// 	"                     Created    %d    records\n",
			// 	0,
			// )
		},
	}

	createRecordCmd.Flags().IntP("limit", "l", 1, "How many records to be created")

	// Clear records
	clearRecordCmd := &cobra.Command{
		Use:   "clearRecord",
		Short: "Clear fake record",

		PreRunE: commandPreRunInitService(app),
		Run: func(cmd *cobra.Command, args []string) {
			var (
				ctx = context.Background()
				err error
			)

			dataGen := seeder.DataGen(ctx, seeder.DefaultStore, seeder.Faker())

			err = dataGen.ClearFakeRecords()
			cli.HandleError(err)

			fmt.Fprintf(
				cmd.OutOrStdout(),
				"                     Cleared    all    fake    records\n",
			)
		},
	}

	// Clear users
	clearAllFakeDataCmd := &cobra.Command{
		Use:   "clearData",
		Short: "Clear all fake data",

		PreRunE: commandPreRunInitService(app),
		Run: func(cmd *cobra.Command, args []string) {
			var (
				ctx = context.Background()
				err error
			)

			dataGen := seeder.DataGen(ctx, seeder.DefaultStore, seeder.Faker())

			err = dataGen.ClearAllFakeData()
			cli.HandleError(err)

			fmt.Fprintf(
				cmd.OutOrStdout(),
				"                     Cleared    all    fake    data\n",
			)
		},
	}

	cmd.AddCommand(
		createUserCmd,
		clearUserCmd,
		createRecordCmd,
		clearRecordCmd,
		clearAllFakeDataCmd,
	)

	return cmd
}
