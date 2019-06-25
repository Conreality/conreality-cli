/* This is free and unencumbered software released into the public domain. */

// conmon subscribes to events in an ongoing game and prints them out as they happen.
package main

import (
	"context"
	"fmt"

	"github.com/conreality/conreality.go/sdk"
)

func main() {
	ctx := context.Background()

	games, err := sdk.DiscoverGames(ctx)
	if err != nil {
		panic(err)
	}

	for game := range games {
		fmt.Printf("%s: Connecting to grpc://%s:%d...\n", game.Name, game.Host, game.Port)
		go monitorGame(ctx, game)
	}

	<-ctx.Done()
}

func monitorGame(ctx context.Context, gameEndpoint *sdk.GameEndpoint) {
	gameConnection, err := sdk.ConnectToGame(ctx, gameEndpoint)
	if err != nil {
		panic(err)
	}

	events, err := gameConnection.ReceiveEvents(ctx, 0)
	if err != nil {
		panic(err)
	}

	for event := range events {
		fmt.Printf("%s: Event #%d\n", gameEndpoint.Name, event.ID) // TODO
	}
}
