/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// ConfigInitCmd describes and implements the `concli config init` command
var ConfigInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the $HOME/.conreality configuration directory",
	Long:  `This is the command-line interface (CLI) for Conreality.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {

		configDirPath := fmt.Sprintf("%s/.conreality", os.Getenv("HOME"))
		gamesDirPath := fmt.Sprintf("%s/games", configDirPath)

		// Create the $HOME/.conreality directory, if it doesn't exist:
		if err := os.Mkdir(configDirPath, 0700); err != nil && !os.IsExist(err) {
			panic(err)
		}

		// Set the correct permissions on the $HOME/.conreality directory:
		if err := os.Chmod(configDirPath, 0700); err != nil {
			panic(err)
		}

		// Create the $HOME/.conreality/games directory, if it doesn't exist:
		if err := os.Mkdir(gamesDirPath, 0700); err != nil && !os.IsExist(err) {
			panic(err)
		}

		// Set the correct permissions on the $HOME/.conreality/games directory:
		if err := os.Chmod(gamesDirPath, 0700); err != nil {
			panic(err)
		}
	},
}

func init() {
	ConfigCmd.AddCommand(ConfigInitCmd)
}
