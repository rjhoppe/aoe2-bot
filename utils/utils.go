package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func SelectRandomArrayEle(arr []string) string {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	randomIndex := r.Intn(len(arr))
	return arr[randomIndex]
}

func GetCurDate() string {
	now := time.Now()
	dateFormat := now.Format("01/02/2006")
	return dateFormat
}

// Check text for the Discord bot cmd initiator char
func CheckMsgForExclamation(msg string) bool {
	cmdChar := '!'
	return strings.ContainsRune(msg, cmdChar)
}

func PrintCmds(s *discordgo.Session, m *discordgo.MessageCreate) {
	var commandList strings.Builder
	commandList.WriteString("**!civ** -> Return a random civ")
	commandList.WriteString("\n")
	commandList.WriteString("**!random** -> Returns 3 random civs")
	commandList.WriteString("\n")
	commandList.WriteString("**!archciv** -> Return a random archer civ")
	commandList.WriteString("\n")
	commandList.WriteString("**!cavciv** -> Return a random cavalary civ")
	commandList.WriteString("\n")
	commandList.WriteString("**!infciv** -> Return a random infantry civ")
	commandList.WriteString("\n")
	commandList.WriteString("**!info <CIV>** -> Returns info on a particular civ")
	commandList.WriteString("\n")
	commandList.WriteString("**!summon** -> Pings everyone in the channel")
	commandList.WriteString("\n")
	commandList.WriteString("**!map** -> Return a random map")
	commandList.WriteString("\n")
	commandList.WriteString("**!game1** -> Random the settings for a game with a 1v1 singleplayer game")
	commandList.WriteString("\n")
	commandList.WriteString("**!game2** -> Random the settings for a game with 2 human players")
	commandList.WriteString("\n")
	commandList.WriteString("**!game3** -> Random the settings for a game with 3 human players")
	cmds := commandList.String()
	_, err := s.ChannelMessageSend(m.ChannelID, cmds)
	if err != nil {
		fmt.Printf("Error sending message to %v \n", cmds)
	}
}
