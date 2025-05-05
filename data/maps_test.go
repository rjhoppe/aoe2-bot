package data

import (
	"testing"
)


func TestGetRandomWaterMap(t *testing.T) {
	mapType := "water"
	randomMap := GetRandomMap(mapType)
	if randomMap == "" {
		t.Errorf("Expected a map, got empty string")
	}
}

func TestGetRandomLandMap(t *testing.T) {
	mapType := "land"
	randomMap := GetRandomMap(mapType)
	if randomMap == "" {
		t.Errorf("Expected a map, got empty string")
	}
}

func TestPrintRandomMap(t *testing.T) {
	mockSession := &MockSession{}
	mockMessage := &discordgo.MessageCreate{
		Message: &discordgo.Message{
			Content: "!map",
		},
	}
	PrintRandomMap(randomMap, mockSession, mockMessage)
	if mockSession.LastChannelID == "" || mockSession.LastMessage == "" {
		t.Errorf("Expected ChannelMessageSend to be called, but it wasn't")
	}
	if mockSession.CallCount != 1 {
		t.Errorf("Expected ChannelMessageSend to be called once, but it was called %d times", mockSession.CallCount)
	}
}

func TestGetRandomMap_All(t *testing.T) {
	mapType := "all"
	randomMap := GetRandomMap(mapType)
	if randomMap == "" {
		t.Errorf("Expected a map, got empty string")
	}
	found := false
	for _, mapName := range allMaps {
		if randomMap == mapName {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Expected a valid map, got %v", randomMap)
	}
}
