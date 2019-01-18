/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"context"
	"fmt"

	"github.com/conreality/conreality.go/sdk"
	"github.com/spf13/cobra"
)

// GameStateCmd describes and implements the `concli game state` command
var GameStateCmd = &cobra.Command{
	Use:     "state",
	Aliases: []string{"status"},
	Short:   "Get game state",
	Long:    "Conreality Command-Line Interface (CLI): Get Game State",
	Args:    cobra.NoArgs,
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

		state, err := session.GetGameState(ctx)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s\n", state)
	},
}

func init() {
	GameCmd.AddCommand(GameStateCmd)
}
