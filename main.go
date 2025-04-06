package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/rjhoppe/aoe-bot/bot"
	"github.com/rjhoppe/aoe-bot/data"
)

func init() {
	err := godotenv.Load("local.env")
	if err != nil {
		fmt.Printf("error loading local.env file: %e", err)
		fmt.Println("local.env file may not be present")
	}
}

func main() {
	config := &data.BotConfig{}

	textChanId := os.Getenv("TEXT_CHAN_ID")
	guildId := os.Getenv("GUILD_ID")
	token := os.Getenv("TOKEN")
	botPrefix := os.Getenv("BOT_PREFIX")

	config.Token = token
	config.BotPrefix = botPrefix

	botParams := &bot.BotParams{
		TextChanId: textChanId,
		GuildId:    guildId,
		Config:     config,
	}
	fmt.Println("Starting bot...")
	bot.Start(botParams)

	// blocks program from exiting
	<-make(chan struct{})
}
