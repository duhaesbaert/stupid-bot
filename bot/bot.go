package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
	"stupid-bot/common"
	"stupid-bot/config"
)

var (
	BotId string
	goBot *discordgo.Session
	log common.Logger
)

// Start initializes the bot functionality, using the configuration already loaded from the config.json.
func Start() {
	log = common.NewLogger()
	log.InfoLog("initializing bot")
	// creates a new session for the bot using the respective Token.
	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		log.ErrorLog(fmt.Sprintf("error creating bot session on Discord: %s", err.Error()))
		return
	}

	log.InfoLog("assigning user to bot")
	u, err := goBot.User("@me")
	if err != nil {
		log.ErrorLog(err.Error())
		return
	}
	BotId = u.ID

	goBot.AddHandler(messageHandler)

	log.InfoLog("connecting bot to discord")
	err = goBot.Open()
	if err != nil {
		log.ErrorLog(fmt.Sprintf("error opening connection: %s", err.Error()))
		return
	}

	log.InfoLog("bot running")
}

// messageHandler watches for messages sent on the discord channel by other users and interacts with them, either by sending new messages or by performing actions.
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotId {
		log.WarningLog("message sent by bot -> ignoring")
		return
	}

	err := messageSelector(s, m)
	if err != nil {
		log.ErrorLog(fmt.Sprintf("could not sent message: %s", err.Error()))
	}
}

func messageSelector(s *discordgo.Session, m *discordgo.MessageCreate) error {
	log.InfoLog(fmt.Sprintf("found message on thread: %s", m.Content))

	if strings.Contains(m.Content, "cs") {
		return csgoMessage(s, m)
	}

	if strings.Contains(m.Content, "gamersclub") {
		return gamersclubMessage(s, m)
	}

	if strings.Contains(m.Content, "vitor") {
		return vitorMessage(s, m)
	}

	if strings.Contains(m.Content, "wolke") {
		return vitorMessage(s, m)
	}

	if strings.Contains(m.Content, "f") {
		return fMessage(s, m)
	}

	if strings.Contains(m.Content, "bob") {
		return bobMessage(s, m)
	}

	if strings.Contains(m.Content, "chinela"){
		return chinelaMessage(s, m)
	}

	if strings.Contains(m.Content, "bora") ||
		strings.Contains(m.Content, "online") ||
		strings.Contains(m.Content, "vamo") ||
		strings.Contains(m.Content, "jogar") ||
		strings.Contains(m.Content, "voltei") ||
		strings.Contains(m.Content, "avisa") ||
		strings.Contains(m.Content, "hoje") ||
		strings.Contains(m.Content, "jogando") ||
		strings.Contains(m.Content, "?") {
		return boraMessage(s, m)
	}

	if strings.Contains(m.Content, "monstro"){
		return monstroMessage(s, m)
	}

	if strings.Contains(m.Content, "hess") || strings.Contains(m.Content, "hsz") || strings.Contains(m.Content, "geferson") {
		return hszMessage(s, m)
	}

	log.WarningLog("no message has been identified to be sent.")
	return nil
}