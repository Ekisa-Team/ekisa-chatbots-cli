package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

const (
	ENV_CLIENT_ID               = "CLIENT_ID"
	ENV_CONN_STRING             = "CONN_STRING"
	ENV_UPLOAD_APPOINTMENTS_URI = "UPLOAD_APPOINTMENTS_URI"
)

type Config struct {
	Application struct {
		ClientID string `yaml:"client_id"`
	} `yaml:"application"`
	Database struct {
		ConnectionString string `yaml:"connection_string"`
	} `yaml:"database"`
	Webhooks struct {
		UploadAppointmentsUri string `yaml:"upload_appointments_uri"`
	} `yaml:"webhooks"`
}

func (c *Config) ReadConfig(path string) (*Config, error) {
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
