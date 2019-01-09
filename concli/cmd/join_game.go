/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"context"
	"fmt"

	"github.com/conreality/conreality.go/sdk"
	"github.com/spf13/cobra"
)

// JoinGameCmd describes and implements the `concli join-game` command
var JoinGameCmd = &cobra.Command{
	Use:   "join-game",
	Short: "Join the game",
	Long:  `This is the command-line interface (CLI) for Conreality.`,
	Args:  cobra.NoArgs,
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

		player, err := session.JoinGame(ctx)
		if err != nil {
			panic(err)
		}

		if verbose {
			fmt.Printf("%d\n", player.ID)
		}
	},
}

func init() {
	RootCmd.AddCommand(JoinGameCmd)
}
