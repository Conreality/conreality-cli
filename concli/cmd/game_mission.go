/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"context"
	"io/ioutil"
	"os"

	"github.com/conreality/conreality.go/sdk"
	"github.com/spf13/cobra"
)

// GameMissionCmd describes and implements the `concli game mission` command
var GameMissionCmd = &cobra.Command{
	Use:   "mission",
	Short: "Define the game mission objective",
	Long:  "Conreality Command-Line Interface (CLI): Define Mission",
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
	GameCmd.AddCommand(GameMissionCmd)
}
