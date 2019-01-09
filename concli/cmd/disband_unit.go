/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"context"
	"strconv"

	"github.com/conreality/conreality.go/sdk"
	"github.com/spf13/cobra"
)

// DisbandUnitCmd describes and implements the `concli disband-unit` command
var DisbandUnitCmd = &cobra.Command{
	Use:   "disband-unit UNIT-NAME",
	Short: "Disband a unit",
	Long:  `This is the command-line interface (CLI) for Conreality.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		client, err := sdk.Connect(masterURL)
		if err != nil {
			panic(err)
		}
		defer client.Disconnect()

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		session, err := client.Authenticate(ctx, playerNick)
		if err != nil {
			panic(err)
		}
		defer session.Close()

		unitID, err := strconv.Atoi(args[0]) // TODO: support unit names as well
		if err != nil {
			panic(err)
		}

		err = session.DisbandUnit(ctx, sdk.UnitID(unitID))
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(DisbandUnitCmd)
}
