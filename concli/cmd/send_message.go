/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/conreality/conreality.go/sdk"
	"github.com/spf13/cobra"
)

// SendMessageCmd describes and implements the `concli send-message` command
var SendMessageCmd = &cobra.Command{
	Use:   "send-message", // TODO: [FILE-PATH]
	Short: "Send a broadcast message",
	Long:  `This is the command-line interface (CLI) for Conreality.`,
	Args:  cobra.NoArgs, // TODO: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		bytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
		input := string(bytes)

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

		message, err := session.SendMessage(ctx, input)
		if err != nil {
			panic(err)
		}

		if verbose {
			fmt.Printf("%d\n", message.ID)
		}
	},
}

func init() {
	RootCmd.AddCommand(SendMessageCmd)
}
