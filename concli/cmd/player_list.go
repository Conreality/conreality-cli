/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"context"
	"fmt"
	"strconv"

	"github.com/conreality/conreality.go/sdk"
	"github.com/spf13/cobra"
)

// PlayerListCmd describes and implements the `concli player list` command
var PlayerListCmd = &cobra.Command{
	Use:   "list [UNIT-NAME]",
	Short: "List all players on the team or a unit",
	Long:  "Conreality Command-Line Interface (CLI): List Players",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

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

		unitID := 0
		if len(args) > 0 {
			unitID, err = strconv.Atoi(args[0]) // TODO: support unit names as well
			if err != nil {
				panic(err)
			}
		}

		players, err := session.ListPlayers(ctx, sdk.UnitID(unitID))
		if err != nil {
			panic(err)
		}

		for _, player := range players {
			if debug {
				fmt.Printf("%d\n", player.ID)
			} else if verbose {
				fmt.Printf("Player #%d\n", player.ID) // TODO
			} else {
				fmt.Printf("#%d\n", player.ID) // TODO
			}
		}
	},
}

func init() {
	PlayerCmd.AddCommand(PlayerListCmd)
}
