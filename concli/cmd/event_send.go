/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"context"
	"fmt"
	"strconv"

	"github.com/conreality/conreality.go/sdk"
	"github.com/spf13/cobra"
)

// EventSendCmd describes and implements the `concli event send` command
var EventSendCmd = &cobra.Command{
	Use:   "send PREDICATE-ID SUBJECT-NAME OBJECT-NAME",
	Short: "Send a broadcast event",
	Long:  `This is the command-line interface (CLI) for Conreality.`,
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {

		client, err := sdk.Connect(masterURL)
		if err != nil {
			panic(err)
		}
		defer client.Disconnect()

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		session, err := client.Authenticate(ctx, playerNick)
		if err != nil {
			panic(err)
		}
		defer session.Close()

		predicate := args[0] // TODO: validate against controlled vocabulary

		subjectID, err := strconv.Atoi(args[1]) // TODO: support object names as well
		if err != nil {
			panic(err)
		}

		objectID, err := strconv.Atoi(args[2]) // TODO: support object names as well
		if err != nil {
			panic(err)
		}

		event, err := session.SendEvent(ctx, predicate, sdk.ObjectID(subjectID), sdk.ObjectID(objectID))
		if err != nil {
			panic(err)
		}

		if verbose {
			fmt.Printf("%d\n", event.ID)
		}
	},
}

func init() {
	EventCmd.AddCommand(EventSendCmd)
}
