package bot

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/bwmarrin/discordgo"
	"github.com/rjhoppe/aoe-bot/data"
	"github.com/rjhoppe/aoe-bot/utils"
)

var botId string

type BotParams struct {
	TextChanId  string
	VoiceChanId string
	GuildId     string
	Config      *data.BotConfig
}

func Start(params *BotParams) {
	var botSession *discordgo.Session

	botSession, err := discordgo.New("Bot " + params.Config.Token)
	if err != nil {
		log.Fatalf("something went wrong when initializing the bot: %v", err)
	}

	botUser, err := botSession.User("@me")
	if err != nil {
		log.Fatalf("something went wrong when initializing bot user: %v", err)
	}

	botId = botUser.ID
	botSession.AddHandler(messageHandler)
	err = botSession.Open()
	if err != nil {
		log.Fatalf("something went wrong when starting the bot: %v", err)
	}

	select {}
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == botId {
		return
	}

	if utils.CheckMsgForExclamation(m.Content) {
		switch {
		case strings.Contains(strings.ToLower(m.Content), "!summon"):
			summonLads(s, m)
		case strings.Contains(m.Content, "!cmds"):
			utils.PrintCmds(s, m)
		case strings.Contains(m.Content, "!civ"):
			civ := data.GetNewRandomCiv("all")
			data.PrintCivOutput("", civ, s, m)
		case strings.Contains(m.Content, "!random"):
			data.GetThreeRandomCivs(s, m)
		case strings.Contains(m.Content, "!archciv"):
			civ := data.GetNewRandomCiv("archer")
			data.PrintCivOutput("Archer ", civ, s, m)
		case strings.Contains(m.Content, "!cavciv"):
			civ := data.GetNewRandomCiv("cavalry")
			data.PrintCivOutput("Cavalry ", civ, s, m)
		case strings.Contains(m.Content, "!infciv"):
			civ := data.GetNewRandomCiv("infantry")
			data.PrintCivOutput("Infantry ", civ, s, m)
		case strings.Contains(m.Content, "!info"):
			data.GetCivInfo(s, m)
		case strings.Contains(m.Content, "!map"):
			randomMap := data.GetRandomMap("all")
			data.PrintRandomMap(randomMap, s, m)
		case strings.Contains(m.Content, "!game1"):
			players := data.BuildGame("1v1")
			data.PrintGame(players, s, m)
		case strings.Contains(m.Content, "!game2"):
			players := data.BuildGame("2v2")
			data.PrintGame(players, s, m)
		case strings.Contains(m.Content, "!game3") || strings.Contains(m.Content, "!game"):
			players := data.BuildGame("3v3")
			data.PrintGame(players, s, m)
		case strings.Contains(m.Content, "!strat"):
			fmt.Println("Placeholders")
		case strings.Contains(m.Content, "!stratcivs"):
			fmt.Println("Placeholders")
		case strings.Contains(m.Content, "!civstrat"):
			fmt.Println("Placeholders")
		case strings.Contains(m.Content, "!help"):
			msg := "Try using the **!cmds** command to get a list of all commands this bot accepts"
			_, err := s.ChannelMessageSend(m.ChannelID, msg)
			if err != nil {
				fmt.Printf("Error sending message to %v \n", msg)
			}
		}
	}
}

func summonLads(s *discordgo.Session, m *discordgo.MessageCreate) {
	userIds := []string{
		"305206328039833601",
		"305198119618871299",
		"1294366970631819296",
	}

	var wg sync.WaitGroup
	for _, userId := range userIds {
		wg.Add(1)

		go func(userId string) {
			defer wg.Done()

			mention := "<@" + userId + ">"
			message := fmt.Sprintf("Summoning...%s", mention)

			_, err := s.ChannelMessageSend(m.ChannelID, message)
			if err != nil {
				fmt.Printf("Error sending message to %v \n", userId)
			}
		}(userId)
	}
	wg.Wait()
}
