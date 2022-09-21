package main

import(
	"stupid-bot/bot"
	"stupid-bot/common"
	"stupid-bot/config"
)

// main initializes the bot, reading the necessary configuration respectively starts the bot.
func main() {
	err := config.ReadConfig()
	if err != nil {
		common.ErrorLog(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}
