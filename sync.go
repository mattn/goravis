package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/alecthomas/kingpin"
)

var syncCommand = kingpin.Command("sync", "triggers a new sync with GitHub").Action(func(ctx *kingpin.ParseContext) error {
	err := auth()
	if err != nil {
		return err
	}

	loop := true
	var wg sync.WaitGroup

	fmt.Print("Synchronizing: ")
	wg.Add(1)
	go func() {
		_, err = client.Users.Sync()
		loop = false
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		for loop {
			fmt.Print(".")
			time.Sleep(time.Second)
		}
		wg.Done()
	}()
	wg.Wait()
	return err
})
var syncRepoFlag = syncCommand.Flag("repo", "repository").Short('r').String()
