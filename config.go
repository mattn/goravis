package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v2"
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

	EndPoints map[string]struct {
		AccessToken string `yaml:"access_token"`
	} `yaml:"endpoints"`
	Repos map[string]struct {
		EndPoint string `yaml:"endpoint"`
	} `yaml:"repos"`
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
