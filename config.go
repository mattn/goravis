package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v2"
)

type LastCheck struct {
	ETag    string `yaml:"etag"`
	Version string `yaml:"version"`
	At      string `yaml:"at"`
}

type EndPoint struct {
	AccessToken string `yaml:"access_token"`
}

type Repo struct {
	EndPoint string `yaml:"endpoint"`
}

type Config struct {
	path              string
	LastCheck         *LastCheck           `yaml:"last_check"`
	CheckedCompletion bool                 `yaml:"checked_completion"`
	CompletionVersion string               `yaml:"completion_version"`
	EndPoints         map[string]*EndPoint `yaml:"endpoints"`
	Repos             map[string]*Repo     `yaml:"repos"`
}

func newConfig() *Config {
	home := os.Getenv("HOME")
	if home == "" && runtime.GOOS == "windows" {
		home = os.Getenv("USERPROFILE")
	}
	c := &Config{
		EndPoints: make(map[string]*EndPoint),
		Repos:     make(map[string]*Repo),
	}
	c.path = filepath.Join(home, ".travis", "config.yml")
	os.MkdirAll(filepath.Dir(c.path), 0700)
	return c
}

func (c *Config) Load() error {
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
