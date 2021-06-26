package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

const envFile = "env.yaml"

type Config struct {
	AppPort int      `yaml:"app_port"`
	DB      DBConfig `yaml:"db"`
}

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func (c *Config) load() error {
	yamlFile, err := ioutil.ReadFile(envFile)

	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return err
	}

	return nil
}

func New() (*Config, error) {
	var c Config
	err := c.load()

	return &c, err
}
