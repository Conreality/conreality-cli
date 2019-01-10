/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"github.com/spf13/cobra"
)

// DebugCmd describes and implements the `concli debug` command
var DebugCmd = &cobra.Command{
	Use:   "debug",
	Short: "TODO",
	Long:  `This is the command-line interface (CLI) for Conreality.`,
	Args:  cobra.NoArgs,
}

func init() {
	RootCmd.AddCommand(DebugCmd)
}
