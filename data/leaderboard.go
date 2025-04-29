package data

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

// leaderboard
func GetCivLeaderBoardAll(s *discordgo.Session, m *discordgo.MessageCreate) error {
	content, err := os.ReadFile("leaderboard.json")
	if err != nil {
		return fmt.Errorf("error writing file: %w", err)
	}

	var parsedContent map[string]string
	err = json.Unmarshal(content, &parsedContent)
	if err != nil {
		return fmt.Errorf("error parsing leaderboard.json: %w", err)
	}

	var msg string
	for key, value := range parsedContent {
		msg += fmt.Sprintf("%v: %v", key, value)
		msg += fmt.Sprint("\n")
	}

	_, err = s.ChannelMessageSend(m.ChannelID, msg)
	if err != nil {
		fmt.Printf("Error sending message to %v \n", msg)
	}

	return nil
}
