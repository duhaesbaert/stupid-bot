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
	Token     string `json : "Token"`
	BotPrefix string `json : "BotPrefix"`
}

// ReadConfig reads the config.json file on the directory to use the bot information for connection.
func ReadConfig() error {
	common.NormalizedLog("reading config.json file to load configurations", common.Info)
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		common.NormalizedLog(err.Error(), common.Error)
		return err
	}
	common.NormalizedLog("config.json loaded successfully", common.Info)

	err = json.Unmarshal(file, &config)
	if err != nil {
		common.NormalizedLog(err.Error(), common.Error)
		return err
	}

	Token = config.Token
	BotPrefix = config.BotPrefix
	common.NormalizedLog("bot configuration loaded from config files", common.Info)
	return nil
}