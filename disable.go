package main

import (
	"github.com/alecthomas/kingpin"
)

var disableCommand = kingpin.Command("disable", "disable a project").Action(func(ctx *kingpin.ParseContext) error {
	err := auth()
	if err != nil {
		return err
	}

	s := slug(ctx)
	repo, _, err := client.Repositories.GetFromSlug(s)
	if err != nil {
		return err
	}
	if repo.LastBuildId == 0 {
		fatal("no build yet for "+s, nil)
	}

	type Hook struct {
		Id     uint `json:"id"`
		Active bool `json:"active"`
	}

	req, err := client.NewRequest("PUT", "/hooks/", struct {
		Hook `json:"hook"`
	}{
		Hook{repo.Id, false},
	}, nil)
	if err != nil {
		return err
	}
	_, err = client.Do(req, nil)
	if err != nil {
		return err
	}
	return nil
})
var disableRepoFlag = disableCommand.Flag("repo", "repository").Short('r').String()
