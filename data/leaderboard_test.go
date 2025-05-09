package data

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/bwmarrin/discordgo"
)

func TestGetLeaderboardData(t *testing.T) {
	mockJSON := `{
		"Britons": "56.2",
		"Franks": "52.8",
		"Aztecs": "60.1"
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
	if leaderboard["Britons"] != "56.2" {
		t.Errorf("Expected Britons winrate to be 56.2, got %s", leaderboard["Britons"])
	}
}

func TestFormatLeaderboardData(t *testing.T) {
	leaderboardData := map[string]string{
		"Britons": "56.2",
		"Franks":  "52.8",
		"Aztecs":  "60.1",
	}

	formattedData, err := FormatLeaderboardData(leaderboardData)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	expectedOutput := "Aztecs: 60.1\nBritons: 56.2\nFranks: 52.8\n"
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
		"Byzantines": "58.7",
		"Incas": "52.7",
		"Magyars": "60.3",
		"Malay": "56.2",
		"Mayans": "52.1",
		"Portuguese": "60.2",
		"Spanish": "56.8",
		"Turks": "52.8",
		"Vietnamese": "60.1"
	}`

	expectedOutput := "Magyars: 60.3\nPortuguese: 60.2\nVietnamese: 60.1\nByzantines: 58.7\nSpanish: 56.8\nMalay: 56.2\nTurks: 52.8\nIncas: 52.7\nMayans: 52.1\n"

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
