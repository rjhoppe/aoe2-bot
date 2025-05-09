package data

import (
	"strings"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/rjhoppe/aoe-bot/utils"
)

func TestGetNewRandomCiv_All(t *testing.T) {
	civ := GetNewRandomCiv("all")
	if civ == nil {
		t.Fatal("Expected a Civilization, got nil")
	}
	if civ.Name == "" || civ.Type == "" {
		t.Errorf("Expected civ to have Name and Type, got %+v", civ)
	}
}

func TestGetNewRandomCiv_HasAllProperties(t *testing.T) {
	civ := GetNewRandomCiv("all")
	if civ == nil {
		t.Fatal("Expected a Civilization, got nil")
	}
	if civ.Name == "" {
		t.Error("Expected civ.Name to be set, got empty string")
	}
	if civ.Type == "" {
		t.Error("Expected civ.Type to be set, got empty string")
	}
	if CivTypeToEmoji[civ.Type] == "" {
		t.Errorf("Expected civ.Type to be set, got empty string")
	}
	if civ.Strengths == "" {
		t.Error("Expected civ.Strengths to be set, got empty string")
	}
	if civ.Weaknesses == "" {
		t.Error("Expected civ.Weaknesses to be set, got empty string")
	}
}

func TestGetNewRandomCiv_Archer(t *testing.T) {
	civ := GetNewRandomCiv("archer")
	if civ == nil {
		t.Fatal("Expected a Civilization, got nil")
	}

	if civ.Name == "" || civ.Type == "" {
		t.Errorf("Expected civ to have Name and Type, got %+v", civ)
	}
	found := false
	for _, name := range archerCivs {
		if civ.Name == name {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Expected archer civ, got %v", civ.Name)
	}
}

func TestGetNewRandomCiv_Cavalry(t *testing.T) {
	civ := GetNewRandomCiv("cavalry")
	if civ == nil {
		t.Fatal("Expected a Civilization, got nil")
	}

	if civ.Name == "" || civ.Type == "" {
		t.Errorf("Expected civ to have Name and Type, got %+v", civ)
	}
	found := false
	for _, name := range cavCivs {
		if civ.Name == name {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Expected cavalry civ, got %v", civ.Name)
	}
}

func TestPrintCivOutput(t *testing.T) {
	mockSession := &MockSession{}
	mockCiv := Civilization{
		Name:       "Britons",
		Type:       "Foot archer",
		Strengths:  "Flush archer, Crossbowmen rush",
		Weaknesses: "Weak cavalry, Easily countered with Siege Rams, Reliant on Trebuchets for siege, Mediocre archer options outside of Longbowmen",
	}
	mockMessage := &discordgo.MessageCreate{
		Message: &discordgo.Message{
			Content:   "!civ",
			ChannelID: "test-channel",
		},
	}
	PrintCivOutput("all", &mockCiv, mockSession, mockMessage)

	if mockSession.LastChannelID == "" || mockSession.LastMessage == "" {
		t.Errorf("Expected ChannelMessageSend to be called, but it wasn't")
	}
	if mockSession.CallCount != 1 {
		t.Errorf("Expected ChannelMessageSend to be called once, but it was called %d times", mockSession.CallCount)
	}
}

func TestGetThreeRandomCivs(t *testing.T) {
	mockSession := &MockSession{}
	mockMessage := &discordgo.MessageCreate{
		Message: &discordgo.Message{
			Content:   "!civ",
			ChannelID: "test-channel",
		},
	}
	GetThreeRandomCivs(mockSession, mockMessage)
	if mockSession.LastChannelID == "" || mockSession.LastMessage == "" {
		t.Errorf("Expected ChannelMessageSend to be called, but it wasn't")
	}
	if mockSession.CallCount != 1 {
		t.Errorf("Expected ChannelMessageSend to be called once, but it was called %d times", mockSession.CallCount)
	}

	civs := strings.Split(mockSession.LastMessage, ",")
	if len(civs) != 3 {
		t.Errorf("Expected 3 civs, got %d", len(civs))
	}
	for _, civ := range civs {
		civ = strings.TrimSpace(civ)
		if civ == "" {
			t.Errorf("Expected a civ, got empty string")
		}
		found := false
		for _, name := range allCivs {
			if civ == name {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected a valid civ, got %v", civ)
		}
	}
}

func TestListAllStrengths_ValidCiv(t *testing.T) {
	// Optionally, mock IsValidCmd to always return true
	origIsValidCmd := utils.IsValidCmd
	utils.IsValidCmd = func(_ int, _ utils.DiscordSession, _ *discordgo.MessageCreate) bool { return true }
	defer func() { utils.IsValidCmd = origIsValidCmd }()

	mockSession := &MockSession{}
	mockMessage := &discordgo.MessageCreate{
		Message: &discordgo.Message{
			Content:   "!civstrat Britons", // "Britons" is a valid civ
			ChannelID: "test-channel",
		},
	}
	ListAllStrengths(mockSession, mockMessage)

	if mockSession.CallCount != 1 {
		t.Errorf("Expected ChannelMessageSend to be called once, got %d", mockSession.CallCount)
	}
	if !strings.Contains(mockSession.LastMessage, "Britons") {
		t.Errorf("Expected message to contain 'Britons', got: %s", mockSession.LastMessage)
	}
}

// Test for invalid command (should not send a message)
func TestListAllStrengths_InvalidCmd(t *testing.T) {
	origIsValidCmd := utils.IsValidCmd
	utils.IsValidCmd = func(_ int, _ utils.DiscordSession, _ *discordgo.MessageCreate) bool { return false }
	defer func() { utils.IsValidCmd = origIsValidCmd }()

	mockSession := &MockSession{}
	mockMessage := &discordgo.MessageCreate{
		Message: &discordgo.Message{
			Content:   "!civstrat",
			ChannelID: "test-channel",
		},
	}
	ListAllStrengths(mockSession, mockMessage)

	if mockSession.CallCount != 0 {
		t.Errorf("Expected ChannelMessageSend to not be called, but it was called %d times", mockSession.CallCount)
	}
}

// Test for unknown civ (should not send a message)
func TestListAllStrengths_UnknownCiv(t *testing.T) {
	origIsValidCmd := utils.IsValidCmd
	utils.IsValidCmd = func(_ int, _ utils.DiscordSession, _ *discordgo.MessageCreate) bool { return true }
	defer func() { utils.IsValidCmd = origIsValidCmd }()

	mockSession := &MockSession{}
	mockMessage := &discordgo.MessageCreate{
		Message: &discordgo.Message{
			Content:   "!civstrat NotACiv",
			ChannelID: "test-channel",
		},
	}
	ListAllStrengths(mockSession, mockMessage)

	if mockSession.CallCount != 0 {
		t.Errorf("Expected ChannelMessageSend to not be called for unknown civ, but it was called %d times", mockSession.CallCount)
	}
}

func TestListAllStrengths_Lowercase(t *testing.T) {
	origIsValidCmd := utils.IsValidCmd
	utils.IsValidCmd = func(_ int, _ utils.DiscordSession, _ *discordgo.MessageCreate) bool { return true }
	defer func() { utils.IsValidCmd = origIsValidCmd }()

	mockSession := &MockSession{}
	mockMessage := &discordgo.MessageCreate{
		Message: &discordgo.Message{
			Content:   "!civstrat mongols",
			ChannelID: "test-channel",
		},
	}

	expectedCivStr := strings.Join(civStrengths["Mongols"], ", ")
	expectedCivType := CivType["Mongols"]
	expectedEmojis := CivTypeToEmoji[expectedCivType]

	ListAllStrengths(mockSession, mockMessage)

	if mockSession.CallCount != 1 {
		t.Errorf("Expected ChannelMessageSend to be called once, got %d", mockSession.CallCount)
	}
	if !strings.Contains(mockSession.LastMessage, "Mongols") {
		t.Errorf("Expected message to contain 'Mongols', got: %s", mockSession.LastMessage)
	}
	if !strings.Contains(mockSession.LastMessage, expectedCivStr) {
		t.Errorf("Expected message to contain %v, got: %s", expectedCivStr, mockSession.LastMessage)
	}
	if !strings.Contains(mockSession.LastMessage, expectedEmojis) {
		t.Errorf("Expected message to contain %v, got: %s", expectedEmojis, mockSession.LastMessage)
	}
}

func TestCivInfo(t *testing.T) {
	mockSession := &MockSession{}
	mockMessage := &discordgo.MessageCreate{
		Message: &discordgo.Message{
			Content:   "!info Turks",
			ChannelID: "test-channel",
		},
	}

	expectedCivStr := strings.Join(civStrengths["Turks"], ", ")
	expectedCivWeak := strings.Join(civWeaknesses["Turks"], ", ")
	expectedCivType := CivType["Turks"]
	expectedEmojis := CivTypeToEmoji[expectedCivType]

	GetCivInfo(mockSession, mockMessage)

	if mockSession.CallCount != 1 {
		t.Errorf("Expected ChannelMessageSend to be called once, got %d", mockSession.CallCount)
	}
	if !strings.Contains(mockSession.LastMessage, "Turks") {
		t.Errorf("Expected message to contain 'Turks', got: %s", mockSession.LastMessage)
	}
	if !strings.Contains(mockSession.LastMessage, expectedCivStr) {
		t.Errorf("Expected message to contain %v, got: %s", expectedCivStr, mockSession.LastMessage)
	}
	if !strings.Contains(mockSession.LastMessage, expectedCivWeak) {
		t.Errorf("Expected message to contain %v, got: %s", expectedCivWeak, mockSession.LastMessage)
	}
	if !strings.Contains(mockSession.LastMessage, expectedEmojis) {
		t.Errorf("Expected message to contain %v, got: %s", expectedEmojis, mockSession.LastMessage)
	}
	if !strings.Contains(mockSession.LastMessage, expectedCivType) {
		t.Errorf("Expected message to contain %v, got: %s", expectedCivType, mockSession.LastMessage)
	}
}

func TestCivInfo_Lowercase(t *testing.T) {
	mockSession := &MockSession{}
	mockMessage := &discordgo.MessageCreate{
		Message: &discordgo.Message{
			Content:   "!info goths",
			ChannelID: "test-channel",
		},
	}
	expectedCivStr := strings.Join(civStrengths["Goths"], ", ")
	expectedCivWeak := strings.Join(civWeaknesses["Goths"], ", ")
	expectedCivType := CivType["Goths"]
	expectedEmojis := CivTypeToEmoji[expectedCivType]

	GetCivInfo(mockSession, mockMessage)

	if mockSession.CallCount != 1 {
		t.Errorf("Expected ChannelMessageSend to be called once, got %d", mockSession.CallCount)
	}
	if !strings.Contains(mockSession.LastMessage, "Goths") {
		t.Errorf("Expected message to contain 'Goths', got: %s", mockSession.LastMessage)
	}

	if !strings.Contains(mockSession.LastMessage, expectedCivStr) {
		t.Errorf("Expected message to contain %v, got: %s", expectedCivStr, mockSession.LastMessage)
	}
	if !strings.Contains(mockSession.LastMessage, expectedCivWeak) {
		t.Errorf("Expected message to contain %v, got: %s", expectedCivWeak, mockSession.LastMessage)
	}
	if !strings.Contains(mockSession.LastMessage, expectedEmojis) {
		t.Errorf("Expected message to contain %v, got: %s", expectedEmojis, mockSession.LastMessage)
	}
	if !strings.Contains(mockSession.LastMessage, expectedCivType) {
		t.Errorf("Expected message to contain %v, got: %s", expectedCivType, mockSession.LastMessage)
	}
}
