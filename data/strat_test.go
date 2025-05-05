package data

import (
	"testing"
)

func TestFormatStratOutput(t *testing.T) {
	mockSession := &MockSession{}
	mockMessage := &discordgo.MessageCreate{
		Message: &discordgo.Message{
			Content: "!strat drush",
			ChannelID: "test-channel",
		},
	}
	FormatStratOutput(mockSession, mockMessage)

	if mockSession.CallCount != 1 {
		t.Errorf("Expected ChannelMessageSend to be called once, got %d", mockSession.CallCount)
	}
	if mockSession.LastChannelID != "test-channel" {
		t.Errorf("Expected channel ID 'test-channel', got %s", mockSession.LastChannelID)
	}
	if mockSession.LastMessage == "" {
		t.Error("Expected a message to be sent, got empty string")
	}
	// Optionally, check for specific content in mockSession.LastMessage
	if !strings.Contains(mockSession.LastMessage, "drush") {
		t.Errorf("Expected message to contain 'drush', got %s", mockSession.LastMessage)
	}
}

func TestStrategiesInfo_TurtleIsPopulated(t *testing.T) {
	strategy, ok := StrategiesInfo["turtle"]
	if !ok {
		t.Fatal("Expected 'drush' strategy to exist in StrategiesInfo")
	}
	if strategy.Name == "" {
		t.Error("Expected strategy.Name to be set, got empty string")
	}
	if strategy.Emoji == "" {
		t.Error("Expected strategy.Emoji to be set, got empty string")
	}
	if strategy.Description == "" {
		t.Error("Expected strategy.Description to be set, got empty string")
	}
	if strategy.Pros == "" {
		t.Error("Expected strategy.Pros to be set, got empty string")
	}
	if strategy.Cons == "" {
		t.Error("Expected strategy.Cons to be set, got empty string")
	}
	if strategy.Tips == "" {
		t.Error("Expected strategy.Tips to be set, got empty string")
	}
	if len(strategy.Civs) == 0 {
		t.Error("Expected strategy.Civs to be set, got empty slice")
	}
}

func TestListAllStrats(t *testing.T) {
	mockSession := &MockSession{}
	mockMessage := &discordgo.MessageCreate{
		Message: &discordgo.Message{
			Content: "!stratlist",		
			ChannelID: "test-channel",
		},
	}
	ListAllStrats(mockSession, mockMessage)

	if mockSession.CallCount != 1 {
		t.Errorf("Expected ChannelMessageSend to be called once, got %d", mockSession.CallCount)
	}	
	if mockSession.LastChannelID != "test-channel" {
		t.Errorf("Expected channel ID 'test-channel', got %s", mockSession.LastChannelID)
	}
	if mockSession.LastMessage == "" {
		t.Error("Expected a message to be sent, got empty string")
	}
	if !strings.Contains(mockSession.LastMessage, "drush") {
		t.Errorf("Expected message to contain 'drush', got %s", mockSession.LastMessage)
	}
}



func TestFormatStratOutputUserInputTooShort(t *testing.T) {
	// Save the original function and restore after test
	origIsValidCmd := utils.IsValidCmd
	defer func() { utils.IsValidCmd = origIsValidCmd }()

	// Mock IsValidCmd to always return false
	utils.IsValidCmd = func(_ int, _ *discordgo.Session, _ *discordgo.MessageCreate) bool {
		return false
	}
	
	mockSession := &MockSession{}
	mockMessage := &discordgo.MessageCreate{
		Message: &discordgo.Message{
			Content: "!strat",		
			ChannelID: "test-channel",
		},
	}
	FormatStratOutput(mockSession, mockMessage)

	if mockSession.CallCount != 0 {
		t.Errorf("Expected ChannelMessageSend to not be called, but it was called %d times", mockSession.CallCount)
	}
}

func TestFormatStratOutputUserInputInvalid(t *testing.T) {
	mockSession := &MockSession{}
	mockMessage := &discordgo.MessageCreate{
		Message: &discordgo.Message{
			Content: "!strat flush",		
			ChannelID: "test-channel",
		},
	}
	FormatStratOutput(mockSession, mockMessage)

	if mockSession.CallCount != 1 {
		t.Errorf("Expected ChannelMessageSend to be called once, got %d", mockSession.CallCount)
	}	
	if mockSession.LastMessage == "" {
		t.Error("Expected a message to be sent, got empty string")
	}
}