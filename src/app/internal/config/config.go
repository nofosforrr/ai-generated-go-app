package config

import (
	"os"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type Service struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}


type Config struct {
	Service *Service `yaml:"service"`
}

func LoadConfig() (*Config, error) {
	cfgPath := os.Getenv("CONFIG_PATH")
	if cfgPath == "" {
		logrus.Fatal("Missing CONFIG_PATH variable.")
	}

	file, err := os.ReadFile(cfgPath)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
