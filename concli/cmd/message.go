/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"github.com/spf13/cobra"
)

// MessageCmd describes and implements the `concli message` command
var MessageCmd = &cobra.Command{
	Use:   "message",
	Short: "TODO",
	Long:  `This is the command-line interface (CLI) for Conreality.`,
	Args:  cobra.NoArgs,
}

func init() {
	RootCmd.AddCommand(MessageCmd)
}
