package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"stupid-bot/common"
	"stupid-bot/config"
)

type Bot struct {
	BotId string
	goBot *discordgo.Session
	botListening bool
	log common.Logger
	config *config.ConfigStruct
}

// NewBot instantiates and returns a new Bot struct.
func NewBot(log common.Logger, config *config.ConfigStruct, listening bool) Bot {
	return Bot{
		botListening: listening,
		log: log,
		config: config,
	}
}

// Start initializes the bot functionality, using the configuration already loaded from the config.json.
func (b Bot) Start() {
	b.log = common.NewLogger()
	b.log.InfoLog("starting bot")

	// creates a new session for the bot using the respective Token.
	bot, err := discordgo.New("Bot " + b.config.Token)
	if err != nil {
		b.log.ErrorLog(fmt.Sprintf("error creating bot session on Discord: %s", err.Error()))
		return
	}
	b.goBot = bot


	b.log.InfoLog("assigning user to bot")
	u, err := b.goBot.User("@me")
	if err != nil {
		b.log.ErrorLog(err.Error())
		return
	}
	b.BotId = u.ID

	b.goBot.AddHandler(b.messageHandler)
	b.log.InfoLog("connecting bot to discord")
	err = b.goBot.Open()
	if err != nil {
		b.log.ErrorLog(fmt.Sprintf("error opening connection: %s", err.Error()))
		return
	}

	b.log.InfoLog("bot running")
}

// stopListening switches teh flag botListening, from Bot struct to false.
func (b Bot) stopListening() {
	b.botListening = false
}

// startListening switches teh flag botListening, from Bot struct to true.
func (b Bot) startListening() {
	b.botListening = true
}
