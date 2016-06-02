package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/alecthomas/kingpin"
	"github.com/mattn/go-colorable"
)

var (
	logsCommand  = kingpin.Command("logs", "streams test logs")
	logsRepoFlag = logsCommand.Flag("repo", "repository").Short('r').String()
)

func init() {
	logsCommand.Action(logsAction)
}

func logsAction(ctx *kingpin.ParseContext) error {
	err := auth()
	if err != nil {
		return err
	}

	s := slug(logsRepoFlag)
	repo, _, err := client.Repositories.GetFromSlug(s)
	if err != nil {
		return err
	}
	if repo.LastBuildId == 0 {
		fatal("no build yet for "+s, nil)
	}

	_, jobs, _, _, err := client.Builds.Get(repo.LastBuildId)
	if err != nil {
		return err
	}

	u, err := client.BaseURL.Parse(fmt.Sprintf("/jobs/%d/log", jobs[0].Id))
	if err != nil {
		return err
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(colorable.NewColorableStdout(), resp.Body)
	return nil
}
