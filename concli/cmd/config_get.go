/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ConfigGetCmd describes and implements the `concli config get` command
var ConfigGetCmd = &cobra.Command{
	Use:   "get [VAR-NAME]",
	Short: "Show configuration variables",
	Long:  "Conreality Command-Line Interface (CLI): Get Configuration",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			for _, key := range viper.AllKeys() {
				fmt.Printf("%s=%s\n", key, viper.GetString(key))
			}
		} else {
			for _, key := range args {
				fmt.Printf("%s=%s\n", key, viper.GetString(key))
			}
		}
	},
}

func init() {
	ConfigCmd.AddCommand(ConfigGetCmd)
}
