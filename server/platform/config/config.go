package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
}

var config *Config

func init() {
	configFilePath := ""
	f, err := os.ReadFile(configFilePath)
	if err != nil {
		log.Printf("Could not read config file: %v\n", err)
		os.Exit(-1)
	}
	config = new(Config)
	yaml.Unmarshal(f, config)
}
func GetConfig() *Config {
	return config
}
