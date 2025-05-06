package data

import (
	"testing"

	"github.com/bwmarrin/discordgo"
)

func TestPrintGame(t *testing.T) {
	mockPlayers := []Player{}
	mockSession := &MockSession{}
	mockMessage := &discordgo.MessageCreate{
		Message: &discordgo.Message{
			Content: "!game",
		},
	}
	PrintGame(mockPlayers, mockSession, mockMessage)
	if mockSession.LastChannelID == "" || mockSession.LastMessage == "" {
		t.Errorf("Expected ChannelMessageSend to be called, but it wasn't")
	}
	if mockSession.CallCount != 1 {
		t.Errorf("Expected ChannelMessageSend to be called once, but it was called %d times", mockSession.CallCount)
	}
}

func TestBuildGame_1v1(t *testing.T) {
	gameType := "1v1"
	players := BuildGame(gameType)
	if len(players) != 2 {
		t.Errorf("Expected 2 players, got %d", len(players))
	}

	for _, player := range players {
		if player.Civ.Name == "" {
			t.Errorf("Expected player to have a civ, got %v", player.Civ)
		}
	}
}

func TestBuildGame_2v2(t *testing.T) {
	gameType := "2v2"
	players := BuildGame(gameType)
	if len(players) != 4 {
		t.Errorf("Expected 4 players, got %d", len(players))
	}

	for _, player := range players {
		if player.Civ.Name == "" {
			t.Errorf("Expected player to have a civ, got %v", player.Civ)
		}
	}
}

func TestBuildGame_3v3(t *testing.T) {
	gameType := "3v3"
	players := BuildGame(gameType)
	if len(players) != 6 {
		t.Errorf("Expected 6 players, got %d", len(players))
	}

	for _, player := range players {
		if player.Civ.Name == "" {
			t.Errorf("Expected player to have a civ, got %v", player.Civ)
		}
	}
}

func TestBuildGame_4v4(t *testing.T) {
	gameType := "4v4"
	players := BuildGame(gameType)
	if len(players) != 8 {
		t.Errorf("Expected 8 players, got %d", len(players))
	}

	for _, player := range players {
		if player.Civ.Name == "" {
			t.Errorf("Expected player to have a civ, got %v", player.Civ)
		}
	}
}
