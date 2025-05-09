package data

import (
	"github.com/bwmarrin/discordgo"
)

type MockSession struct {
	LastChannelID string
	LastMessage   string
	CallCount     int
}

type DiscordSession interface {
	ChannelMessageSend(channelID string, content string, options ...discordgo.RequestOption) (*discordgo.Message, error)
}

func (m *MockSession) ChannelMessageSend(channelID string, content string, options ...discordgo.RequestOption) (*discordgo.Message, error) {
	m.LastChannelID = channelID
	m.LastMessage = content
	m.CallCount++
	return &discordgo.Message{}, nil
}
