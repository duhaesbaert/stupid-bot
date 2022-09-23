package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"stupid-bot/common"
)

type ConfigStruct struct {
	Token        string `json:"Token"`
	BotPrefix    string `json:"BotPrefix"`
	log          common.Logger
	BotListening bool
}

// NewConfig reads the config.json file contained on the directory, and instantiates a new ConfigStruct to be used by the bot.
func NewConfig(log common.Logger) (*ConfigStruct, error) {
	var config *ConfigStruct

	log.InfoLog("reading config.json file to load configurations")
	config, err := readConfig(config)
	if err != nil {
		log.ErrorLog(fmt.Sprintf("error loading config from json file: %s", err.Error()))
		return config, err
	}
	log.InfoLog("configuration loaded from json file successfully")

	config.log = log
	log.InfoLog("all configurations loaded successfully")
	return config, nil
}

func readConfig(config *ConfigStruct) (*ConfigStruct, error) {
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		return &ConfigStruct{}, err
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		return &ConfigStruct{}, err
	}
	return config, nil
}

func (cs ConfigStruct) PrintConfiguration() {
	cs.log.DebugLog("Configurations:")
	cs.log.DebugLog("Token: " + cs.Token)
	cs.log.DebugLog("Prefix: " + cs.BotPrefix)
	cs.log.DebugLog(fmt.Sprintf("Listening: %b", cs.BotListening))
}
