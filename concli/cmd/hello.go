/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"fmt"
	"time"

	api "github.com/conreality/conreality.go/sdk/client"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

// HelloCmd describes and implements the `concli hello` command
var HelloCmd = &cobra.Command{
	Use:   "hello",
	Short: "TODO", // TODO
	Long:  `This is the command-line interface (CLI) for Conreality.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {

		client, err := api.Connect(masterURL)
		if err != nil {
			panic(err)
		}
		defer client.Disconnect()

		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		masterVersion, err := client.Hello(ctx)
		if err != nil {
			panic(err)
		}
		fmt.Printf("version: %s\n", masterVersion)
	},
}

func init() {
	RootCmd.AddCommand(HelloCmd)
}
