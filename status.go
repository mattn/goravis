package main

import (
	"fmt"

	"github.com/alecthomas/kingpin"
)

var (
	statusCommand  = kingpin.Command("status", "checks status of the latest build")
	statusRepoFlag = statusCommand.Flag("repo", "repository").Short('r').String()
)

func init() {
	statusCommand.Action(statusAction)
}

func statusAction(ctx *kingpin.ParseContext) error {
	err := auth()
	if err != nil {
		return err
	}

	s := slug(statusRepoFlag)
	repo, _, err := client.Repositories.GetFromSlug(s)
	if err != nil {
		return err
	}
	if repo.LastBuildId == 0 {
		fatal("no build yet for "+s, nil)
	}

	fmt.Printf("build #%s %s\n", repo.LastBuildNumber, repo.LastBuildState)
	return nil
}
