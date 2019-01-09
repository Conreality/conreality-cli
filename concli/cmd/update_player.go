/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"context"

	"github.com/conreality/conreality.go/sdk"
	"github.com/spf13/cobra"
)

// UpdatePlayerCmd describes and implements the `concli update-player` command
var UpdatePlayerCmd = &cobra.Command{
	Use:   "update-player",
	Short: "Update player status",
	Long:  `This is the command-line interface (CLI) for Conreality.`,
	Args:  cobra.MinimumNArgs(0),
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

		err = session.UpdatePlayer(ctx, 42) // TODO: --heartbeat
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(UpdatePlayerCmd)
}
