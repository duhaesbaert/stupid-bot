package bot

import (
	"github.com/bwmarrin/discordgo"
	"math/rand"
)

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

func fMessage(s *discordgo.Session, m *discordgo.MessageCreate) error {
	_, err := s.ChannelMessageSend(m.ChannelID, "F")
	return err
}

func bobMessage(s *discordgo.Session, m *discordgo.MessageCreate) error {
	_, err := s.ChannelMessageSend(m.ChannelID, "bob uma vez foi level 20 GC")
	return err
}

func chinelaMessage(s *discordgo.Session, m *discordgo.MessageCreate) error {
	_, err := s.ChannelMessageSend(m.ChannelID, "De acordo com o leetify, chinela tem a maior taxa de amigos cegos por flash. Ótimo aproveitamento de utilitários.")
	return err
}

func monstroMessage(s *discordgo.Session, m *discordgo.MessageCreate) error {
	_, err := s.ChannelMessageSend(m.ChannelID, "SAI DA JAULA")
	return err
}

func hszMessage(s *discordgo.Session, m *discordgo.MessageCreate) error {
	_, err := s.ChannelMessageSend(m.ChannelID, "\"vou só levar um colchao na sogra, volto em 15 min\"")
	return err
}

func boraMessage(s *discordgo.Session, m *discordgo.MessageCreate) error {
	max := 100
	min := 1

	randomNum := rand.Intn(max - min) + min

	if randomNum < 25 {
		_, err := s.ChannelMessageSend(m.ChannelID, "to on.")
		return err
	}

	if randomNum > 25 && randomNum < 50 {
		_, err := s.ChannelMessageSend(m.ChannelID, "vitor disse que em 5 min ta entrando")
		return err
	}

	if randomNum > 50 && randomNum < 75 {
		_, err := s.ChannelMessageSend(m.ChannelID, "\"nunca mais jogo cs\"")
		return err
	}

	if randomNum > 75 {
		_, err := s.ChannelMessageSend(m.ChannelID, "hoje tem")
		return err
	}

	return nil
}
