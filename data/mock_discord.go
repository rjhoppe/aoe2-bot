package data

import (
	"github.com/bwmarrin/discordgo"
)

type MockSession struct {
	LastChannelID string
	LastMessage   string
	CallCount     int
}

func (m *MockSession) ChannelMessageSend(channelID, content string) (*discordgo.Message, error) {
	m.LastChannelID = channelID
	m.LastMessage = content
	m.CallCount++
	return &discordgo.Message{}, nil
}
