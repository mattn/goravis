package main

import (
	"fmt"
	"github.com/alecthomas/kingpin"
)

var branchesCommand = kingpin.Command("branches", "displays the most recent build for each branch").Action(func(ctx *kingpin.ParseContext) error {
	err := auth()
	if err != nil {
		return err
	}

	s := slug(ctx)
	branches, _, err := client.Branches.ListFromRepository(s)
	if err != nil {
		return err
	}

	for _, branch := range branches {
		b, _, err := client.Commits.GetFromBuild(branch.Id)
		if err != nil {
			return err
		}

		fmt.Printf("%s %s %s\n", b.Branch, branch.Number, branch.State)
	}
	return nil
})
var branchesRepoFlag = branchesCommand.Flag("repo", "repository").Short('r').String()
