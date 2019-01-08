/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// PingCmd describes and implements the `concli ping` command
var PingCmd = &cobra.Command{
	Use:   "ping",
	Short: "TODO",
	Long:  `This is the command-line interface (CLI) for Conreality.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ping has not been implemented yet") // TODO
		os.Exit(1)
	},
}

func init() {
	RootCmd.AddCommand(PingCmd)
}
