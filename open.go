package main

import (
	"fmt"

	"github.com/alecthomas/kingpin"
	"github.com/skratchdot/open-golang/open"
)

var (
	openCommand  = kingpin.Command("open", "opens a build or job in the browser")
	openRepoFlag = openCommand.Flag("repo", "repository").Short('r').String()
)

func init() {
	openCommand.Action(openAction)
}

func openAction(ctx *kingpin.ParseContext) error {
	u := fmt.Sprintf("https://travis-ci.org/%s", slug(openRepoFlag))
	return open.Run(u)
}
