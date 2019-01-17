/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"github.com/spf13/cobra"
)

// TargetCmd describes and implements the `concli target` command
var TargetCmd = &cobra.Command{
	Use:   "target",
	Short: "Target designation [WIP]",
	Long:  "Conreality Command-Line Interface (CLI): Target Designation",
	Args:  cobra.NoArgs,
}

func init() {
	RootCmd.AddCommand(TargetCmd)
}
