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

// MessageSendCmd describes and implements the `concli message send` command
var MessageSendCmd = &cobra.Command{
	Use:   "send", // TODO: [FILE-PATH]
	Short: "Send a broadcast message",
	Long:  "Conreality Command-Line Interface (CLI): Send Message",
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
	MessageCmd.AddCommand(MessageSendCmd)
}
