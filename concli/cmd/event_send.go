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
	Long:  "Conreality Command-Line Interface (CLI): Send Event",
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

		subjectID, err := resolveEntityArg(ctx, session, args[1])
		if err != nil {
			panic(err)
		}

		objectID, err := resolveEntityArg(ctx, session, args[2])
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

func resolveEntityArg(ctx context.Context, session *sdk.Session, arg string) (sdk.EntityID, error) {
	entityID, err := strconv.Atoi(arg)
	if err != nil {
		entity, err := session.LookupEntityByName(ctx, arg)
		if err != nil {
			panic(err)
		}
		if entity == nil {
			return 0, fmt.Errorf("Unknown entity: %s", arg)
		}
		entityID = int(entity.ID)
	}
	return sdk.EntityID(entityID), nil
}

func init() {
	EventCmd.AddCommand(EventSendCmd)
}
