package main

import (
	"fmt"
	"github.com/Ableton/go-travis"
	"github.com/alecthomas/kingpin"
)

var whatsUpCommand = kingpin.Command("whatsup", "Lists most recent builds.").Action(func(ctx *kingpin.ParseContext) error {
	err := auth()
	if err != nil {
		return err
	}

	u, resp, err := client.Users.GetAuthenticated()
	if err != nil {
		return err
	}
	resp.Body.Close()
	repos, resp, err := client.Repositories.Find(&travis.RepositoryListOptions{Member: u.Name})
	if err != nil {
		return err
	}
	resp.Body.Close()
	for _, repo := range repos {
		build, _, _, resp, err := client.Builds.Get(repo.LastBuildId)
		if err != nil {
			continue
		}
		resp.Body.Close()
		fmt.Printf("%s %s: #%s\n", repo.Slug, build.State, build.Number)
	}
	return nil
})
