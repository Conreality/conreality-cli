/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	api "github.com/conreality/conreality.go/sdk/client"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

// UpdatePlayerCmd describes and implements the `concli update-player` command
var UpdatePlayerCmd = &cobra.Command{
	Use:   "update-player",
	Short: "Update player status",
	Long:  `This is the command-line interface (CLI) for Conreality.`,
	Args:  cobra.MinimumNArgs(0),
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
	RootCmd.AddCommand(UpdatePlayerCmd)
}
