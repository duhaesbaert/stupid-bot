package config

import (
	"encoding/json"
	"io/ioutil"
	"stupid-bot/common"
)

var (
	Token     string // value of Token from config.json
	BotPrefix string // value of BotPrefix from config.json

	config *configStruct // value from config.json
)

type configStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
}

// ReadConfig reads the config.json file on the directory to use the bot information for connection.
func ReadConfig() error {
	log := common.NewLogger()

	log.InfoLog("reading config.json file to load configurations")
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.ErrorLog(err.Error())
		return err
	}
	log.InfoLog("config.json loaded successfully")

	err = json.Unmarshal(file, &config)
	if err != nil {
		log.ErrorLog(err.Error())
		return err
	}

	Token = config.Token
	BotPrefix = config.BotPrefix
	log.InfoLog("bot configuration loaded from config files")
	return nil
}
