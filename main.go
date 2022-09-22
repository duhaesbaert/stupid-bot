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
	err := config.ReadConfig()
	if err != nil {
		log.ErrorLog(err.Error())
		return
	}

	myBot := bot.NewBot(true, log)
	myBot.Start()

	<-make(chan struct{})
	return
}
