/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"github.com/spf13/cobra"
)

// PlayerCmd describes and implements the `concli player` command
var PlayerCmd = &cobra.Command{
	Use:   "player",
	Short: "Player management",
	Long:  "Conreality Command-Line Interface (CLI): Player Management",
	Args:  cobra.NoArgs,
}

func init() {
	RootCmd.AddCommand(PlayerCmd)
}
