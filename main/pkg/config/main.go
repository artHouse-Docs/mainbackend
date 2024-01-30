package config

import (
	"flag"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type ServerConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Salt string `yaml:"salt"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"user"`
	Password string `yaml:"password"`
}

type YamlConfig struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
}

var configPath = flag.String("cfg", "config.yaml", "path to config file")

func Configure() YamlConfig {
	var config YamlConfig

	flag.Parse()
	print(*configPath)

	configFile, err := os.ReadFile(*configPath)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}
