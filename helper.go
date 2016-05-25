package main

import (
	"os"
	"os/exec"
	"regexp"
	"strings"
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
	m := regexp.MustCompile(`[:/]([^/]+/[^/]+)(\.git)?$`).FindStringSubmatch(s)
	if len(m) != 3 {
		return ""
	}
	return m[1]
}

func token() string {
	t := os.Getenv("TRAVIS_TOKEN")
	ep, ok := config.EndPoints["https://api.travis-ci.org/"]
	if ok {
		t = ep.AccessToken
	}
	return t
}

func auth() error {
	return client.Authentication.UsingTravisToken(token())
}
