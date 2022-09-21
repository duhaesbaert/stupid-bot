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
)

// Start initializes the bot functionality, using the configuration already loaded from the config.json.
func Start() {
	common.NormalizedLog("initializing bot", common.Info)
	// creates a new session for the bot using the respective Token.
	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		common.NormalizedLog(err.Error(), common.Error)
		return
	}

	common.NormalizedLog("assigning user to bot", common.Info)
	u, err := goBot.User("@me")
	if err != nil {
		common.NormalizedLog(err.Error(), common.Error)
		return
	}
	BotId = u.ID

	goBot.AddHandler(messageHandler)

	common.NormalizedLog("connecting bot to discord", common.Info)
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

	err := messageSelector(s,m)
	if err != nil {
		common.NormalizedLog(fmt.Sprintf("could not sent message: %s", err.Error()), common.Error)
	}
}

func messageSelector(s *discordgo.Session, m *discordgo.MessageCreate) error {
	if strings.Contains(m.Content, "cs") {
		return csgoMessage(s, m)
	}

	if strings.Contains(m.Content, "gamersclub") {
		gamersclubMessage(s, m)
	}

	if strings.Contains(m.Content, "vitor") {
		vitorMessage(s, m)
	}

	if strings.Contains(m.Content, "wolke") {
		vitorMessage(s, m)
	}
	return nil
}


func csgoMessage(s *discordgo.Session, m *discordgo.MessageCreate) error {
	_, err := s.ChannelMessageSend(m.ChannelID, "BORA JOGAR UM CSGO PORRA")
	return err
}

func gamersclubMessage(s *discordgo.Session, m *discordgo.MessageCreate) error {
	_, err := s.ChannelMessageSend(m.ChannelID, "Gamers Club é muito coisa de try hard")
	return err
}

func vitorMessage(s *discordgo.Session, m *discordgo.MessageCreate) error {
	_, err := s.ChannelMessageSend(m.ChannelID, "vitor = baiter")
	if err != nil {
		return err
	}
	_, err = s.ChannelMessageSend(m.ChannelID, "ele tb nunca planta a bomba")
	return err
}