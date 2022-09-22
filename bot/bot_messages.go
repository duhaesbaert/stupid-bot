package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"math/rand"
)

func csgoMessage(s *discordgo.Session, m *discordgo.MessageCreate) error {
	message := "BORA JOGAR UM CSGO PORRA"
	log.InfoLog(fmt.Sprintf("sending message: %s", message))

	_, err := s.ChannelMessageSend(m.ChannelID, message)
	return err
}

func gamersclubMessage(s *discordgo.Session, m *discordgo.MessageCreate) error {
	message := "Gamers Club é muito coisa de try hard"
	log.InfoLog(fmt.Sprintf("sending message: %s", message))

	_, err := s.ChannelMessageSend(m.ChannelID, message)
	return err
}

func vitorMessage(s *discordgo.Session, m *discordgo.MessageCreate) error {
	message := "vitor = baiter"
	log.InfoLog(fmt.Sprintf("sending message: %s", message))
	_, err := s.ChannelMessageSend(m.ChannelID, message)
	if err != nil {
		return err
	}

	message = "ele tb nunca planta a bomba"
	log.InfoLog(fmt.Sprintf("sending message: %s", message))
	_, err = s.ChannelMessageSend(m.ChannelID, message)
	return err
}

func fMessage(s *discordgo.Session, m *discordgo.MessageCreate) error {
	message := "F"
	log.InfoLog(fmt.Sprintf("sending message: %s", message))
	_, err := s.ChannelMessageSend(m.ChannelID, message)
	return err
}

func bobMessage(s *discordgo.Session, m *discordgo.MessageCreate) error {
	message := "bob uma vez foi level 20 GC"
	log.InfoLog(fmt.Sprintf("sending message: %s", message))
	_, err := s.ChannelMessageSend(m.ChannelID, message)
	return err
}

func chinelaMessage(s *discordgo.Session, m *discordgo.MessageCreate) error {
	message := "De acordo com o leetify, chinela tem a maior taxa de amigos cegos por flash. Ótimo aproveitamento de utilitários."
	log.InfoLog(fmt.Sprintf("sending message: %s", message))
	_, err := s.ChannelMessageSend(m.ChannelID, message)
	return err
}

func monstroMessage(s *discordgo.Session, m *discordgo.MessageCreate) error {
	message := "SAI DA JAULA"
	log.InfoLog(fmt.Sprintf("sending message: %s", message))
	_, err := s.ChannelMessageSend(m.ChannelID, message)
	return err
}

func hszMessage(s *discordgo.Session, m *discordgo.MessageCreate) error {
	message := "\"vou só levar um colchao na sogra, volto em 15 min\""
	log.InfoLog(fmt.Sprintf("sending message: %s", message))
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
		log.InfoLog(fmt.Sprintf("sending message: %s", message))
		_, err := s.ChannelMessageSend(m.ChannelID, message)
		return err
	}

	log.WarningLog("a problem happened when identifying a message to send. No message has been sent")
	return nil
}
