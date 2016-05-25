package main

import (
	"fmt"

	"github.com/alecthomas/kingpin"
)

var envCommand = kingpin.Command("env", "show or modify build environment variables")
var envListCommand = envCommand.Command("list", "").Action(func(ctx *kingpin.ParseContext) error {
	err := auth()
	if err != nil {
		return err
	}

	s := slug()
	repo, _, err := client.Repositories.GetFromSlug(s)
	if err != nil {
		return err
	}

	req, err := client.NewRequest("GET", "/settings/env_vars/?repository_id="+fmt.Sprint(repo.Id), nil, nil)
	if err != nil {
		return err
	}
	var envVarsRes struct {
		EnvVers []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"env_vars"`
	}
	_, err = client.Do(req, &envVarsRes)
	if err != nil {
		return err
	}
	for _, env_ver := range envVarsRes.EnvVers {
		fmt.Println(env_ver)
	}
	return nil
})
var envSetCommand = envCommand.Command("set", "").Action(func(ctx *kingpin.ParseContext) error {
	if len(ctx.Elements) != 4 {
		kingpin.Usage()
		return nil
	}
	err := auth()
	if err != nil {
		return err
	}

	s := slug()
	repo, _, err := client.Repositories.GetFromSlug(s)
	if err != nil {
		return err
	}

	req, err := client.NewRequest("POST", "/settings/env_vars/?repository_id="+fmt.Sprint(repo.Id), struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}{
		Name:  *ctx.Elements[2].Value,
		Value: *ctx.Elements[3].Value,
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
var envSetArg = envSetCommand.Arg("env", "env to set").Strings()
