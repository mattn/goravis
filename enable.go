package main

import (
	"os/exec"
	"regexp"
	"strings"

	"github.com/alecthomas/kingpin"
)

func slug() string {
	b, _ := exec.Command("git", "config", "--get", "travis.slug").CombinedOutput()
	if len(b) > 0 {
		return strings.TrimSpace(string(b))
	}

	b, err := exec.Command("git", "name-rev", "--name-only", "HEAD").CombinedOutput()
	if err != nil {
		return ""
	}
	remote := "origin"
	s := strings.TrimSpace(string(b))
	b, err = exec.Command("git", "config", "--get", "branch."+s+".remote").CombinedOutput()
	if err == nil {
		remote = string(b)
	}
	b, err = exec.Command("git", "ls-remote", "--get-url", remote).CombinedOutput()
	if err != nil {
		return ""
	}
	s = strings.TrimSpace(string(b))
	m := regexp.MustCompile(`([^/]+/[^/]+)(\.git)?$`).FindStringSubmatch(s)
	if len(m) != 3 {
		return ""
	}
	return m[1]
}

var enableCommand = kingpin.Command("enable", "Enables a project.").Action(func(ctx *kingpin.ParseContext) error {
	err := client.Authentication.UsingTravisToken(config.EndPoints["https://api.travis-ci.org/"].AccessToken)
	if err != nil {
		return err
	}

	/*
		resp, err := client.Users.Sync()
		if resp != nil {
			println("resp")
			resp.Body.Close()
			if resp.StatusCode != 409 && err != nil {
				return err
			}
		} else if err != nil {
			return err
		}
	*/

	s := slug()
	repo, resp, err := client.Repositories.GetFromSlug(s)
	if err != nil {
		return err
	}
	resp.Body.Close()

	type Hook struct {
		Id     uint `json:"id"`
		Active bool `json:"active"`
	}

	req, err := client.NewRequest("PUT", "/hooks/", struct {
		Hook `json:"hook"`
	}{
		Hook{repo.Id, true},
	}, nil)
	if err != nil {
		return err
	}
	resp, err = client.Do(req, nil)
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil
})
