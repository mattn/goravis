package main

import (
	"fmt"

	"github.com/alecthomas/kingpin"
)

var (
	envCommand     = kingpin.Command("env", "show or modify build environment variables")
	envListCommand = envCommand.Command("list", "")
	envSetCommand  = envCommand.Command("set", "")
	envSetArg      = envSetCommand.Arg("env", "env to set").Strings()
	envRepoFlag    = envCommand.Flag("repo", "repository").Short('r').String()
)

func init() {
	envListCommand.Action(envListAction)
	envSetCommand.Action(envSetAction)
}

func envListAction(ctx *kingpin.ParseContext) error {
	err := auth()
	if err != nil {
		return err
	}

	s := slug(envRepoFlag)
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
	fmt.Println("# environment variables for " + s)
	for _, env_ver := range envVarsRes.EnvVers {
		fmt.Println(env_ver)
	}
	return nil
}

func envSetAction(ctx *kingpin.ParseContext) error {
	if len(ctx.Elements) != 4 {
		kingpin.Usage()
		return nil
	}
	err := auth()
	if err != nil {
		return err
	}

	s := slug(envRepoFlag)
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
}
