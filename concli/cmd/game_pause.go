/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"context"

	"github.com/conreality/conreality.go/sdk"
	"github.com/spf13/cobra"
)

// GamePauseCmd describes and implements the `concli game pause` command
var GamePauseCmd = &cobra.Command{
	Use:   "pause",
	Short: "Pause the game temporarily",
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

		err = session.PauseGame(ctx, "") // TODO: --notice
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	GameCmd.AddCommand(GamePauseCmd)
}
