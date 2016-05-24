package main

import (
	"fmt"

	"github.com/alecthomas/kingpin"
)

var whoamiCommand = kingpin.Command("whoami", "Displays accounts and their subscription status.").Action(func(ctx *kingpin.ParseContext) error {
	err := client.Authentication.UsingTravisToken(config.EndPoints["https://api.travis-ci.org/"].AccessToken)
	if err != nil {
		return err
	}
	u, resp, err := client.Users.GetAuthenticated()
	if err != nil {
		return err
	}
	resp.Body.Close()
	login := u.Login
	if login == "" {
		login = u.Name
	}
	fmt.Printf("You are %s (%s)\n", login, u.Name)
	return nil
})
