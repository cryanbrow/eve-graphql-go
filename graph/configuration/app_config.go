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
	Esi struct {
		Default struct {
			Query_params []Key_value `yaml:"query_params"`
			Url          string      `yaml:"url"`
		} `yaml:"default"`
	} `yaml:"esi"`
}

type Key_value struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
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

func SetupLogging() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)
}

func init() {
	SetupLogging()
	ReadFile()
	ReadEnv()
}
