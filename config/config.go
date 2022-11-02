package config

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type Configuration struct {
	OldRedis          Redis `yaml:"old_redis"`
	ConcurrentWorkers int   `yaml:"concurrent_workers"`
	NewRedis          Redis `yaml:"new_redis"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	Database []int  `yaml:"database"`
}

// ParseConfig is to parse YAML configuration file
func ParseConfig(configFile string) Configuration {
	var configContent Configuration
	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		logrus.Errorf("Error while reading file %v", err)
	}
	err = yaml.Unmarshal(yamlFile, &configContent)
	if err != nil {
		logrus.Errorf("Error in parsing file to yaml content %v", err)
	}
	return configContent
}
