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

	if m.Author.Username == "chinela" {
		if m.Content == "/stop_listening" {
			b.stopListening()
		} else if m.Content == "/start_listening" {
			b.startListening()
		}
	}

	if b.botListening {
		if m.Content == "" {
			content, attachments, err := b.channelMessageLookup(s, m)
			if err != nil {
				b.log.ErrorLog(err.Error())
				return
			}

			b.log.InfoLog(content)
			m.Content = content
			m.Attachments = attachments
		}

		err := b.messageSelector(s, m)
		if err != nil {
			b.log.ErrorLog(fmt.Sprintf("could not sent message: %s", err.Error()))
		}
	}
}

func (b Bot) channelMessageLookup(s *discordgo.Session, m *discordgo.MessageCreate) (string, []*discordgo.MessageAttachment, error) {
	chanMsgs, err := s.ChannelMessages(m.ChannelID, 1, "", "", m.ID)
	if err != nil {
		b.log.ErrorLog(fmt.Sprintf("unable to get messages: %s", err))
		return "", []*discordgo.MessageAttachment{}, err
	}

	return chanMsgs[0].Content, chanMsgs[0].Attachments, err
}

func (b Bot) messageSelector(s *discordgo.Session, m *discordgo.MessageCreate) error {
	b.log.DebugLog(m.Author.Username)
	b.log.DebugLog(fmt.Sprintf("found message on thread: %s", m.Content))

	var err error
	if rand.Intn(100) > 50 {
		err = b.contentBasedInteraction(s, m)
		if err != nil {
			if err.Error() == "no_content_based_interaction_found" {
				err = b.authorBasedInteractions(s, m)
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

func (b Bot) contentBasedInteraction(s *discordgo.Session, m *discordgo.MessageCreate) error {
	if strings.Contains(m.Content, "cs") {
		return b.csgoMessage(s, m)
	}

	if strings.Contains(m.Content, "pubg"){
		return b.pubgMessage(s, m)
	}

	if strings.Contains(m.Content, "bora") ||
		strings.Contains(m.Content, "online") ||
		strings.Contains(m.Content, "vamo") ||
		strings.Contains(m.Content, "voltei") {
		return b.boraMessage(s, m)
	}

	if strings.Contains(m.Content, "gamersclub") {
		return b.gamersclubMessage(s, m)
	}

	if strings.Contains(m.Content, "wolke") || strings.Contains(m.Content, "vitor"){
		return b.vitorMentionedMessage(s, m)
	}

	if strings.Contains(m.Content, "bob") {
		return b.bobMessage(s, m)
	}

	if strings.Contains(m.Content, "monstro"){
		return b.monstroMessage(s, m)
	}

	if strings.Contains(m.Content, "hess") || strings.Contains(m.Content, "hsz") || strings.Contains(m.Content, "geferson") {
		return b.hszMessage(s, m)
	}

	if strings.EqualFold(m.Content, "F") || strings.EqualFold(m.Content, "f") {
		return b.fMessage(s, m)
	}

	b.log.WarningLog("no content based interactions found")
	return fmt.Errorf("no_content_based_interaction_found")
}

func (b Bot) authorBasedInteractions(s *discordgo.Session, m *discordgo.MessageCreate) error {
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