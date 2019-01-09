/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"context"

	"github.com/conreality/conreality.go/sdk"
	"github.com/spf13/cobra"
)

// StopGameCmd describes and implements the `concli stop-game` command
var StopGameCmd = &cobra.Command{
	Use:   "stop-game",
	Short: "Stop the game altogether",
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

		err = session.StopGame(ctx, "") // TODO: --notice
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(StopGameCmd)
}
