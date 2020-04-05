package main

import (
	"log"

	"github.com/thepatrik/hellogrpc/pkg/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		log.Fatal(err)
	}
}
