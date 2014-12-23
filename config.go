package main

import (
	"io"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// Config contains the configuration options for the service
type Config struct {
}

// NewConfig creates new configuration
func NewConfig() *Config {
	return &Config{}
}

// Load reads the configuration file from the provided file name
func (c *Config) Load(fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	return c.LoadStream(f)
}

// LoadStream reads the configuration from the provided reader
func (c *Config) LoadStream(r io.Reader) error {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(data, &c); err != nil {
		return err
	}
	return nil
}
