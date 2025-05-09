package data

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/rjhoppe/aoe-bot/utils"
)

// leaderboard
func GetCivLeaderBoardAll(s DiscordSession, m *discordgo.MessageCreate, fp string) error {
	leaderboard, err := GetLeaderboardData(fp)
	if err != nil {
		return fmt.Errorf("error getting leaderboard data: %w", err)
	}

	msg, err := FormatLeaderboardData(leaderboard)
	if err != nil {
		return fmt.Errorf("error formatting leaderboard data: %w", err)
	}

	_, err = s.ChannelMessageSend(m.ChannelID, msg)
	if err != nil {
		fmt.Printf("Error sending message to %v \n", msg)
	}

	return nil
}

// helper function to get the leaderboard data
func GetLeaderboardData(fp string) (map[string]string, error) {
	content, err := os.ReadFile(fp)
	if err != nil {
		return nil, fmt.Errorf("error reading leaderboard.json: %w", err)
	}

	var parsedContent map[string]string
	err = json.Unmarshal(content, &parsedContent)
	if err != nil {
		return nil, fmt.Errorf("error parsing leaderboard.json: %w", err)
	}

	return parsedContent, nil
}

// helper function to format the leaderboard data
func FormatLeaderboardData(data map[string]string) (string, error) {
	type pair struct {
		Civ     string
		Winrate float64
	}

	var pairs []pair
	for civ, winRateStr := range data {
		winrate, err := strconv.ParseFloat(winRateStr, 64)
		if err != nil {
			return "", fmt.Errorf("error parsing winrate: %w", err)
		}
		pairs = append(pairs, pair{Civ: civ, Winrate: winrate})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Winrate > pairs[j].Winrate
	})

	var msg string
	for _, pair := range pairs {
		msg += fmt.Sprintf("%v: %v\n", pair.Civ, pair.Winrate)
	}

	return msg, nil
}

// !winrate
func GetCivWinRate(s DiscordSession, m *discordgo.MessageCreate) {
	isCmdValid := utils.IsValidCmd(10, s, m)
	if !isCmdValid {
		return
	}

	civRaw := m.Content[9:]

	content, err := os.ReadFile("data/leaderboard.json")
	if err != nil {
		fmt.Printf("error writing file: %v", err)
		return
	}

	var data map[string]interface{}
	err = json.Unmarshal(content, &data)
	if err != nil {
		fmt.Printf("Error parsing JSON: %v", err)
		return
	}

	if civWinRate, ok := data[civRaw].(string); ok {
		civType := CivType[civRaw]
		civEmojis := CivTypeToEmoji[civType]

		msg := fmt.Sprintf(`
		%v: %v
---------------------------
Win Rate: %v`, civRaw, civEmojis, civWinRate)
		_, err := s.ChannelMessageSend(m.ChannelID, msg)
		if err != nil {
			fmt.Printf("Error sending message to %v \n", msg)
		}
	} else {
		errMsg := fmt.Sprintf("Could not find data for '%v'", civRaw)
		_, err := s.ChannelMessageSend(m.ChannelID, errMsg)
		if err != nil {
			fmt.Printf("Error sending message to %v \n", errMsg)
		}
	}
}
