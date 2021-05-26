package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jessevdk/go-flags"
)

type Options struct {
	Dir string `short:"d" long:"directory" description:"Source directory to scan for md files" default:"./"`
}

var options Options

var parser = flags.NewParser(&options, flags.Default)

func main() {
	if _, err := parser.Parse(); err != nil {
		switch flagsErr := err.(type) {
		case flags.ErrorType:
			if flagsErr == flags.ErrHelp {
				os.Exit(0)
			}
			os.Exit(1)
		default:

			os.Exit(1)
		}
	}

	// We've got our options, let's get the show on the road
	fmt.Printf("Scanning %v for .md files\n", options.Dir)
	if _, err := os.Stat(options.Dir); os.IsNotExist(err) {
		log.Fatalf("%v does not exist", options.Dir)
	}
	Todo(options.Dir)

}
