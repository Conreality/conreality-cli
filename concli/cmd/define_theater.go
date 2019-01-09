/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"time"

	api "github.com/conreality/conreality.go/sdk/client"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

// DefineTheaterCmd describes and implements the `concli define-theater` command
var DefineTheaterCmd = &cobra.Command{
	Use:   "define-theater NW-GPS NE-GPS SW-GPS SE-GPS",
	Short: "Define the theater boundaries",
	Long:  `This is the command-line interface (CLI) for Conreality.`,
	Args:  cobra.ExactArgs(4),
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
	RootCmd.AddCommand(DefineTheaterCmd)
}
