package main

import (
	"log"
	"os"

	"github.com/jessevdk/go-flags"
)

const (
	ExitCodeOK int = iota
	ExitCodeError
)

type Options struct {
	Dir string `short:"d" long:"directory" description:"Source directory to scan for md files" default:"./"`
}

var options Options

var parser = flags.NewParser(&options, flags.Default)

func main() {
	os.Exit(Run())
}

func Run() int {
	if _, err := parser.Parse(); err != nil {
		switch flagsErr := err.(type) {
		case flags.ErrorType:
			if flagsErr == flags.ErrHelp {
				return ExitCodeOK
			}
			return ExitCodeError
		default:
			return ExitCodeError
		}
	}

	// We've got our options, let's get the show on the road
	log.Printf("Scanning %v for .md files\n", options.Dir)
	if _, err := os.Stat(options.Dir); os.IsNotExist(err) {
		log.Printf("%v does not exist", options.Dir)
		return ExitCodeError
	}

	err := Todo(options.Dir)
	if err != nil {
		log.Printf("error: %v", err)
		return ExitCodeError
	}

	return ExitCodeOK
}


