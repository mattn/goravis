package main

import (
	"fmt"

	"github.com/alecthomas/kingpin"
	"github.com/skratchdot/open-golang/open"
)

var openCommand = kingpin.Command("open", "opens a build or job in the browser").Action(func(ctx *kingpin.ParseContext) error {
	u := fmt.Sprintf("https://github.com/%s", slug(ctx))
	return open.Run(u)
})
var openRepoFlag = openCommand.Flag("repo", "repository").Short('r').String()
