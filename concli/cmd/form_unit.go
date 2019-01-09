/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"context"
	"fmt"

	"github.com/conreality/conreality.go/sdk"
	"github.com/spf13/cobra"
)

// FormUnitCmd describes and implements the `concli form-unit` command
var FormUnitCmd = &cobra.Command{
	Use:   "form-unit UNIT-NAME",
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
	RootCmd.AddCommand(FormUnitCmd)
}
