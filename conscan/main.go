/* This is free and unencumbered software released into the public domain. */

// conscan attempts to discover ongoing games on the local network.
package main

import (
	"context"
	"fmt"

	"github.com/conreality/conreality.go/sdk"
)

func main() {
	games := make(chan *sdk.GameEndpoint)

	ctx := context.Background()
	err := sdk.DiscoverGames(ctx, games)
	if err != nil {
		panic(err)
	}

	go func(games <-chan *sdk.GameEndpoint) {
		for game := range games {
			fmt.Printf("Discovered '%s' at grpc://%s:%d...\n", game.Name, game.Host, game.Port)
		}
	}(games)

	<-ctx.Done()
}
