package main

import (
	"fmt"
	"github.com/alecthomas/kingpin"
)

var statusCommand = kingpin.Command("status", "checks status of the latest build").Action(func(ctx *kingpin.ParseContext) error {
	err := auth()
	if err != nil {
		return err
	}

	s := slug()
	repo, _, err := client.Repositories.GetFromSlug(s)
	if err != nil {
		return err
	}

	fmt.Printf("build #%s %s\n", repo.LastBuildNumber, repo.LastBuildState)
	return nil
})
