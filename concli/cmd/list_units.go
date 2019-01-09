/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"context"
	"fmt"
	"strconv"

	"github.com/conreality/conreality.go/sdk"
	"github.com/spf13/cobra"
)

// ListUnitsCmd describes and implements the `concli list-units` command
var ListUnitsCmd = &cobra.Command{
	Use:   "list-units [UNIT-NAME]",
	Short: "List all units on the team or a parent unit",
	Long:  `This is the command-line interface (CLI) for Conreality.`,
	Args:  cobra.MaximumNArgs(1),
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

		unitID := 0
		if len(args) > 0 {
			unitID, err = strconv.Atoi(args[0]) // TODO: support unit names as well
			if err != nil {
				panic(err)
			}
		}

		units, err := session.ListUnits(ctx, sdk.UnitID(unitID))
		if err != nil {
			panic(err)
		}

		for _, unit := range units {
			if debug {
				fmt.Printf("%d\n", unit.ID)
			} else if verbose {
				fmt.Printf("%d\n", unit.ID) // TODO
			} else {
				fmt.Printf("%d\n", unit.ID) // TODO
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(ListUnitsCmd)
}
