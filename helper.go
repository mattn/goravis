package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/alecthomas/kingpin"
)

func fatal(msg string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", msg, err)
	} else {
		fmt.Fprintf(os.Stderr, "%s\n", msg)
	}
	os.Exit(1)
}

func slug(ctx *kingpin.ParseContext) string {
	r := ctx.SelectedCommand.GetFlag("repo")
	if r != nil {
		rs := r.String()
		if rs != nil && *rs != "" {
			return *rs
		}
	}
	b, err := exec.Command("git", "config", "--get", "travis.slug").CombinedOutput()
	if len(b) > 0 {
		return strings.TrimSpace(string(b))
	}

	b, err = exec.Command("git", "name-rev", "--name-only", "HEAD").CombinedOutput()
	if err != nil {
		fatal(`Can't figure out GitHub repo name. Ensure you're in the repo directory, or specify the repo name via the -r option (e.g. travis <command> -r <owner>/<repo>)`, nil)
	}
	remote := "origin"
	s := strings.TrimSpace(string(b))
	b, err = exec.Command("git", "config", "--get", "branch."+s+".remote").CombinedOutput()
	if err == nil {
		remote = string(b)
	}
	b, err = exec.Command("git", "ls-remote", "--get-url", remote).CombinedOutput()
	if err != nil {
		fatal(`Can't figure out GitHub repo name. Ensure you're in the repo directory, or specify the repo name via the -r option (e.g. travis <command> -r <owner>/<repo>)`, nil)
	}
	s = strings.TrimSpace(string(b))
	m := regexp.MustCompile(`[:/]([^/]+/[^/]+?)(\.git)?$`).FindStringSubmatch(s)
	if len(m) != 3 {
		fatal(`Can't figure out GitHub repo name. Ensure you're in the repo directory, or specify the repo name via the -r option (e.g. travis <command> -r <owner>/<repo>)`, nil)
	}
	err = exec.Command("git", "config", "travis.slug", m[1]).Run()
	if err != nil {
		fatal(`Can't figure out GitHub repo name. Ensure you're in the repo directory, or specify the repo name via the -r option (e.g. travis <command> -r <owner>/<repo>)`, nil)
	}
	return m[1]
}

func token() string {
	err := config.Load()
	if err != nil {
		fatal("not logged in, please run travis login", nil)
	}
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

func githubAuth(githubToken string) error {
	travisToken, _, err := client.Authentication.UsingGithubToken(githubToken)
	if err != nil {
		return err
	}
	config.EndPoints["https://api.travis-ci.org/"] = &EndPoint{
		AccessToken: string(travisToken),
	}
	return config.Save()
}
