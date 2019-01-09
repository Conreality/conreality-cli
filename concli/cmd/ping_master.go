/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"context"

	"github.com/conreality/conreality.go/sdk"
	"github.com/spf13/cobra"
)

// PingMasterCmd describes and implements the `concli ping-master` command
var PingMasterCmd = &cobra.Command{
	Use:   "ping-master",
	Short: "Ping the master server",
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

		err = client.Ping(ctx)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(PingMasterCmd)
}
