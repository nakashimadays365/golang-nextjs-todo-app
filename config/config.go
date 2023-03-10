package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"todo/logger"
)

type Config struct {
	DBConfig DBConfig `yaml:"db"`
}

type DBConfig struct {
	Name string `yaml:"name"`
	Host string `yaml:"host"`
	Pass string `yaml:"pass"`
	User string `yaml:"user"`
}

func NewConfig(env string) (*Config, error) {
	f, err := os.Open(fmt.Sprintf("%s.yaml", env))
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	defer f.Close()
	conf := Config{}
	if err := yaml.NewDecoder(f).Decode(&conf); err != nil {
		logger.Error(err)
		return nil, err
	}

	return &conf, nil
}
