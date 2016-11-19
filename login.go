package main

import (
	"bufio"
	"fmt"
	"os"
	"syscall"
	"time"

	"github.com/alecthomas/kingpin"
	"github.com/octokit/go-octokit/octokit"
	"golang.org/x/crypto/ssh/terminal"
)

var loginCommand = kingpin.Command("login", "authenticates against the API and stores the token").Action(func(ctx *kingpin.ParseContext) error {
	if token() != "" {
		return nil
	}

	if err := tryHubConfig(); err == nil {
		return nil
	}

	r := bufio.NewReader(os.Stdin)
	fmt.Print("Username: ")
	b, _, err := r.ReadLine()
	if err != nil {
		return err
	}
	username := string(b)

	fmt.Print("Password: ")
	b, err = terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return err
	}
	password := string(b)
	fmt.Println()

	host, err := os.Hostname()
	if err != nil {
		return err
	}
	param := octokit.Authorization{
		Scopes: []string{"user", "user:email", "repo"},
		Note:   "goravis on " + host + " " + time.Now().Format("2006/01/02 15:04:05"),
	}
	uri, err := octokit.AuthorizationsURL.Expand(nil)
	if err != nil {
		return err
	}
	authMethod := octokit.BasicAuth{
		Login:    username,
		Password: password,
	}
	client := octokit.NewClient(authMethod)
	authToken, result := client.Authorizations(uri).Create(param)
	if result.Err != nil {
		rerr, ok := result.Err.(*octokit.ResponseError)
		if !ok || rerr.Type != octokit.ErrorOneTimePasswordRequired {
			return result.Err
		}
		fmt.Print("OTP: ")
		b, _, err = r.ReadLine()
		if err != nil {
			return err
		}
		authMethod.OneTimePassword = string(b)
		client = octokit.NewClient(authMethod)
		authToken, result = client.Authorizations(uri).Create(param)
		if result.Err != nil {
			return result.Err
		}
	}

	return githubAuth(authToken.Token)
})
