package main

import (
	"github.com/Ableton/go-travis"
	"github.com/alecthomas/kingpin"
)

const (
	VERSION = "0.0.1"
)

var (
	config = newConfig()
	client = travis.NewDefaultClient("")
)

func main() {
	client.Headers["User-Agent"] = "Travis/Goravis"
	kingpin.Parse()
}
