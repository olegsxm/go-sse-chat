package config

import (
	"flag"
	"log/slog"
	"os"

	"gopkg.in/yaml.v2"

	_ "gopkg.in/yaml.v2"
)

type AppConfig struct {
	Production bool   `yaml:"production"`
	JWTSecret  string `yaml:"jwt_secret"`
	Server     struct {
		Address    string `yaml:"addr"`
		DevAddress string `yaml:"dev_addr"`
	} `yaml:"server"`

	Swagger struct {
		Url string `yaml:"url"`
	} `yaml:"swagger"`
}

var Config *AppConfig

func New() (*AppConfig, error) {
	if Config != nil {
		return Config, nil
	}

	configPath := flag.String("config", "config.yaml", "config file path")
	flag.Parse()

	file, err := os.Open(*configPath)

	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			slog.Error("Error closing config file", err.Error())
		}
	}(file)

	decoder := yaml.NewDecoder(file)

	if err := decoder.Decode(&Config); err != nil {
		return nil, err
	}

	return Config, nil
}
