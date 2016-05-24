package main

import (
	"fmt"
	"github.com/Ableton/go-travis"
	"github.com/alecthomas/kingpin"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Config struct {
	LastCheck struct {
		ETag    string `yaml:"etag"`
		Version string `yaml:"version"`
		At      string `yaml:"at"`
	} `yaml:"last_check"`
	CheckedCompletion bool   `yaml:"checked_completion"`
	CompletionVersion string `yaml:"completion_version"`
	EndPoints         map[string]struct {
		AccessToken string `yaml:"access_token"`
	} `yaml:"endpoints"`
}

var (
	config Config
	client = travis.NewDefaultClient("")

	whoamiCommand = kingpin.Command("whoami", "Displays accounts and their subscription status.").Action(func(ctx *kingpin.ParseContext) error {
		err := client.Authentication.UsingTravisToken(config.EndPoints["https://api.travis-ci.org/"].AccessToken)
		if err != nil {
			return err
		}
		u, resp, err := client.Users.GetAuthenticated()
		if err != nil {
			return err
		}
		resp.Body.Close()
		fmt.Printf("You are %s (%s)\n", u.Login, u.Name)
		return nil
	})
)

func main() {
	home := os.Getenv("HOME")
	if home == "" && runtime.GOOS == "windows" {
		home = os.Getenv("USERPROFILE")
	}
	fn := filepath.Join(home, ".travis", "config.yml")
	b, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(b, &config)
	if err != nil {
		log.Fatal(err)
	}
	kingpin.Parse()
}
