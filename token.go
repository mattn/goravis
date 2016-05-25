package main

import (
	"fmt"
	"github.com/alecthomas/kingpin"
)

var tokenCommand = kingpin.Command("token", "outputs the secret API token").Action(func(ctx *kingpin.ParseContext) error {
	err := auth()
	if err != nil {
		return err
	}

	fmt.Printf("Your access token is %s\n", token())
	return nil
})
