/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"github.com/spf13/cobra"
)

// EventCmd describes and implements the `concli event` command
var EventCmd = &cobra.Command{
	Use:   "event",
	Short: "Event sending [WIP]",
	Long:  "Conreality Command-Line Interface (CLI): Event Sending",
	Args:  cobra.NoArgs,
}

func init() {
	RootCmd.AddCommand(EventCmd)
}
