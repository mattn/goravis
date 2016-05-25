package main

import (
	"log"

	"github.com/Ableton/go-travis"
	"github.com/alecthomas/kingpin"
)

const (
	VERSION = "0.0.1"
)

var (
	config Config
	client = travis.NewDefaultClient("")
)

func main() {
	err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	kingpin.Parse()
}
