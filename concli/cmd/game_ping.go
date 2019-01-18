/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"context"

	"github.com/conreality/conreality.go/sdk"
	"github.com/spf13/cobra"
)

// GamePingCmd describes and implements the `concli game ping` command
var GamePingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Ping the game's master server",
	Long:  "Conreality Command-Line Interface (CLI): Ping Game",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {

		client, err := sdk.Connect(masterURL)
		if err != nil {
			panic(err)
		}
		defer client.Disconnect()

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		err = client.Ping(ctx)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	GameCmd.AddCommand(GamePingCmd)
}
