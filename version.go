package main

import (
	"fmt"
	"github.com/alecthomas/kingpin"
)

var versionCommand = kingpin.Command("version", "outputs the client version").Action(func(ctx *kingpin.ParseContext) error {
	fmt.Println(VERSION)
	return nil
})
