package data

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/rjhoppe/aoe-bot/utils"
)

var MapTypes = []string{
	"Random",
	"Land",
	"Water",
	"Unique",
	"Real World",
}

var RandomMaps = []string{
	"Black Forest",
	"Arabia",
	"Arena",
	"Nomad",
	"Amazon Tunnel",
	"Team Islands",
	"Islands",
	"Snake Forest",
	"Gold Rush",
	"Steppe",
	"Seize the Mountain",
	"Team Moats",
	"Hideout",
	"Land Madness",
	"Oasis",
	"Lombardia",
	"Migration",
	"Pilgrims",
	"Canyons",
	"Sherwood Forest",
	"Water Nomads",
	"Nomads",
	"Britain",
	"Aral Sea",
	"Bohemia",
	"Madagascar",
	"Central America",
	"Sea of Japan",
	"Mideast",
	"France",
	"Norse Lands",
	"El Dorado",
	"Land Madness",
	"Team Acropolis",
}

var LandMaps = []string{
	"Black Forest",
	"Arena",
	"Nomad",
	"Amazon Tunnel",
	"Snake Forest",
	"Gold Rush",
	"Seize the Mountain",
	"Team Moats",
	"Hideout",
	"Land Madness",
	"Steppe",
	"Oasis",
	"Lombardia",
	"Canyons",
	"Sherwood Forest",
	"El Dorado",
	"Land Madness",
	"Team Acropolis",
}

var WaterMaps = []string{
	"Islands",
	"Team Islands",
	"Mediterranean",
	"Continental",
	"Migration",
	"Aral Sea",
	"Bohemia",
	"Madagascar",
	"Central America",
	"Sea of Japan",
	"Mideast",
	"France",
	"Norse Lands",
}

var UniqueMaps = []string{
	"Pilgrims",
	"Water Nomads",
	"Nomads",
}

var RealWorldMaps = []string{
	"Britain",
	"Aral Sea",
	"Bohemia",
	"Madagascar",
	"Central America",
	"Sea of Japan",
	"Mideast",
	"France",
	"Norse Lands",
}

func GetRandomMap(mapType string) string {
	switch mapType {
	case "water":
		randomMap := utils.SelectRandomArrayEle(WaterMaps)
		return randomMap
	case "land":
		randomMap := utils.SelectRandomArrayEle(LandMaps)
		return randomMap
	default:
		randomMap := utils.SelectRandomArrayEle(RandomMaps)
		return randomMap
	}
}

func PrintRandomMap(randomMap string, s DiscordSession, m *discordgo.MessageCreate) {
	msg := fmt.Sprintf("%v", randomMap)
	_, err := s.ChannelMessageSend(m.ChannelID, msg)
	if err != nil {
		fmt.Printf("Error sending message to %v \n", msg)
	}
}
