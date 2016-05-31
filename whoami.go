package main

import (
	"fmt"

	"github.com/alecthomas/kingpin"
)

var whoamiCommand = kingpin.Command("whoami", "displays accounts and their subscription status").Action(func(ctx *kingpin.ParseContext) error {
	err := auth()
	if err != nil {
		return err
	}

	err = client.Authentication.UsingTravisToken(config.EndPoints["https://api.travis-ci.org/"].AccessToken)
	if err != nil {
		return err
	}
	u, _, err := client.Users.GetAuthenticated()
	if err != nil {
		return err
	}

	login := u.Login
	if login == "" {
		login = u.Name
	}
	fmt.Printf("You are %s (%s)\n", login, u.Name)
	return nil
})
