package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"math/rand"
	"stupid-bot/common"
)

func csgoMessage(s *discordgo.Session, m *discordgo.MessageCreate) error {
	message := "BORA JOGAR UM CSGO PORRA"
	common.InfoLog(fmt.Sprintf("sending message: %s", message))

	_, err := s.ChannelMessageSend(m.ChannelID, message)
	return err
}

func gamersclubMessage(s *discordgo.Session, m *discordgo.MessageCreate) error {
	message := "Gamers Club é muito coisa de try hard"
	common.InfoLog(fmt.Sprintf("sending message: %s", message))

	_, err := s.ChannelMessageSend(m.ChannelID, message)
	return err
}

func vitorMessage(s *discordgo.Session, m *discordgo.MessageCreate) error {
	message := "vitor = baiter"
	common.InfoLog(fmt.Sprintf("sending message: %s", message))
	_, err := s.ChannelMessageSend(m.ChannelID, message)
	if err != nil {
		return err
	}

	message = "ele tb nunca planta a bomba"
	common.InfoLog(fmt.Sprintf("sending message: %s", message))
	_, err = s.ChannelMessageSend(m.ChannelID, message)
	return err
}

func fMessage(s *discordgo.Session, m *discordgo.MessageCreate) error {
	message := "F"
	common.InfoLog(fmt.Sprintf("sending message: %s", message))
	_, err := s.ChannelMessageSend(m.ChannelID, message)
	return err
}

func bobMessage(s *discordgo.Session, m *discordgo.MessageCreate) error {
	message := "bob uma vez foi level 20 GC"
	common.InfoLog(fmt.Sprintf("sending message: %s", message))
	_, err := s.ChannelMessageSend(m.ChannelID, message)
	return err
}

func chinelaMessage(s *discordgo.Session, m *discordgo.MessageCreate) error {
	message := "De acordo com o leetify, chinela tem a maior taxa de amigos cegos por flash. Ótimo aproveitamento de utilitários."
	common.InfoLog(fmt.Sprintf("sending message: %s", message))
	_, err := s.ChannelMessageSend(m.ChannelID, message)
	return err
}

func monstroMessage(s *discordgo.Session, m *discordgo.MessageCreate) error {
	message := "SAI DA JAULA"
	common.InfoLog(fmt.Sprintf("sending message: %s", message))
	_, err := s.ChannelMessageSend(m.ChannelID, message)
	return err
}

func hszMessage(s *discordgo.Session, m *discordgo.MessageCreate) error {
	message := "\"vou só levar um colchao na sogra, volto em 15 min\""
	common.InfoLog(fmt.Sprintf("sending message: %s", message))
	_, err := s.ChannelMessageSend(m.ChannelID, message)
	return err
}

func boraMessage(s *discordgo.Session, m *discordgo.MessageCreate) error {
	message := ""
	max := 100
	min := 1
	randomNum := rand.Intn(max - min) + min

	if randomNum < 25 {
		message = "to on."
	}

	if randomNum > 25 && randomNum < 50 {
		message = "vitor disse que em 5 min ta entrando"
	}

	if randomNum > 50 && randomNum < 75 {
		message = "\"nunca mais jogo cs\""
	}

	if randomNum > 75 {
		message = "hoje tem"
	}

	if message != "" {
		common.InfoLog(fmt.Sprintf("sending message: %s", message))
		_, err := s.ChannelMessageSend(m.ChannelID, message)
		return err
	}

	common.WarningLog("a problem happened when identifying a message to send. No message has been sent")
	return nil
}
