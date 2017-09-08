package config

import (
	"os"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Db     Db
	Kafka  Kafka
	Reddit Reddit
}
type Db struct {
	Host     string
	Name     string
	User     string
	Password string
}

type Kafka struct {
	Host  string
	Topic string
}
type Reddit struct {
	Url string
}

const defaultConfigPath = "/config.yml"

func GetConfig() (con *Config) {
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
	con = new(Config)
	yaml.Unmarshal(config, &con)
	return con
}
