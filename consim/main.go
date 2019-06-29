/* This is free and unencumbered software released into the public domain. */

// consim simulates an agent based on a script.
package main

import (
	"fmt"
	"os"

	"github.com/conreality/conreality-gdk/rt"
	"github.com/conreality/conreality.go/sdk"
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
		model := &rt.Model{
			Self:    &Agent{},
			Unit:    &Unit{},
			Theater: &Theater{},
			Game:    &Game{},
			Headset: &Headset{},
		}
		thread, err := rt.NewThread(model, true)
		if err != nil {
			panic(err)
		}
		defer thread.Destroy()

		for _, scriptPath := range args {
			if verbose || debug {
				fmt.Printf("Executing %s...\n", scriptPath)
			}

			if err := thread.EvalFile(scriptPath); err != nil {
				panic(err)
			}

			if debug {
				fmt.Printf("Invoking the %s hook...\n", "on_load")
			}
			if thread.HasFunction("on_load") {
				if err := thread.CallFunction("on_load"); err != nil {
					panic(err)
				}
			}

			if debug {
				fmt.Printf("Invoking the %s hook...\n", "on_activity")
			}
			if thread.HasFunction("on_activity") {
				if err := thread.CallFunction("on_activity", map[string]string{"activity": "running"}); err != nil { // FIXME
					panic(err)
				}
			}

			if debug {
				fmt.Printf("Invoking the %s hook...\n", "on_unload")
			}
			if thread.HasFunction("on_unload") {
				if err := thread.CallFunction("on_unload"); err != nil {
					panic(err)
				}
			}
		}

		if debug {
			thread.DumpStack()
		}
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
