package config

import (
	"os"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Kafka   Kafka `yaml:kafka`
	Reddit  Reddit `yaml:reddit`
	Article Article `yaml: article`
}
type Kafka struct {
	Broker []string `yaml: broker`
	Topic  string    `yaml: topic`
	File   string    `yaml: file`
}
type Reddit struct {
	Url string `yaml: url`
}
type Article struct {
	Key        string    `yaml: key`
	Url        string    `yaml: url`
	Source     string    `yaml: source`
	SortOption string    `yaml: sort_option `
}

const defaultConfigPath = "/src/github.com/webservice_go/config.yml"

func GetConfig() (conf *Config) {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + defaultConfigPath)
	if err != nil {
		panic(err.Error())
	}

	if err != nil {
		panic(err.Error())
	}
	config, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	conf = new(Config)
	err = yaml.Unmarshal(config, &conf)
	if err != nil {
		panic(err)
	}
	return conf
}
