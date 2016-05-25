package main

import (
	"github.com/alecthomas/kingpin"
)

var enableCommand = kingpin.Command("enable", "enables a project").Action(func(ctx *kingpin.ParseContext) error {
	err := auth()
	if err != nil {
		return err
	}

	s := slug(ctx)
	repo, _, err := client.Repositories.GetFromSlug(s)
	if err != nil {
		return err
	}

	type Hook struct {
		Id     uint `json:"id"`
		Active bool `json:"active"`
	}

	req, err := client.NewRequest("PUT", "/hooks/", struct {
		Hook `json:"hook"`
	}{
		Hook{repo.Id, true},
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
var enableRepoFlag = enableCommand.Flag("repo", "repository").Short('r').String()
