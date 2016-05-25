package main

import (
	"github.com/alecthomas/kingpin"
)

var logoutCommand = kingpin.Command("logout", "deletes the stored API token").Action(func(ctx *kingpin.ParseContext) error {
	ep := config.EndPoints["https://api.travis-ci.org/"]
	ep.AccessToken = ""
	return config.Save()
})
