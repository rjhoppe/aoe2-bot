package utils

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

type DiscordSession interface {
	ChannelMessageSend(channelID string, content string, options ...discordgo.RequestOption) (*discordgo.Message, error)
}

func CreateDataDir() {
	err := os.MkdirAll("data", 0755)
	if err != nil {
		// handle error, e.g. log or return
		fmt.Printf("Failed to create data directory: %v\n", err)
		return
	}
}

func SelectRandomArrayEle(arr []string) string {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	randomIndex := r.Intn(len(arr))
	return arr[randomIndex]
}

func GetAllKeys(dataMap map[string][]string) string {
	keyList := make([]string, 0, len(dataMap))
	for key := range dataMap {
		keyList = append(keyList, key)
	}

	keys := strings.Join(keyList, ", ")
	return keys
}

func GetCurDate() string {
	now := time.Now()
	dateFormat := now.Format("01/02/2006")
	return dateFormat
}

func FirstCharToUpper(s string) string {
	if s != "" {
		return strings.ToUpper(string(s[0])) + s[1:]
	}
	return s
}

// Check text for the Discord bot cmd initiator char
func CheckMsgForExclamation(msg string) bool {
	cmdChar := '!'
	return strings.ContainsRune(msg, cmdChar)
}

var IsValidCmd = func(validLen int, s DiscordSession, m *discordgo.MessageCreate) bool {
	if len(m.Content) < validLen {
		errMsg := "Invalid input, not enough chars"
		_, err := s.ChannelMessageSend(m.ChannelID, errMsg)
		if err != nil {
			fmt.Printf("Error sending message to %v \n", errMsg)
		}
		return false
	}
	return true
}

func PrintCmds(s DiscordSession, m *discordgo.MessageCreate) {
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
	commandList.WriteString("\n")
	commandList.WriteString("**!strat <STRATEGY>** -> Returns information on a given strategy")
	commandList.WriteString("\n")
	commandList.WriteString("**!stratlist** -> Lists all strategies you can pass to the !strat or !stratciv cmds")
	commandList.WriteString("\n")
	commandList.WriteString("**!stratcivs <STRATEGY>** -> Returns all the civs that can employ a specified strategy effectively")
	commandList.WriteString("\n")
	commandList.WriteString("**!civstrat <CIV>** -> Returns all the common strategies associated with particular civilization")
	commandList.WriteString("\n")
	commandList.WriteString("**!leaderboard** -> Returns the win rate for civs in competitive play")
	cmds := commandList.String()
	_, err := s.ChannelMessageSend(m.ChannelID, cmds)
	if err != nil {
		fmt.Printf("Error sending message to %v \n", cmds)
	}
}
