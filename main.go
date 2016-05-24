package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/Ableton/go-travis"
	"github.com/alecthomas/kingpin"
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
