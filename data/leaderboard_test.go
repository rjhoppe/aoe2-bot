package data

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/bwmarrin/discordgo"
)

func TestGetLeaderboardData(t *testing.T) {
	mockJSON := `{
		"Britons": "56.20%",
		"Franks": "52.80%",
		"Aztecs": "60.10%"
	}`

	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "test_leaderboard.json")

	err := os.WriteFile(testFile, []byte(mockJSON), 0644)
	if err != nil {
		t.Fatalf("Failed to write test leaderboard.json: %v", err)
	}

	leaderboard, err := GetLeaderboardData(testFile)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	// Assert
	if len(leaderboard) != 3 {
		t.Errorf("Expected 3 entries, got %d", len(leaderboard))
	}
	if leaderboard["Britons"] != "56.20%" {
		t.Errorf("Expected Britons winrate to be 56.20%%, got %s", leaderboard["Britons"])
	}
}

func TestFormatLeaderboardData(t *testing.T) {
	leaderboardData := map[string]string{
		"Britons": "56.20%",
		"Franks":  "52.80%",
		"Aztecs":  "60.10%",
	}

	formattedData, err := FormatLeaderboardData(leaderboardData)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	expectedOutput := "Aztecs: 60.10%\nBritons: 56.20%\nFranks: 52.80%\n"
	if formattedData != expectedOutput {
		t.Errorf("Expected formatted data to be %q, got %q", expectedOutput, formattedData)
	}
}

func TestGetCivLeaderBoardAll(t *testing.T) {
	mockSession := &MockSession{}
	mockMessage := &discordgo.MessageCreate{
		Message: &discordgo.Message{
			Content:   "!leaderboard",
			ChannelID: "test-channel",
		},
	}

	mockJSON := `{
		"Byzantines": "58.70%",
		"Incas": "52.70%",
		"Magyars": "60.30%",
		"Malay": "56.20%",
		"Mayans": "52.10%",
		"Portuguese": "60.20%",
		"Spanish": "56.80%",
		"Turks": "52.80%",
		"Vietnamese": "60.10%"
	}`

	expectedOutput := "Magyars: 60.30%\nPortuguese: 60.20%\nVietnamese: 60.10%\nByzantines: 58.70%\nSpanish: 56.80%\nMalay: 56.20%\nTurks: 52.80%\nIncas: 52.70%\nMayans: 52.10%\n"

	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "test_leaderboard.json")

	// write mockJSON to test_leaderboard.json
	err := os.WriteFile(testFile, []byte(mockJSON), 0644)
	if err != nil {
		t.Fatalf("Failed to write test leaderboard.json: %v", err)
	}

	GetCivLeaderBoardAll(mockSession, mockMessage, testFile)
	if mockSession.LastChannelID == "" || mockSession.LastMessage == "" {
		t.Errorf("Expected ChannelMessageSend to be called, but it wasn't")
	}
	if mockSession.CallCount != 1 {
		t.Errorf("Expected ChannelMessageSend to be called once, but it was called %d times", mockSession.CallCount)
	}
	if mockSession.LastMessage != expectedOutput {
		t.Errorf("Expected Discord message to be %q, got %q", expectedOutput, mockSession.LastMessage)
	}
}

func TestGetCivWinRate(t *testing.T) {
	mockSession := &MockSession{}
	mockMessage := &discordgo.MessageCreate{
		Message: &discordgo.Message{
			Content:   "!winrate Vikings",
			ChannelID: "test-channel",
		},
	}

	mockJSON := `{
		"Britons": "56.20%",
		"Franks": "52.80%",
		"Aztecs": "60.10%",
		"Vikings": "51.23%"
	}`

	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "test_leaderboard.json")

	err := os.WriteFile(testFile, []byte(mockJSON), 0644)
	if err != nil {
		t.Fatalf("Failed to write test leaderboard.json: %v", err)
	}

	expectedCivWinRate := "51.23%"
	expectedCivType := CivType["Vikings"]
	expectedEmojis := CivTypeToEmoji[expectedCivType]

	GetCivWinRate(mockSession, mockMessage, testFile)

	if mockSession.CallCount != 1 {
		t.Errorf("Expected ChannelMessageSend to be called once, got %d", mockSession.CallCount)
	}
	if !strings.Contains(mockSession.LastMessage, "Vikings") {
		t.Errorf("Expected message to contain 'Vikings', got: %s", mockSession.LastMessage)
	}
	if !strings.Contains(mockSession.LastMessage, expectedCivWinRate) {
		t.Errorf("Expected message to contain %v, got: %s", expectedCivWinRate, mockSession.LastMessage)
	}
	if !strings.Contains(mockSession.LastMessage, expectedEmojis) {
		t.Errorf("Expected message to contain %v, got: %s", expectedEmojis, mockSession.LastMessage)
	}
}

func TestGetCivWinRate_Lowercase(t *testing.T) {
	mockSession := &MockSession{}
	mockMessage := &discordgo.MessageCreate{
		Message: &discordgo.Message{
			Content:   "!winrate poles",
			ChannelID: "test-channel",
		},
	}

	mockJSON := `{
		"Britons": "56.20%",
		"Franks": "52.80%",
		"Poles": "48.91%",
		"Aztecs": "60.10%",
		"Vikings": "51.23%"
	}`

	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "test_leaderboard.json")

	err := os.WriteFile(testFile, []byte(mockJSON), 0644)
	if err != nil {
		t.Fatalf("Failed to write test leaderboard.json: %v", err)
	}

	expectedCivWinRate := "48.91%"
	expectedCivType := CivType["Poles"]
	expectedEmojis := CivTypeToEmoji[expectedCivType]

	GetCivWinRate(mockSession, mockMessage, testFile)

	if mockSession.CallCount != 1 {
		t.Errorf("Expected ChannelMessageSend to be called once, got %d", mockSession.CallCount)
	}
	if !strings.Contains(mockSession.LastMessage, "Poles") {
		t.Errorf("Expected message to contain 'Poles', got: %s", mockSession.LastMessage)
	}
	if !strings.Contains(mockSession.LastMessage, expectedCivWinRate) {
		t.Errorf("Expected message to contain %v, got: %s", expectedCivWinRate, mockSession.LastMessage)
	}
	if !strings.Contains(mockSession.LastMessage, expectedEmojis) {
		t.Errorf("Expected message to contain %v, got: %s", expectedEmojis, mockSession.LastMessage)
	}
}
