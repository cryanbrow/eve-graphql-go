package configuration

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

const configPath = "config.yml"

type Config struct {
	Redis struct {
		Url      string `yaml:"url"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"redis"`
}

var AppConfig Config

func ReadFile() {
	file, err := os.Open(configPath)
	if err != nil {
		processError(err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&AppConfig)
	if err != nil {
		processError(err)
	}
}

func ReadEnv() {
	err := envconfig.Process("", &AppConfig)
	if err != nil {
		processError(err)
	}
}

func processError(err error) {
	log.Errorf("Could not load config. : %v", err)
	os.Exit(2)
}

func init() {
	ReadFile()
	ReadEnv()
}
