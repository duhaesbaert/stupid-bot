package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"math/rand"
	"strings"
	"stupid-bot/common"
)

type BotMessages struct {
	bot Bot
	log common.Logger
	s   *discordgo.Session
	m   *discordgo.MessageCreate
}

// messageSelector executes the 50-50 pseudo-randomized value to reply or not with a message for the received message on the channel.
func (b BotMessages) messageSelector() error {
	b.log.DebugLog(b.m.Author.Username)
	b.log.DebugLog(fmt.Sprintf("found message on thread: %s", b.m.Content))

	var err error
	if rand.Intn(100) > 50 {
		err = b.contentBasedInteraction()
		if err != nil {
			if err.Error() == "no_content_based_interaction_found" {
				err = b.authorBasedInteraction()
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
func (b BotMessages) contentBasedInteraction() error {
	if strings.Contains(strings.ToLower(b.m.Content), "cs") {
		return b.csgoMessage()
	}

	if strings.Contains(strings.ToLower(b.m.Content), "pubg") {
		return b.pubgMessage()
	}

	if strings.Contains(strings.ToLower(b.m.Content), "bora") ||
		strings.Contains(strings.ToLower(b.m.Content), "online") ||
		strings.Contains(strings.ToLower(b.m.Content), "vamo") ||
		strings.Contains(strings.ToLower(b.m.Content), "voltei") {
		return b.boraMessage()
	}

	if strings.Contains(strings.ToLower(b.m.Content), "gamersclub") {
		return b.gamersclubMessage()
	}

	if strings.Contains(strings.ToLower(b.m.Content), "kim") {
		return b.kimMessage()
	}

	if strings.Contains(strings.ToLower(b.m.Content), "schenk") || strings.Contains(strings.ToLower(b.m.Content), "marcel") {
		return b.schenkMessage()
	}

	if strings.Contains(strings.ToLower(b.m.Content), "wolke") || strings.Contains(strings.ToLower(b.m.Content), "vitor") {
		return b.vitorMentionedMessage()
	}

	if strings.Contains(strings.ToLower(b.m.Content), "bob") {
		return b.bobMessage()
	}

	if strings.Contains(strings.ToLower(b.m.Content), "monstro") {
		return b.monstroMessage()
	}

	if strings.Contains(strings.ToLower(b.m.Content), "hess") || strings.Contains(strings.ToLower(b.m.Content), "hsz") || strings.Contains(strings.ToLower(b.m.Content), "geferson") {
		return b.hszMessage()
	}

	if strings.EqualFold(strings.ToLower(b.m.Content), "f") {
		return b.fMessage()
	}

	b.log.WarningLog("no content based interactions found")
	return fmt.Errorf("no_content_based_interaction_found")
}

// authorBasedInteraction based on who sent the message into the channel, reply back with some specific content.
func (b BotMessages) authorBasedInteraction() error {
	if strings.Contains(b.m.Author.Username, "chinela") {
		return b.chinelaMessage()
	}

	if strings.Contains(b.m.Author.Username, "vitor") {
		return b.vitorMessage()
	}

	if strings.Contains(b.m.Author.Username, "nico") {
		return b.nicoMessage()
	}

	b.log.WarningLog("no author based interactions found")
	return fmt.Errorf("no_author_based_interaction_found")
}

func (b BotMessages) csgoMessage() error {
	message := "BORA JOGAR UM CSGO PORRA"
	b.log.InfoLog(fmt.Sprintf("sending message: %s", message))
	_, err := b.s.ChannelMessageSendReply(b.m.ChannelID, message, &discordgo.MessageReference{ChannelID: b.m.ChannelID, MessageID: b.m.Message.ID})
	return err
}

func (b BotMessages) gamersclubMessage() error {
	message := "try hard"
	b.log.InfoLog(fmt.Sprintf("sending message: %s", message))
	_, err := b.s.ChannelMessageSendReply(b.m.ChannelID, message, &discordgo.MessageReference{ChannelID: b.m.ChannelID, MessageID: b.m.Message.ID})
	return err
}

func (b BotMessages) vitorMessage() error {
	message := "baiter"
	b.log.InfoLog(fmt.Sprintf("sending message: %s", message))
	_, err := b.s.ChannelMessageSendReply(b.m.ChannelID, message, &discordgo.MessageReference{ChannelID: b.m.ChannelID, MessageID: b.m.Message.ID})
	return err
}

func (b BotMessages) vitorMentionedMessage() error {
	message := "esse ai nunca planta a bomba"
	b.log.InfoLog(fmt.Sprintf("sending message: %s", message))
	_, err := b.s.ChannelMessageSendReply(b.m.ChannelID, message, &discordgo.MessageReference{ChannelID: b.m.ChannelID, MessageID: b.m.Message.ID})
	return err
}

func (b BotMessages) fMessage() error {
	message := "https://tenor.com/view/keyboard-hyperx-rgb-hyperx-family-hyperx-gaming-gif-17743649"
	b.log.InfoLog(fmt.Sprintf("sending message: %s", message))
	_, err := b.s.ChannelMessageSendReply(b.m.ChannelID, message, &discordgo.MessageReference{ChannelID: b.m.ChannelID, MessageID: b.m.Message.ID})
	return err
}

func (b BotMessages) bobMessage() error {
	message := "bob uma vez foi level 20 GC"
	b.log.InfoLog(fmt.Sprintf("sending message: %s", message))
	_, err := b.s.ChannelMessageSendReply(b.m.ChannelID, message, &discordgo.MessageReference{ChannelID: b.m.ChannelID, MessageID: b.m.Message.ID})
	return err
}

func (b BotMessages) chinelaMessage() error {
	message := "De acordo com o leetify, chinela tem a maior taxa de amigos cegos por flash. Ótimo aproveitamento de utilitários."
	b.log.InfoLog(fmt.Sprintf("sending message: %s", message))
	_, err := b.s.ChannelMessageSendReply(b.m.ChannelID, message, &discordgo.MessageReference{ChannelID: b.m.ChannelID, MessageID: b.m.Message.ID})
	return err
}

func (b BotMessages) monstroMessage() error {
	message := "SAI DA JAULA"
	b.log.InfoLog(fmt.Sprintf("sending message: %s", message))
	_, err := b.s.ChannelMessageSend(b.m.ChannelID, message)
	return err
}

func (b BotMessages) hszMessage() error {
	message := "\"vou só levar um colchao na sogra, volto em 15 min\""
	b.log.InfoLog(fmt.Sprintf("sending message: %s", message))
	_, err := b.s.ChannelMessageSendReply(b.m.ChannelID, message, &discordgo.MessageReference{ChannelID: b.m.ChannelID, MessageID: b.m.Message.ID})
	return err
}

func (b BotMessages) boraMessage() error {
	message := ""
	max := 100
	min := 1
	randomNum := rand.Intn(max-min) + min

	if randomNum < 25 {
		message = "csgo > valorant"
	}

	if randomNum > 25 && randomNum < 50 {
		message = "vitor disse que em 5 min ta entrando"
	}

	if randomNum > 50 && randomNum < 75 {
		message = "\"hoje tem\""
	}

	if randomNum > 75 {
		message = "to on."
	}

	if message != "" {
		b.log.InfoLog(fmt.Sprintf("sending message: %s", message))
		_, err := b.s.ChannelMessageSendReply(b.m.ChannelID, message, &discordgo.MessageReference{ChannelID: b.m.ChannelID, MessageID: b.m.Message.ID})
		return err
	}

	b.log.WarningLog("a problem happened when identifying a message to send. No message has been sent")
	return nil
}

func (b BotMessages) nicoMessage() error {
	message := "https://tenor.com/view/communiste-communist-hugs-heart-red-gif-14360509"
	b.log.InfoLog(fmt.Sprintf("sending message: %s", message))
	_, err := b.s.ChannelMessageSendReply(b.m.ChannelID, message, &discordgo.MessageReference{ChannelID: b.m.ChannelID, MessageID: b.m.Message.ID})
	return err
}

func (b BotMessages) pubgMessage() error {
	message := "jogo de nerdola"
	b.log.InfoLog(fmt.Sprintf("sending message: %s", message))
	_, err := b.s.ChannelMessageSendReply(b.m.ChannelID, message, &discordgo.MessageReference{ChannelID: b.m.ChannelID, MessageID: b.m.Message.ID})
	return err
}

func (b BotMessages) kimMessage() error {
	message := "nosso muambeiro favorito"
	b.log.InfoLog(fmt.Sprintf("sending message: %s", message))
	_, err := b.s.ChannelMessageSendReply(b.m.ChannelID, message, &discordgo.MessageReference{ChannelID: b.m.ChannelID, MessageID: b.m.Message.ID})
	return err
}

func (b BotMessages) schenkMessage() error {
	message := "foi pra germania e nos esqueceu"
	b.log.InfoLog(fmt.Sprintf("sending message: %s", message))
	_, err := b.s.ChannelMessageSendReply(b.m.ChannelID, message, &discordgo.MessageReference{ChannelID: b.m.ChannelID, MessageID: b.m.Message.ID})
	return err
}
