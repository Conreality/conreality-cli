/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"github.com/spf13/cobra"
)

// GameCmd describes and implements the `concli game` command
var GameCmd = &cobra.Command{
	Use:   "game",
	Short: "Game management [WIP]",
	Long:  "Conreality Command-Line Interface (CLI): Game Management",
	Args:  cobra.NoArgs,
}

func init() {
	RootCmd.AddCommand(GameCmd)
}
