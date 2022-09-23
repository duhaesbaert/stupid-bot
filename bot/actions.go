package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
	"stupid-bot/common"
	"time"
)

type BotActions struct {
	bot Bot
	log common.Logger
	s   *discordgo.Session
	m   *discordgo.MessageCreate
}

// executeActions identifies which command has been written by the member, and requests for the responsible function.
func (b BotActions) executeActions() {
	if strings.ToLower(b.m.Content) == "/stop_listening" {
		b.stopListening()
	} else if strings.ToLower(b.m.Content) == "/start_listening" {
		b.startListening()
	} else if strings.HasPrefix(strings.ToLower(b.m.Content), "/play") {
		b.callForGaming()
	} else if strings.HasPrefix(strings.ToLower(b.m.Content), "/poll") {
		b.startPoll()
	}
}

// stopListening switches teh flag botListening, from Bot struct to false.
func (b BotActions) stopListening() {
	if b.checkUserAllowed(b.m.Author.ID) {
		b.log.InfoLog("disabling listening on bot")
		b.bot.config.BotListening = false
	} else {
		b.log.InfoLog("Only authorized users are allowed to disable or enable listening on bot")
	}
}

// startListening switches teh flag botListening, from Bot struct to true.
func (b BotActions) startListening() {
	if b.checkUserAllowed(b.m.Author.ID) {
		b.log.InfoLog("enabling listening on bot")
		b.bot.config.BotListening = true
	} else {
		b.log.InfoLog("Only authorized users are allowed to disable or enable listening on bot")
	}
}

// callForGaming sends a message into the channel mentioning @here to notify all users.
func (b BotActions) callForGaming() {
	b.log.InfoLog("calling everyone on server to play")
	jogo := b.gameToPlay(b.m.Content)

	message := fmt.Sprintf("@here BORA JOGAR %s", jogo)

	_, err := b.s.ChannelMessageSendReply(b.m.ChannelID, message, &discordgo.MessageReference{ChannelID: b.m.ChannelID, MessageID: b.m.Message.ID})
	if err != nil {
		b.log.ErrorLog(fmt.Sprintf("error while calling everyone to play: %s", err.Error()))
	}
}

// gameToPlay selects the message from a predefined list and returns.
func (b BotActions) gameToPlay(message string) string {
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

// checkUserAllowed identifies if the user executing the action is allowed to.
func (b BotActions) checkUserAllowed(userid string) bool {
	return userid == "343136401711169539"
}

// startPoll starts a poll into the channel, which will delete the message who requested the poll. The poll will be an EmbedMessage,
// that last for 5 minutes and automatically adds votes for up and down. Once the 5 minutes have passed, the poll is deleted, and
// the results are posted back into the channel.
func (b BotActions) startPoll() {
	pollMessage := strings.Replace(strings.ToUpper(b.m.Content), "/POLL", "", -1)

	pollMessage = strings.Trim(pollMessage, " ")

	if len(pollMessage) <= 0 {
		message := "Para iniciar uma poll, digite o topico logo após /poll"
		_, err := b.s.ChannelMessageSendReply(b.m.ChannelID, message, &discordgo.MessageReference{ChannelID: b.m.ChannelID, MessageID: b.m.Message.ID})
		if err != nil {
			b.log.ErrorLog(fmt.Sprintf("error initiating poll: %s", err.Error()))
		}
		return
	}

	timer := common.Newtimer(5)
	showTime := timer.ShowNormalizedTime()

	msgEmbed := generatePollEmbed(pollMessage, b.m.Author.Username, b.m.Author.AvatarURL(""), showTime)

	// delete original message from user
	err := b.s.ChannelMessageDelete(b.m.ChannelID, b.m.Message.ID)
	if err != nil {
		b.log.ErrorLog(fmt.Sprintf("error while trying to delete message: %s", err.Error()))
	}

	// send new message with poll
	newPoll, err := b.s.ChannelMessageSendEmbed(b.m.ChannelID, msgEmbed)
	if err != nil {
		b.log.ErrorLog(fmt.Sprintf("error initiating poll: %s", err.Error()))
	}

	// add reaction to indicate how to vote
	err = b.s.MessageReactionAdd(newPoll.ChannelID, newPoll.ID, "👍")
	if err != nil {
		b.log.ErrorLog(fmt.Sprintf("error adding reactions to poll: %s", err.Error()))
	}
	err = b.s.MessageReactionAdd(newPoll.ChannelID, newPoll.ID, "👎")
	if err != nil {
		b.log.ErrorLog(fmt.Sprintf("error adding reactions to poll: %s", err.Error()))
	}

	go b.runPollTicker(pollMessage, b.m.Author, newPoll, timer)
}

// runPollTicker is a ticker which controls the time and updates the message back into the channel, updating the timer for the poll to finish.
func (b BotActions) runPollTicker(pollMessage string, originalAuthor *discordgo.User, m *discordgo.Message, timer common.Timer) {
	ticker := time.NewTicker(time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				tup, err := b.s.MessageReactions(m.ChannelID, m.ID, "👍", 100, "", "")
				if err != nil {
					b.log.ErrorLog(fmt.Sprintf("error reading poll results: %s", err.Error()))
				}

				tdown, err := b.s.MessageReactions(m.ChannelID, m.ID, "👎", 100, "", "")
				if err != nil {
					b.log.ErrorLog(fmt.Sprintf("error reading poll results: %s", err.Error()))
				}

				err = b.s.ChannelMessageDelete(m.ChannelID, m.ID)
				if err != nil {
					b.log.ErrorLog(fmt.Sprintf("error deleting poll message: %s", err.Error()))
				}

				pollMessage = fmt.Sprintf("**%s** \n 👍 %d x %d 👎", pollMessage, len(tup)-1, len(tdown)-1)
				_, err = b.s.ChannelMessageSend(m.ChannelID, pollMessage)
				if err != nil {
					b.log.ErrorLog(fmt.Sprintf("error sending poll results back to channel: %s", err.Error()))
				}

			case <-ticker.C:
				timer = timer.Countdown()
				_, err := b.s.ChannelMessageEditEmbed(m.ChannelID, m.ID, generatePollEmbed(pollMessage, originalAuthor.Username, originalAuthor.AvatarURL(""), timer.ShowNormalizedTime()))
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

// generatePollEmbed generates a message object of poll.
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
