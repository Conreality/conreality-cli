/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"context"
	"strconv"

	api "github.com/conreality/conreality.go/sdk/client"
	"github.com/conreality/conreality.go/sdk/model"
	"github.com/spf13/cobra"
)

// JoinUnitCmd describes and implements the `concli join-unit` command
var JoinUnitCmd = &cobra.Command{
	Use:   "join-unit UNIT-NAME",
	Short: "Join a unit",
	Long:  `This is the command-line interface (CLI) for Conreality.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		client, err := api.Connect(masterURL)
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

		err = session.JoinUnit(ctx, model.UnitID(unitID))
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(JoinUnitCmd)
}
