/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"context"
	"fmt"

	"github.com/conreality/conreality.go/sdk"
	"github.com/spf13/cobra"
)

// UnitFormCmd describes and implements the `concli unit form` command
var UnitFormCmd = &cobra.Command{
	Use:   "form UNIT-NAME",
	Short: "Form a new unit",
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

		unit, err := session.FormUnit(ctx, args[0])
		if err != nil {
			panic(err)
		}

		if verbose {
			fmt.Printf("%d\n", unit.ID)
		}
	},
}

func init() {
	UnitCmd.AddCommand(UnitFormCmd)
}
