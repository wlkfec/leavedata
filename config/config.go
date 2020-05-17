package config

import (
	"github.com/BurntSushi/toml"
)

var (
	Config Configs
)

func init() {
	var cpath string = "config.toml"
	if _, err := toml.DecodeFile(cpath, &Config); err != nil {
	}
}

type Configs struct {
	MySQLDataSource string
	MQuri           string
	MQexchangeName  string
	MQexchangeType  string
	MQroutingKey    string
	MQqueueName     string
}
