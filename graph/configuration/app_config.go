package configuration

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

const configPath = "config.yml"

var TestConfigPath string
var AppConfig Config

type Config struct {
	Application struct {
		Name        string `default:"eve-graphql-go" yaml:"name"`
		Environment string `default:"test" yaml:"environment"`
	} `yaml:"application"`
	Server struct {
		Port string `default:"8080" yaml:"port"`
	} `yaml:"server"`
	Caching struct {
		Impl string `default:"memory" yaml:"impl"`
	} `yaml:"cache"`
	Redis struct {
		Url      string `default:"localhost" yaml:"url"`
		Port     string `default:"30893" yaml:"port"`
		User     string `default:"" yaml:"user"`
		Password string `default:"" yaml:"password"`
	} `yaml:"redis"`
	Esi struct {
		Default struct {
			QueryParams []Key_value `yaml:"queryParams"`
			Url         string      `default:"https://esi.evetech.net/latest" yaml:"url"`
		} `yaml:"default"`
	} `yaml:"esi"`
	Jaeger struct {
		Hostname string `default:"localhost" yaml:"hostname"`
		Port     string `default:"14268" yaml:"port"`
		Protocol string `default:"http" yaml:"protocol"`
		Route    string `default:"api/traces" yaml:"route"`
		Sample   struct {
			Percent int `default:"0" yaml:"percent"`
		} `yaml:"sample"`
	} `yaml:"jaeger"`
}

type Key_value struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}

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
	//ReadEnv()
}
