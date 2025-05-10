package utils

import (
	"strings"
	"testing"

	"github.com/bwmarrin/discordgo"
)

func TestSelectRandomArrayEle(t *testing.T) {
	arr := []string{"a", "b", "c"}
	result := SelectRandomArrayEle(arr)
	found := false
	for _, v := range arr {
		if v == result {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Expected result to be one of %v, got %v", arr, result)
	}
}

func TestGetCurDate(t *testing.T) {
	date := GetCurDate()
	if date == "" {
		t.Error("Expected non-empty date string")
	}
	// Optionally, check format: MM/DD/YYYY
	if len(date) != 10 || date[2] != '/' || date[5] != '/' {
		t.Errorf("Expected date in MM/DD/YYYY format, got %v", date)
	}
}

func TestFirstCharToUpper(t *testing.T) {
	tests := map[string]string{
		"hello": "Hello",
		"Hello": "Hello",
		"":      "",
		"a":     "A",
	}
	for input, expected := range tests {
		if got := FirstCharToUpper(input); got != expected {
			t.Errorf("FirstCharToUpper(%q) = %q; want %q", input, got, expected)
		}
	}
}

func TestCheckMsgForExclamation(t *testing.T) {
	if !CheckMsgForExclamation("!cmd") {
		t.Error("Expected true for string with '!'")
	}
	if CheckMsgForExclamation("cmd") {
		t.Error("Expected false for string without '!'")
	}
}

type MockSession struct {
	LastChannelID string
	LastMessage   string
	CallCount     int
}

func (m *MockSession) ChannelMessageSend(channelID string, content string, options ...discordgo.RequestOption) (*discordgo.Message, error) {
	m.LastChannelID = channelID
	m.LastMessage = content
	m.CallCount++
	return &discordgo.Message{}, nil
}

func TestIsValidCmd(t *testing.T) {
	mockSession := &MockSession{}
	mockMessage := &discordgo.MessageCreate{
		Message: &discordgo.Message{
			Content:   "short",
			ChannelID: "test-channel",
		},
	}
	// Should be invalid (too short)
	if IsValidCmd(10, mockSession, mockMessage) {
		t.Error("Expected IsValidCmd to return false for short input")
	}
	if mockSession.CallCount != 1 {
		t.Error("Expected ChannelMessageSend to be called for invalid input")
	}

	// Should be valid
	mockSession = &MockSession{}
	mockMessage.Message.Content = "this is long enough"
	if !IsValidCmd(10, mockSession, mockMessage) {
		t.Error("Expected IsValidCmd to return true for valid input")
	}
	if mockSession.CallCount != 0 {
		t.Error("Expected ChannelMessageSend not to be called for valid input")
	}
}

func TestPrintCmds(t *testing.T) {
	mockSession := &MockSession{}
	mockMessage := &discordgo.MessageCreate{
		Message: &discordgo.Message{
			ChannelID: "test-channel",
			Content:   "!cmds",
		},
	}

	PrintCmds(mockSession, mockMessage)

	if mockSession.CallCount != 1 {
		t.Errorf("Expected ChannelMessageSend to be called once, got %d", mockSession.CallCount)
	}
	if mockSession.LastChannelID != "test-channel" {
		t.Errorf("Expected ChannelID to be 'test-channel', got %s", mockSession.LastChannelID)
	}
	if !strings.Contains(mockSession.LastMessage, "!civ") {
		t.Errorf("Expected command list to contain '!civ', got %q", mockSession.LastMessage)
	}
	if !strings.Contains(mockSession.LastMessage, "!leaderboard") {
		t.Errorf("Expected command list to contain '!leaderboard', got %q", mockSession.LastMessage)
	}
}
