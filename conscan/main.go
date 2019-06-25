/* This is free and unencumbered software released into the public domain. */

// conscan attempts to discover ongoing games on the local network.
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
		fmt.Printf("Discovered '%s' at grpc://%s:%d...\n", game.Name, game.Host, game.Port)
	}

	<-ctx.Done()
}
