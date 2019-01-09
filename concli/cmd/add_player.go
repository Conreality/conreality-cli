/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"context"
	"fmt"

	"github.com/conreality/conreality.go/sdk"
	"github.com/spf13/cobra"
)

// AddPlayerCmd describes and implements the `concli add-player` command
var AddPlayerCmd = &cobra.Command{
	Use:   "add-player PLAYER-NAME",
	Short: "Add a new player to the game",
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

		player, err := session.AddPlayer(ctx, args[0])
		if err != nil {
			panic(err)
		}

		if verbose {
			fmt.Printf("%d\n", player.ID)
		}
	},
}

func init() {
	RootCmd.AddCommand(AddPlayerCmd)
}
