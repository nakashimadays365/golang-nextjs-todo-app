package config

import "os"

type Config struct {
	DBConfig DBConfig `yaml:"db"`
}

type DBConfig struct {
	Name string `yaml:"name"`
	Host string `yaml:"host"`
	Pass string `yaml:"pass"`
	User string `yaml:"user"`
}

func NewConfig() (*Config, error) {

	var conf Config
	conf.DBConfig.Name = os.Getenv("MYSQL_NAME")
	conf.DBConfig.Host = os.Getenv("MYSQL_HOST")
	conf.DBConfig.Pass = os.Getenv("MYSQL_PASS")
	conf.DBConfig.User = os.Getenv("MYSQL_USER")

	return &conf, nil
}
