/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"context"

	"github.com/conreality/conreality.go/sdk"
	"github.com/spf13/cobra"
)

// GameStartCmd describes and implements the `concli game start` command
var GameStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the game",
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

		err = session.StartGame(ctx, "") // TODO: --notice
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	GameCmd.AddCommand(GameStartCmd)
}
