/* This is free and unencumbered software released into the public domain. */

// consim simulates an agent based on a script.
package main

import (
	"fmt"
	"os"

	"github.com/conreality/conreality.go/sdk"
	"github.com/conreality/conreality-cli/consim/gdkvm"
	"github.com/spf13/cobra"
)

var debug bool
var verbose bool

// RootCmd describes the `consim` command
var RootCmd = &cobra.Command{
	Use:     "consim SCRIPT...",
	Short:   "Conreality simulator (consim)",
	Version: sdk.Version,
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		sim := gdkvm.NewSimulator()
		for _, scriptPath := range args {
			if verbose || debug {
				fmt.Printf("Executing %s...\n", scriptPath)
			}
			if err := sim.EvalFile(scriptPath); err != nil {
				panic(err)
			}
		}
		sim.Dump()
	},
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Enable debugging")
	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Be verbose")
	RootCmd.SetVersionTemplate("consim {{printf \"%s\" .Version}}\n")
}

func main() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
