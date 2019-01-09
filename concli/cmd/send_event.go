/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"context"

	api "github.com/conreality/conreality.go/sdk/client"
	"github.com/spf13/cobra"
)

// SendEventCmd describes and implements the `concli send-event` command
var SendEventCmd = &cobra.Command{
	Use:   "send-event PREDICATE-ID SUBJECT-NAME OBJECT-NAME",
	Short: "Send a broadcast event",
	Long:  `This is the command-line interface (CLI) for Conreality.`,
	Args:  cobra.ExactArgs(3),
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
	RootCmd.AddCommand(SendEventCmd)
}
