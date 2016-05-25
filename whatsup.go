package main

import (
	"fmt"
	"github.com/Ableton/go-travis"
	"github.com/alecthomas/kingpin"
)

var whatsUpCommand = kingpin.Command("whatsup", "lists most recent builds").Action(func(ctx *kingpin.ParseContext) error {
	err := auth()
	if err != nil {
		return err
	}

	u, _, err := client.Users.GetAuthenticated()
	if err != nil {
		return err
	}

	repos, _, err := client.Repositories.Find(&travis.RepositoryListOptions{Member: u.Name})
	if err != nil {
		return err
	}

	for _, repo := range repos {
		build, _, _, _, err := client.Builds.Get(repo.LastBuildId)
		if err != nil {
			continue
		}

		fmt.Printf("%s %s: #%s\n", repo.Slug, build.State, build.Number)
	}
	return nil
})
