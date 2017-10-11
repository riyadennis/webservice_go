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


func GetConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	if err != nil {
		panic(err.Error())
	}
	config, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	conf := new(Config)
	err = yaml.Unmarshal(config, &conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
