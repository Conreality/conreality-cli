/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"github.com/spf13/cobra"
)

// ConfigCmd describes and implements the `concli config` command
var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "TODO",
	Long:  `This is the command-line interface (CLI) for Conreality.`,
	Args:  cobra.NoArgs,
}

func init() {
	RootCmd.AddCommand(ConfigCmd)
}
