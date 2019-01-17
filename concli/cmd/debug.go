/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"github.com/spf13/cobra"
)

// DebugCmd describes and implements the `concli debug` command
var DebugCmd = &cobra.Command{
	Use:   "debug",
	Short: "Debugging support",
	Long:  "Conreality Command-Line Interface (CLI): Debugging Support",
	Args:  cobra.NoArgs,
}

func init() {
	RootCmd.AddCommand(DebugCmd)
}
