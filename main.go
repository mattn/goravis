package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/Ableton/go-travis"
	"github.com/alecthomas/kingpin"
	"gopkg.in/yaml.v2"
)

const (
	VERSION = "0.0.1"
)

type Config struct {
	path      string
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

func (c *Config) Load() error {
	home := os.Getenv("HOME")
	if home == "" && runtime.GOOS == "windows" {
		home = os.Getenv("USERPROFILE")
	}
	c.path = filepath.Join(home, ".travis", "config.yml")
	b, err := ioutil.ReadFile(c.path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(b, &config)
}

func (c *Config) Save() error {
	b, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(c.path, b, 0600)
}

var (
	config Config
	client = travis.NewDefaultClient("")
)

func main() {
	err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	kingpin.Parse()
}
