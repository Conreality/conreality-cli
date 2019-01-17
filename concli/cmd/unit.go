/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"github.com/spf13/cobra"
)

// UnitCmd describes and implements the `concli unit` command
var UnitCmd = &cobra.Command{
	Use:   "unit",
	Short: "Unit management [WIP]",
	Long:  "Conreality Command-Line Interface (CLI): Unit Management",
	Args:  cobra.NoArgs,
}

func init() {
	RootCmd.AddCommand(UnitCmd)
}
