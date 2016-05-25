package main

import (
	"fmt"

	"github.com/alecthomas/kingpin"
	"github.com/mattn/go-colorable"
)

var logsCommand = kingpin.Command("logs", "streams test logs").Action(func(ctx *kingpin.ParseContext) error {
	err := auth()
	if err != nil {
		return err
	}

	s := slug()
	builds, _, _, _, err := client.Builds.ListFromRepository(s, nil)
	if err != nil {
		return err
	}

	job, _, err := client.Jobs.Get(builds[0].JobIds[0])
	if err != nil {
		return err
	}

	log, _, err := client.Logs.Get(job.LogId)
	if err != nil {
		return err
	}

	fmt.Fprint(colorable.NewColorableStdout(), log.Body)
	return nil
})
