/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"context"
	"fmt"

	"github.com/conreality/conreality.go/sdk"
	"github.com/spf13/cobra"
)

// DebugLookupCmd describes and implements the `concli debug lookup` command
var DebugLookupCmd = &cobra.Command{
	Use:   "lookup [ENTITY-NAME...]",
	Short: "Lookup an entity ID based on its name",
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

		entity, err := session.LookupEntityByName(ctx, args[0])
		if err != nil {
			panic(err)
		}

		if verbose && entity != nil {
			fmt.Printf("%d\n", entity.ID)
		}
	},
}

func init() {
	DebugCmd.AddCommand(DebugLookupCmd)
}
