/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"time"

	api "github.com/conreality/conreality.go/sdk/client"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

// DesignateTargetCmd describes and implements the `concli designate-target` command
var DesignateTargetCmd = &cobra.Command{
	Use:   "designate-target X-POS Y-POS [Z-POS]",
	Short: "Designate a target for the unit",
	Long:  `This is the command-line interface (CLI) for Conreality.`,
	Args:  cobra.RangeArgs(2, 3),
	Run: func(cmd *cobra.Command, args []string) {

		client, err := api.Connect(master)
		if err != nil {
			panic(err)
		}
		defer client.Disconnect()

		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		err = client.Ping(ctx) // TODO
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(DesignateTargetCmd)
}
