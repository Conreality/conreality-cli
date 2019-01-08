/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// HelloCmd describes and implements the `concli hello` command
var HelloCmd = &cobra.Command{
	Use:   "hello",
	Short: "TODO",
	Long:  `This is the command-line interface (CLI) for Conreality.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello has not been implemented yet") // TODO
		os.Exit(1)
	},
}

func init() {
	RootCmd.AddCommand(HelloCmd)
}
