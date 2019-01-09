/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"context"

	"github.com/conreality/conreality.go/sdk"
	"github.com/spf13/cobra"
)

// LeaveGameCmd describes and implements the `concli leave-game` command
var LeaveGameCmd = &cobra.Command{
	Use:   "leave-game",
	Short: "Leave the game",
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

		err = session.LeaveGame(ctx, "") // TODO: --reason
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(LeaveGameCmd)
}
