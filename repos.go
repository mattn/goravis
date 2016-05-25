package main

import (
	"fmt"
	"github.com/alecthomas/kingpin"
)

var reposCommand = kingpin.Command("repos", "lists repositories the user has certain permissions on").Action(func(ctx *kingpin.ParseContext) error {
	err := auth()
	if err != nil {
		return err
	}

	repos, _, err := client.Repositories.Find(nil)
	if err != nil {
		return err
	}

	for _, repo := range repos {
		fmt.Println(repo.Slug)
		//{ "active" => repo.active?, "admin" => repo.admin?, "push" => repo.push?, "pull" => repo.pull? }
		fmt.Printf("Description: %s\n", repo.Description)
	}
	return nil
})
