/* This is free and unencumbered software released into the public domain. */

// concli provides a command-line interface (CLI) for Conreality.
package main

import (
	"flag"
	"fmt"
	api "github.com/conreality/conreality.go/sdk/client"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"os"
	"time"
)

const Version = "0.0.0"

const master = "localhost:50051"

func usage() {
	cmd := os.Args[0]
	fmt.Fprintf(os.Stderr, "Usage: %s\n", cmd)
	// TODO
	os.Exit(2)
}

func help(argv0 string, args ...string) error {
	return nil
}

func hello(argv0 string, args ...string) error {
	client, err := api.Connect(master)
	if err != nil {
		return errors.Wrap(err, "connect failed")
	}
	defer client.Disconnect()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	masterVersion, err := client.Hello(ctx)
	if err != nil {
		return errors.Wrap(err, "hello failed")
	}
	fmt.Printf("version: %s\n", masterVersion)
	return nil
}

func ping(argv0 string, args ...string) error {
	client, err := api.Connect(master)
	if err != nil {
		return errors.Wrap(err, "connect failed")
	}
	defer client.Disconnect()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	err = client.Ping(ctx)
	if err != nil {
		return errors.Wrap(err, "ping failed")
	}
	return nil
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}

	argv0 := os.Args[0]
	args := os.Args[2:]

	var err error
	switch os.Args[1] {
	case "help":
		err = help(argv0, args...)
	case "hello":
		err = hello(argv0, args...)
	case "ping":
		err = ping(argv0, args...)
	default:
		usage()
	}

	if err != nil {
		if err != flag.ErrHelp {
			fmt.Fprintf(os.Stderr, "%s: %s\n", argv0, err)
			os.Exit(1)
		}
		os.Exit(2)
	}
}
