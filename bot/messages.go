package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"math/rand"
)

func (b Bot) csgoMessage(s *discordgo.Session, m *discordgo.MessageCreate) error {
	message := "BORA JOGAR UM CSGO PORRA"
	b.log.InfoLog(fmt.Sprintf("sending message: %s", message))
	_, err := s.ChannelMessageSendReply(m.ChannelID, message, &discordgo.MessageReference{ChannelID: m.ChannelID, MessageID: m.Message.ID})
	return err
}

func (b Bot) gamersclubMessage(s *discordgo.Session, m *discordgo.MessageCreate) error {
	message := "try hard"
	b.log.InfoLog(fmt.Sprintf("sending message: %s", message))
	_, err := s.ChannelMessageSendReply(m.ChannelID, message, &discordgo.MessageReference{ChannelID: m.ChannelID, MessageID: m.Message.ID})
	return err
}

func (b Bot) vitorMessage(s *discordgo.Session, m *discordgo.MessageCreate) error {
	message := "baiter"
	b.log.InfoLog(fmt.Sprintf("sending message: %s", message))
	_, err := s.ChannelMessageSendReply(m.ChannelID, message, &discordgo.MessageReference{ChannelID: m.ChannelID, MessageID: m.Message.ID})
	return err
}

func (b Bot) vitorMentionedMessage(s *discordgo.Session, m *discordgo.MessageCreate) error {
	message := "esse ai nunca planta a bomba"
	b.log.InfoLog(fmt.Sprintf("sending message: %s", message))
	_, err := s.ChannelMessageSendReply(m.ChannelID, message, &discordgo.MessageReference{ChannelID: m.ChannelID, MessageID: m.Message.ID})
	return err
}

func (b Bot) fMessage(s *discordgo.Session, m *discordgo.MessageCreate) error {
	message := "https://tenor.com/view/keyboard-hyperx-rgb-hyperx-family-hyperx-gaming-gif-17743649"
	b.log.InfoLog(fmt.Sprintf("sending message: %s", message))
	_, err := s.ChannelMessageSendReply(m.ChannelID, message, &discordgo.MessageReference{ChannelID: m.ChannelID, MessageID: m.Message.ID})
	return err
}

func (b Bot) bobMessage(s *discordgo.Session, m *discordgo.MessageCreate) error {
	message := "bob uma vez foi level 20 GC"
	b.log.InfoLog(fmt.Sprintf("sending message: %s", message))
	_, err := s.ChannelMessageSendReply(m.ChannelID, message, &discordgo.MessageReference{ChannelID: m.ChannelID, MessageID: m.Message.ID})
	return err
}

func (b Bot) chinelaMessage(s *discordgo.Session, m *discordgo.MessageCreate) error {
	message := "De acordo com o leetify, chinela tem a maior taxa de amigos cegos por flash. Ótimo aproveitamento de utilitários."
	b.log.InfoLog(fmt.Sprintf("sending message: %s", message))
	_, err := s.ChannelMessageSendReply(m.ChannelID, message, &discordgo.MessageReference{ChannelID: m.ChannelID, MessageID: m.Message.ID})
	return err
}

func (b Bot) monstroMessage(s *discordgo.Session, m *discordgo.MessageCreate) error {
	message := "SAI DA JAULA"
	b.log.InfoLog(fmt.Sprintf("sending message: %s", message))
	_, err := s.ChannelMessageSend(m.ChannelID, message)
	return err
}

func (b Bot) hszMessage(s *discordgo.Session, m *discordgo.MessageCreate) error {
	message := "\"vou só levar um colchao na sogra, volto em 15 min\""
	b.log.InfoLog(fmt.Sprintf("sending message: %s", message))
	_, err := s.ChannelMessageSendReply(m.ChannelID, message, &discordgo.MessageReference{ChannelID: m.ChannelID, MessageID: m.Message.ID})
	return err
}

func (b Bot) boraMessage(s *discordgo.Session, m *discordgo.MessageCreate) error {
	message := ""
	max := 100
	min := 1
	randomNum := rand.Intn(max - min) + min

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
		_, err := s.ChannelMessageSendReply(m.ChannelID, message, &discordgo.MessageReference{ChannelID: m.ChannelID, MessageID: m.Message.ID})
		return err
	}

	b.log.WarningLog("a problem happened when identifying a message to send. No message has been sent")
	return nil
}

func (b Bot) nicoMessage(s *discordgo.Session, m *discordgo.MessageCreate) error {
	message := "https://tenor.com/view/communiste-communist-hugs-heart-red-gif-14360509"
	b.log.InfoLog(fmt.Sprintf("sending message: %s", message))
	_, err := s.ChannelMessageSendReply(m.ChannelID, message, &discordgo.MessageReference{ChannelID: m.ChannelID, MessageID: m.Message.ID})
	return err
}

func (b Bot) pubgMessage(s *discordgo.Session, m *discordgo.MessageCreate) error {
	message := "jogo de nerdola"
	b.log.InfoLog(fmt.Sprintf("sending message: %s", message))
	_, err := s.ChannelMessageSendReply(m.ChannelID, message, &discordgo.MessageReference{ChannelID: m.ChannelID, MessageID: m.Message.ID})
	return err
}

func (b Bot) kimMessage(s *discordgo.Session, m *discordgo.MessageCreate) error {
	message := "nosso muambeiro favorito"
	b.log.InfoLog(fmt.Sprintf("sending message: %s", message))
	_, err := s.ChannelMessageSendReply(m.ChannelID, message, &discordgo.MessageReference{ChannelID: m.ChannelID, MessageID: m.Message.ID})
	return err
}

func (b Bot) schenkMessage(s *discordgo.Session, m *discordgo.MessageCreate) error {
	message := "foi pra germania e nos esqueceu"
	b.log.InfoLog(fmt.Sprintf("sending message: %s", message))
	_, err := s.ChannelMessageSendReply(m.ChannelID, message, &discordgo.MessageReference{ChannelID: m.ChannelID, MessageID: m.Message.ID})
	return err
}