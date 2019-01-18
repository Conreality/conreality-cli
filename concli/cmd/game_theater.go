/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"context"

	"github.com/conreality/conreality.go/sdk"
	"github.com/spf13/cobra"
)

// GameTheaterCmd describes and implements the `concli game theater` command
var GameTheaterCmd = &cobra.Command{
	Use:   "theater NW-GPS NE-GPS SW-GPS SE-GPS",
	Short: "Define the game theater boundaries [WIP]",
	Long:  "Conreality Command-Line Interface (CLI): Define Theater",
	Args:  cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {

		client, err := sdk.Connect(masterURL)
		if err != nil {
			panic(err)
		}
		defer client.Disconnect()

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		err = client.Ping(ctx) // TODO
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	GameCmd.AddCommand(GameTheaterCmd)
}
