/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	api "github.com/conreality/conreality.go/sdk/client"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

// StopGameCmd describes and implements the `concli stop-game` command
var StopGameCmd = &cobra.Command{
	Use:   "stop-game",
	Short: "Stop the game altogether",
	Long:  `This is the command-line interface (CLI) for Conreality.`,
	Args:  cobra.NoArgs,
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
	RootCmd.AddCommand(StopGameCmd)
}
