package configuration

import (
	"os"
	"path/filepath"

	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

const configPath = "config.yml"

// TestConfigPath is file path used for unit tests.
var TestConfigPath string

// AppConfig is the package level variable exposing the applications configs.
var AppConfig Config

// Config is a struct in yaml format for storing configs for the application.
type Config struct {
	Application struct {
		Name        string `default:"eve-graphql-go" yaml:"name"`
		Environment string `default:"production" yaml:"environment"`
	} `yaml:"application"`
	Server struct {
		Port string `default:"8080" yaml:"port"`
	} `yaml:"server"`
	Caching struct {
		Impl string `default:"memory" yaml:"impl"`
	} `yaml:"cache"`
	Redis struct {
		URL      string `default:"localhost" yaml:"url"`
		Port     string `default:"30893" yaml:"port"`
		User     string `default:"" yaml:"user"`
		Password string `default:"" yaml:"password"`
	} `yaml:"redis"`
	Esi struct {
		Default struct {
			QueryParams []KeyValue `yaml:"queryParams"`
			URL         string     `default:"https://esi.evetech.net/latest" yaml:"url"`
		} `yaml:"default"`
	} `yaml:"esi"`
	Jaeger struct {
		Enabled  bool   `default:"false" yaml:"enabled"`
		Hostname string `default:"localhost" yaml:"hostname"`
		Port     string `default:"14268" yaml:"port"`
		Protocol string `default:"http" yaml:"protocol"`
		Route    string `default:"api/traces" yaml:"route"`
		Sample   struct {
			Percent int `default:"0" yaml:"percent"`
		} `yaml:"sample"`
	} `yaml:"jaeger"`
}

// KeyValue is a struct for representing simple Key Value pairs of configs.
type KeyValue struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}

func readFile() {
	var file *os.File
	var err error
	if TestConfigPath == "" {
		file, err = os.Open(filepath.Clean(configPath))
	} else {
		log.Debug("test config path is populated")
		file, err = os.Open(filepath.Clean(TestConfigPath))
	}
	if err != nil {
		if err.Error() == "open config.yml: The system cannot find the file specified." {
			log.Warn("Did not find config file. Proceeding with default config.")
			readEnv()
			return
		}
		processError(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Errorf("Error closing file: %s\n", err)
		}
	}()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&AppConfig)
	if err != nil {
		processError(err)
	}
}

func readEnv() {
	err := envconfig.Process("", &AppConfig)
	if err != nil {
		processError(err)
	}
}

func processError(err error) {
	log.Errorf("Could not load config. : %v", err)
	os.Exit(2)
}

func setupLogging() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)
}

// LoadConfiguration is for loading config from files and environment and setting it ready to be read.
func LoadConfiguration() {
	setupLogging()
	readFile()
	//ReadEnv()
}
