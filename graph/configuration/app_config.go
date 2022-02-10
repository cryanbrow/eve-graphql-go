package configuration

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

const configPath = "config.yml"

var TestConfigPath string

type Config struct {
	Application struct {
		Name        string `yaml:"name"`
		Environment string `yaml:"environment"`
	} `yaml:"application"`
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	Redis struct {
		Url      string `yaml:"url"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"redis"`
	Esi struct {
		Default struct {
			QueryParams []Key_value `yaml:"queryParams"`
			Url         string      `yaml:"url"`
		} `yaml:"default"`
	} `yaml:"esi"`
	Jaeger struct {
		Hostname string `yaml:"hostname"`
		Port     string `yaml:"port"`
		Protocol string `yaml:"protocol"`
		Route    string `yaml:"route"`
	} `yaml:"jaeger"`
}

type Key_value struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}

var AppConfig Config

func ReadFile() {
	var file *os.File
	var err error
	if TestConfigPath == "" {
		file, err = os.Open(configPath)
	} else {
		log.Debug("test config path is populated")
		file, err = os.Open(TestConfigPath)
	}
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

func LoadConfiguration() {
	SetupLogging()
	ReadFile()
	ReadEnv()
}
