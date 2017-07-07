package main

import (
	"os"

	"github.com/cbguder/guac/commands"
	"github.com/jessevdk/go-flags"
)

func main() {
	_, err := flags.Parse(&commands.Opts)
	if err != nil {
		os.Exit(2)
	}
}
