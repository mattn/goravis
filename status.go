package main

import (
	"fmt"
	"github.com/alecthomas/kingpin"
)

var statusCommand = kingpin.Command("status", "Checks status of the latest build.").Action(func(ctx *kingpin.ParseContext) error {
	err := auth()
	if err != nil {
		return err
	}

	s := slug()
	repo, resp, err := client.Repositories.GetFromSlug(s)
	if err != nil {
		return err
	}
	resp.Body.Close()

	fmt.Printf("build #%s %s\n", repo.LastBuildNumber, repo.LastBuildState)
	return nil
})
