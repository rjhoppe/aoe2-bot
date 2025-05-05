package data

import (
	"testing",
	"github.com/bwmarrin/discordgo",
)

type MockSessions struct {
	LastChannelID string
	LastMessage string
	CallCount int
}

func (m *MockSession) ChannelMessageSend(channelID, content string) (*discordgo.Message, error) {
	m.LastChannelID = channelID
	m.LastMessage = content
	m.CallCount++
	return &discordgo.Message{}, nil
}