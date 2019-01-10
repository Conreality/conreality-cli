/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"github.com/spf13/cobra"
)

// GameCmd describes and implements the `concli game` command
var GameCmd = &cobra.Command{
	Use:   "game",
	Short: "TODO",
	Long:  `This is the command-line interface (CLI) for Conreality.`,
	Args:  cobra.NoArgs,
}

func init() {
	RootCmd.AddCommand(GameCmd)
}
