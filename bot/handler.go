package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"math/rand"
	"strings"
)

// messageHandler watches for messages sent on the discord channel by other users and interacts with them, either by sending new messages or by performing actions.
func (b Bot) messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == b.BotId {
		b.log.WarningLog("message sent by bot -> ignoring")
		return
	}

	if b.isAction(m) {
		b.executeActions(s, m)
		return
	}

	if b.config.BotListening {
		if m.Content == "" {
			content, attachments, err := b.channelMessageLookup(s, m)
			if err != nil {
				b.log.ErrorLog(err.Error())
				return
			}

			m.Content = content
			m.Attachments = attachments
		}

		err := b.messageSelector(s, m)
		if err != nil {
			b.log.ErrorLog(fmt.Sprintf("could not sent message: %s", err.Error()))
		}
	}
}

func (b Bot) isAction(m *discordgo.MessageCreate) bool {
	return strings.HasPrefix(m.Content, "/")
}

// channelMessageLookup uses the channel id and message id to find the received message and read it's content.
func (b Bot) channelMessageLookup(s *discordgo.Session, m *discordgo.MessageCreate) (string, []*discordgo.MessageAttachment, error) {
	chanMsgs, err := s.ChannelMessages(m.ChannelID, 1, "", "", m.ID)
	if err != nil {
		b.log.ErrorLog(fmt.Sprintf("unable to get messages: %s", err))
		return "", []*discordgo.MessageAttachment{}, err
	}

	return chanMsgs[0].Content, chanMsgs[0].Attachments, err
}

// messageSelector executes the 50-50 pseudo-randomized value to reply or not with a message for the received message on the channel.
func (b Bot) messageSelector(s *discordgo.Session, m *discordgo.MessageCreate) error {
	b.log.DebugLog(m.Author.Username)
	b.log.DebugLog(fmt.Sprintf("found message on thread: %s", m.Content))

	var err error
	if rand.Intn(100) > 50 {
		err = b.contentBasedInteraction(s, m)
		if err != nil {
			if err.Error() == "no_content_based_interaction_found" {
				err = b.authorBasedInteraction(s, m)
			}
		}

		if err != nil {
			if err.Error() == "no_author_based_interaction_found" {
				b.log.WarningLog("no message has been identified to be sent.")
				return nil
			}
		}
	}

	return err
}

// contentBasedInteraction based on messages received sends a new different message, regardless of who is the author.
func (b Bot) contentBasedInteraction(s *discordgo.Session, m *discordgo.MessageCreate) error {
	if strings.Contains(strings.ToLower(m.Content), "cs") {
		return b.csgoMessage(s, m)
	}

	if strings.Contains(strings.ToLower(m.Content), "pubg"){
		return b.pubgMessage(s, m)
	}

	if strings.Contains(strings.ToLower(m.Content), "bora") ||
		strings.Contains(strings.ToLower(m.Content), "online") ||
		strings.Contains(strings.ToLower(m.Content), "vamo") ||
		strings.Contains(strings.ToLower(m.Content), "voltei") {
		return b.boraMessage(s, m)
	}

	if strings.Contains(strings.ToLower(m.Content), "gamersclub") {
		return b.gamersclubMessage(s, m)
	}

	if strings.Contains(strings.ToLower(m.Content), "kim") {
		return b.kimMessage(s, m)
	}

	if strings.Contains(strings.ToLower(m.Content), "schenk") || strings.Contains(strings.ToLower(m.Content), "marcel") {
		return b.schenkMessage(s,m)
	}

	if strings.Contains(strings.ToLower(m.Content), "wolke") || strings.Contains(strings.ToLower(m.Content), "vitor"){
		return b.vitorMentionedMessage(s, m)
	}

	if strings.Contains(strings.ToLower(m.Content), "bob") {
		return b.bobMessage(s, m)
	}

	if strings.Contains(strings.ToLower(m.Content), "monstro"){
		return b.monstroMessage(s, m)
	}

	if strings.Contains(strings.ToLower(m.Content), "hess") || strings.Contains(strings.ToLower(m.Content), "hsz") || strings.Contains(strings.ToLower(m.Content), "geferson") {
		return b.hszMessage(s, m)
	}

	if strings.EqualFold(strings.ToLower(m.Content), "f") {
		return b.fMessage(s, m)
	}

	b.log.WarningLog("no content based interactions found")
	return fmt.Errorf("no_content_based_interaction_found")
}

// authorBasedInteraction based on who sent the message into the channel, replies back with some specific content.
func (b Bot) authorBasedInteraction(s *discordgo.Session, m *discordgo.MessageCreate) error {
	if strings.Contains(m.Author.Username, "chinela"){
		return b.chinelaMessage(s, m)
	}

	if strings.Contains(m.Author.Username, "vitor") {
		return b.vitorMessage(s, m)
	}

	if strings.Contains(m.Author.Username, "nico") {
		return b.nicoMessage(s, m)
	}

	b.log.WarningLog("no author based interactions found")
	return fmt.Errorf("no_author_based_interaction_found")
}

func (b Bot) executeActions(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.ToLower(m.Content) == "/stop_listening" {
		b.stopListening(m)
	} else if strings.ToLower(m.Content) == "/start_listening" {
		b.startListening(m)
	} else if strings.HasPrefix(strings.ToLower(m.Content), "/play") {
		b.callForGaming(s, m)
	}
}
