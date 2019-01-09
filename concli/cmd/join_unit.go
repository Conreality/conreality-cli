/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"time"

	api "github.com/conreality/conreality.go/sdk/client"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

// JoinUnitCmd describes and implements the `concli join-unit` command
var JoinUnitCmd = &cobra.Command{
	Use:   "join-unit UNIT-NAME",
	Short: "Join a unit",
	Long:  `This is the command-line interface (CLI) for Conreality.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		client, err := api.Connect(masterURL)
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
	RootCmd.AddCommand(JoinUnitCmd)
}