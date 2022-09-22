package main

import(
	"stupid-bot/bot"
	"stupid-bot/common"
	"stupid-bot/config"
)

// main initializes the bot, reading the necessary configuration respectively starts the bot.
func main() {
	log := common.NewLogger()

	log.InfoLog("initializing bot...")
	config, err := config.NewConfig(log)
	if err != nil {
		log.ErrorLog(err.Error())
		return
	}

	botListening := true
	myBot := bot.NewBot(log, config, botListening)
	myBot.Start()

	<-make(chan struct{})
	return
}
