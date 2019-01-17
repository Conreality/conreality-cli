/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"github.com/spf13/cobra"
)

// ConfigCmd describes and implements the `concli config` command
var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Configuration management",
	Long:  "Conreality Command-Line Interface (CLI): Configuration Management",
	Args:  cobra.NoArgs,
}

func init() {
	RootCmd.AddCommand(ConfigCmd)
}
