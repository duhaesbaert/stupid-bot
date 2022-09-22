package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

// stopListening switches teh flag botListening, from Bot struct to false.
func (b Bot) stopListening(m *discordgo.MessageCreate) {
	if b.checkUserAllowed(m.Author.ID) {
		b.log.InfoLog("disabling listening on bot")
		b.config.BotListening = false
	} else {
		b.log.InfoLog("Only authorized users are allowed to disable or enable listening on bot")
	}
}

// startListening switches teh flag botListening, from Bot struct to true.
func (b Bot) startListening(m *discordgo.MessageCreate) {
	if b.checkUserAllowed(m.Author.ID) {
		b.log.InfoLog("enabling listening on bot")
		b.config.BotListening = true
	} else {
		b.log.InfoLog("Only authorized users are allowed to disable or enable listening on bot")
	}
}

func (b Bot) callForGaming(s *discordgo.Session, m *discordgo.MessageCreate) {
	b.log.InfoLog("calling everyone on server to play")
	jogo := b.gameToPlay(m.Content)

	message := fmt.Sprintf("@here BORA JOGAR %s", jogo)

	_, err := s.ChannelMessageSendReply(m.ChannelID, message, &discordgo.MessageReference{ChannelID: m.ChannelID, MessageID: m.Message.ID})
	if err != nil {
		b.log.ErrorLog(fmt.Sprintf("error while calling everyone to play: %s", err.Error()))
	}
}

func (b Bot) gameToPlay(message string) string {
	if strings.Contains(strings.ToLower(message), "cs") || strings.Contains(strings.ToLower(message), "csgo") {
		return "UM CSZINHO"
	}

	if strings.Contains(strings.ToLower(message), "pubg") {
		return "UM PUBG LIXO"
	}

	if strings.Contains(strings.ToLower(message), "dota") {
		return "UM DOTA DA DESGRAÇA"
	}

	if strings.Contains(strings.ToLower(message), "valorant") || strings.Contains(strings.ToLower(message), "valval") || strings.Contains(strings.ToLower(message), "val") {
		return "UM VALORANT, que é pior que CS."
	}

	return ""
}
func (b Bot) checkUserAllowed(userid string) bool {
	return userid == "343136401711169539"
}