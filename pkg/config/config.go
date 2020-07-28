package config

import (
	"io/ioutil"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// LoadConfig tries to load the config from the given path
func LoadConfig(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "could not load config from path")
	}

	var result Config
	err = yaml.Unmarshal(data, &result)
	if err != nil {
		return nil, errors.Wrap(err, "could not unmarshal config")
	}

	return &result, nil
}
