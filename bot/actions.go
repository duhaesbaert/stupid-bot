package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
	"stupid-bot/common"
	"time"
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

func (b Bot) startPoll(s *discordgo.Session, m *discordgo.MessageCreate) {
	pollMessage := strings.Replace(strings.ToUpper(m.Content), "/POLL", "", -1)

	pollMessage = strings.Trim(pollMessage, " ")

	if len(pollMessage) <= 0 {
		message := "Para iniciar uma poll, digite o topico logo após /poll"
		_, err := s.ChannelMessageSendReply(m.ChannelID, message, &discordgo.MessageReference{ChannelID: m.ChannelID, MessageID: m.Message.ID})
		if err != nil {
			b.log.ErrorLog(fmt.Sprintf("error initiating poll: %s", err.Error()))
		}
		return
	}

	timer := common.Newtimer(1)
	showTime := timer.ShowNormalizedTime()

	msgEmbed := generatePollEmbed(pollMessage, m.Author.Username, m.Author.AvatarURL(""), showTime)

	// delete original message from user
	err := s.ChannelMessageDelete(m.ChannelID, m.Message.ID)
	if err != nil {
		b.log.ErrorLog(fmt.Sprintf("error while trying to delete message: %s", err.Error()))
	}

	// send new message with poll
	newPoll, err := s.ChannelMessageSendEmbed(m.ChannelID, msgEmbed)
	if err != nil {
		b.log.ErrorLog(fmt.Sprintf("error initiating poll: %s", err.Error()))
	}

	// add reaction to indicate how to vote
	err = s.MessageReactionAdd(newPoll.ChannelID, newPoll.ID, "👍")
	if err != nil {
		b.log.ErrorLog(fmt.Sprintf("error adding reactions to poll: %s", err.Error()))
	}
	err = s.MessageReactionAdd(newPoll.ChannelID, newPoll.ID, "👎")
	if err != nil {
		b.log.ErrorLog(fmt.Sprintf("error adding reactions to poll: %s", err.Error()))
	}

	go b.deletePoll(pollMessage, m.Author, s, newPoll, timer)
}

func (b Bot) deletePoll(pollMessage string, originalAuthor *discordgo.User, s *discordgo.Session, m *discordgo.Message, timer common.Timer) {
	ticker := time.NewTicker(time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				err := s.ChannelMessageDelete(m.ChannelID, m.ID)
				if err != nil {
					b.log.ErrorLog(fmt.Sprintf("error deleting poll message: %s", err.Error()))
				}
			case <-ticker.C:
				timer = timer.Countdown()
				_, err := s.ChannelMessageEditEmbed(m.ChannelID, m.ID, generatePollEmbed(pollMessage, originalAuthor.Username, originalAuthor.AvatarURL(""), timer.ShowNormalizedTime()))
				if err != nil {
					b.log.ErrorLog(fmt.Sprintf("error editing poll message with timer: %s", err.Error()))
				}
			}
		}
	}()

	d, err := time.ParseDuration(timer.ShowNormalizedTime())
	if err != nil {
		b.log.ErrorLog(fmt.Sprintf("error parsing duration of timer: %s", err.Error()))
	}

	time.Sleep(d)
	ticker.Stop()
	done <- true
	b.log.InfoLog("poll ticker stopped")
}

func generatePollEmbed(pollMessage, author, avatar, time string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Type:        discordgo.EmbedTypeRich,
		Title:       fmt.Sprintf("%s", pollMessage),
		Description: fmt.Sprintf("%s até fechar.", time),
		Color:       4,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    author,
			IconURL: avatar,
		},
	}
}
