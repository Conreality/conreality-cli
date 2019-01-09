/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"context"
	"io/ioutil"
	"os"

	"github.com/conreality/conreality.go/sdk"
	"github.com/spf13/cobra"
)

// DefineMissionCmd describes and implements the `concli define-mission` command
var DefineMissionCmd = &cobra.Command{
	Use:   "define-mission",
	Short: "Define the mission objective",
	Long:  `This is the command-line interface (CLI) for Conreality.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, _ []string) {

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

		err = session.DefineMission(ctx, input)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(DefineMissionCmd)
}
