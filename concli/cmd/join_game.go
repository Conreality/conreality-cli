/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	api "github.com/conreality/conreality.go/sdk/client"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

// JoinGameCmd describes and implements the `concli join-game` command
var JoinGameCmd = &cobra.Command{
	Use:   "join-game [GAME-URL]",
	Short: "Join a game",
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

		_, err = client.Authenticate(ctx, playerNick)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(JoinGameCmd)
}
