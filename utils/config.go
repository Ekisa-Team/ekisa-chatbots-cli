package utils

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	ClientID         int32  `yaml:"client_id"`
	ApiEndpoint      string `yaml:"api_endpoint"`
	ConnectionString string `yaml:"connection_string"`
}

func (c *Config) LoadConfig(path string) (*Config, error) {
	buffer, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(buffer, c)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %v", path, err)
	}

	return c, nil
}
