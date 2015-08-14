package config

import (
	"encoding/json"
	"io/ioutil"
	"os"

	log "github.com/Sirupsen/logrus"
)

type Config struct {
	Host       string
	Production bool
	Db         struct {
		Driver  string
		User    string
		Name    string
		SslMode string
	}
	Redis struct {
		Network        string
		Host           string
		MaxConnections int
	}
}

type GlobalConfig struct {
	Development Config
	Production  Config
}

var GlobalConf GlobalConfig
var Conf Config

func Load() {
	// Get the config file
	config_file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal("Cant read config.json")
	}

	err = json.Unmarshal(config_file, &GlobalConf)
	if err != nil {
		log.Fatal("Cant Unmarshal config file")
	}

	env := os.Getenv("ENV")

	if env == "production" {
		Conf = GlobalConf.Production
		Conf.Production = true
	} else {
		Conf = GlobalConf.Development
		Conf.Production = false
	}

	log.Info("Config file read")
}
