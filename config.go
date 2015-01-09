package x

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// StoreConfig contains configuration values for the storage engine
type StoreConfig map[string]string

// Handler contains single handler configuration
type Handler struct {
	Sizes   []Size
	Formats []Format
	Store   []StoreConfig
}

// Config contains the configuration options for the service
type Config map[string]Handler

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
	return yaml.Unmarshal(data, c)
}

// UnmarshalYAML unmarshals the Size confing and validates it
func (s *Size) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var tmap map[string]int
	if err := unmarshal(&tmap); err != nil {
		return err
	}
	for k, v := range tmap {
		switch SizeType(k) {
		case Square, Max, MaxHeight, MaxWidth:
			s.Type = SizeType(k)
			s.Value = v
		default:
			return errors.New(fmt.Sprintf("Invalid size in config: %s", k))
		}
	}
	return nil
}

// https://github.com/go-yaml/yaml/issues/67

// func (f *Format) UnmarshalYAML(unmarshal func(interface{}) error) error {
// 	var fm string
// 	if err := unmarshal(&fm); err != nil {
// 		return err
// 	}
// 	format := Format(fm)
// 	switch format {
// 	case JPG, PNG, WebP:
// 		print(format)
// 		f = &format
// 	default:
// 		return errors.New(fmt.Sprintf("Invalid file format in config: %s", fm))
// 	}
// 	return nil
// }
