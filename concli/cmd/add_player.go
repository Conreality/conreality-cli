/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"context"

	api "github.com/conreality/conreality.go/sdk/client"
	"github.com/spf13/cobra"
)

// AddPlayerCmd describes and implements the `concli add-player` command
var AddPlayerCmd = &cobra.Command{
	Use:   "add-player PLAYER-NAME",
	Short: "Add a new player to the game",
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

		err = client.Ping(ctx) // TODO
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(AddPlayerCmd)
}
