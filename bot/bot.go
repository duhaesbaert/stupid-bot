package bot

import (
	"github.com/bwmarrin/discordgo"
	"stupid-bot/common"
	"stupid-bot/config"
)

var (
	BotId string
	goBot *discordgo.Session
)

// Start initializes the bot functionality, using the configuration already loaded from the config.json.
func Start() {
	// creates a new session for the bot using the respective Token.
	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		common.NormalizedLog(err.Error(), common.Error)
		return
	}

	u, err := goBot.User("@me")
	if err != nil {
		common.NormalizedLog(err.Error(), common.Error)
		return
	}

	BotId = u.ID

	// Adding handler function to handle our messages using AddHandler from discordgo package. We will declare messageHandler function later.
	goBot.AddHandler(messageHandler)


	err = goBot.Open()
	if err != nil {
		common.NormalizedLog(err.Error(), common.Error)
		return
	}

	common.NormalizedLog("bot running", common.Info)
}

// messageHandler watches for messages sent on the discord channel by other users and interacts with them, either by sending new messages or by performing actions.
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotId {
		return
	}
	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
	}
}