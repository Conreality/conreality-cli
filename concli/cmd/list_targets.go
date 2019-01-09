/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"context"
	"fmt"
	"strconv"

	"github.com/conreality/conreality.go/sdk"
	"github.com/spf13/cobra"
)

// ListTargetsCmd describes and implements the `concli list-targets` command
var ListTargetsCmd = &cobra.Command{
	Use:   "list-targets [UNIT-NAME]",
	Short: "List all designated targets for the team or a unit",
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

		targets, err := session.ListTargets(ctx, sdk.UnitID(unitID))
		if err != nil {
			panic(err)
		}

		for _, target := range targets {
			if debug {
				fmt.Printf("%d\n", target.ID)
			} else if verbose {
				fmt.Printf("Target #%d\n", target.ID) // TODO
			} else {
				fmt.Printf("#%d\n", target.ID) // TODO
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(ListTargetsCmd)
}
