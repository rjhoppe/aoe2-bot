package data

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/rjhoppe/aoe-bot/utils"
)

type Player struct {
	Name string
	Civ  *Civilization
}

func BuildGame(gameType string) []Player {
	var players []Player

	switch gameType {
	case "1v1":
		players = []Player{
			{Name: "Player 1", Civ: GetNewRandomCiv("all")},
			{Name: "CPU 1", Civ: GetNewRandomCiv("all")},
		}
	case "2v2":
		players = []Player{
			{Name: "Player 1", Civ: GetNewRandomCiv("all")},
			{Name: "Player 2", Civ: GetNewRandomCiv("all")},
			{Name: "CPU 1", Civ: GetNewRandomCiv("all")},
			{Name: "CPU 2", Civ: GetNewRandomCiv("all")},
		}
	case "4v4":
		players = []Player{
			{Name: "Player 1", Civ: GetNewRandomCiv("all")},
			{Name: "Player 2", Civ: GetNewRandomCiv("all")},
			{Name: "Player 3", Civ: GetNewRandomCiv("all")},
			{Name: "Player 4", Civ: GetNewRandomCiv("all")},
			{Name: "CPU 1", Civ: GetNewRandomCiv("all")},
			{Name: "CPU 2", Civ: GetNewRandomCiv("all")},
			{Name: "CPU 3", Civ: GetNewRandomCiv("all")},
			{Name: "CPU 4", Civ: GetNewRandomCiv("all")},
		}
	default:
		players = []Player{
			{Name: "Player 1", Civ: GetNewRandomCiv("all")},
			{Name: "Player 2", Civ: GetNewRandomCiv("all")},
			{Name: "Player 3", Civ: GetNewRandomCiv("all")},
			{Name: "CPU 1", Civ: GetNewRandomCiv("all")},
			{Name: "CPU 2", Civ: GetNewRandomCiv("all")},
			{Name: "CPU 3", Civ: GetNewRandomCiv("all")},
		}
	}

	return players
}

func PrintGame(players []Player, s DiscordSession, m *discordgo.MessageCreate) {
	date := utils.GetCurDate()
	randomMap := GetRandomMap("all")
	msg := fmt.Sprintf(
		`New Game %v
---------------------------
Map: %v`, date, randomMap)
	msg += fmt.Sprintln("")
	msg += fmt.Sprintln("")
	for _, player := range players {
		msg += fmt.Sprintf(
			`%v
---------------------------
%v: %v
Civ Strategies: || %v ||
Civ Weaknesses: || %v ||`, player.Name, player.Civ.Name, CivTypeToEmoji[player.Civ.Type], player.Civ.Strengths, player.Civ.Weaknesses)
		msg += fmt.Sprintln("")
		msg += "--------------------------- \n"
		msg += fmt.Sprintln("")
	}
	_, err := s.ChannelMessageSend(m.ChannelID, msg)
	if err != nil {
		fmt.Printf("Error sending message to %v \n", msg)
	}
}
