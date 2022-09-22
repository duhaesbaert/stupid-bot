package config

import (
	"encoding/json"
	"io/ioutil"
	"stupid-bot/common"
)

var (
	config *ConfigStruct // value from config.json
)

type ConfigStruct struct {
	Token    	string `json:"Token"`
	BotPrefix	string `json:"BotPrefix"`
	log			common.Logger
}

// NewConfig reads the config.json file contained on the directory, and instantiates a new ConfigStruct to be used by the bot.
func NewConfig(log common.Logger) (*ConfigStruct, error) {
	log = log

	log.InfoLog("reading config.json file to load configurations")
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.ErrorLog(err.Error())
		return &ConfigStruct{}, err
	}
	log.InfoLog("config.json loaded successfully")

	err = json.Unmarshal(file, &config)
	if err != nil {
		return &ConfigStruct{}, err
	}

	log.InfoLog("bot configuration loaded from config files")
	return config, nil
}

// UpdateConfig reads the config.json file on the directory to use the bot information for connection.
func (cs ConfigStruct) UpdateConfig() (*ConfigStruct, error) {
	cs.log.InfoLog("reading config.json file to load configurations")
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		cs.log.ErrorLog(err.Error())
		return &ConfigStruct{}, err
	}
	cs.log.InfoLog("config.json loaded successfully")

	err = json.Unmarshal(file, &config)
	if err != nil {
		return &ConfigStruct{}, err
	}

	cs.log.InfoLog("bot configuration loaded from config files")
	return config, nil
}
