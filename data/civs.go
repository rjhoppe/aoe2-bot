package data

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/rjhoppe/aoe-bot/utils"
)

var allCivs = []string{
	"Armenians",
	"Aztecs",
	"Bengalis",
	"Berbers",
	"Bohemians",
	"Britons",
	"Bulgarians",
	"Burgundians",
	"Burmese",
	"Byzantines",
	"Celts",
	"Chinese",
	"Cumans",
	"Dravidians",
	"Ethiopians",
	"Franks",
	"Georgians",
	"Goths",
	"Gurjaras",
	"Hindustanis",
	"Huns",
	"Incas",
	"Italians",
	"Japanese",
	"Khmer",
	"Koreans",
	"Lithuanians",
	"Magyars",
	"Malay",
	"Malians",
	"Mayans",
	"Mongols",
	"Persians",
	"Poles",
	"Portuguese",
	// "Romans" -> Don't have this DLC
	"Saracens",
	"Sicilians",
	"Slavs",
	"Spanish",
	"Tartars",
	"Teutons",
	"Turks",
	"Vietnamese",
	"Vikings",
}

var archerCivs = []string{
	"Bohemians",
	"Britons",
	"Chinese",
	"Ethiopians",
	"Italians",
	"Mayans",
	"Mongols",
	"Portuguese",
	"Spanish",
	"Tartars",
	"Turks",
	"Vietnamese",
}

var infantryCivs = []string{
	"Armenians",
	"Aztecs",
	"Bulgarians",
	"Celts",
	"Dravidians",
	"Goths",
	"Incas",
	"Japanese",
	"Malay",
	"Malians",
	"Sicilians",
	"Slavs",
	"Teutons",
	"Vikings",
}

var cavCivs = []string{
	"Bengalis",
	"Berbers",
	"Bulgarians",
	"Burgundians",
	"Burmese",
	"Cumans",
	"Franks",
	"Georgians",
	"Gurjaras",
	"Hindustanis",
	"Huns",
	"Khmer",
	"Lithuanians",
	"Magyars",
	"Mongols",
	"Persians",
	"Poles",
	"Saracens",
	"Sicilians",
	"Tartars",
}

var CivType = map[string]string{
	"Armenians":   "Infantry and Naval",
	"Aztecs":      "Infantry and Monk",
	"Bengalis":    "Elephant and Naval",
	"Berbers":     "Cavalry and Naval",
	"Bohemians":   "Gunpowder and Monk",
	"Britons":     "Foot archer",
	"Bulgarians":  "Infantry and Cavalry",
	"Burgundians": "Cavalry",
	"Burmese":     "Monk and Elephant",
	"Byzantines":  "Defensive",
	"Celts":       "Infantry and Siege",
	"Chinese":     "Archer",
	"Cumans":      "Cavalry",
	"Dravidians":  "Infantry and Naval",
	"Ethiopians":  "Foot archer",
	"Franks":      "Cavalry",
	"Georgians":   "Cavalry and Defensive",
	"Goths":       "Infantry",
	"Gurjaras":    "Cavalry and Camel",
	"Hindustanis": "Camel and Gunpowder",
	"Huns":        "Cavalry",
	"Incas":       "Infantry",
	"Italians":    "Foot archers and Naval",
	"Japanese":    "Infantry",
	"Khmer":       "Siege and Elephant",
	"Koreans":     "Defensive and Naval",
	"Lithuanians": "Cavalry and Monk",
	"Magyars":     "Cavalry",
	"Malay":       "Infantry and Naval",
	"Malians":     "Infantry",
	"Mayans":      "Foot archer",
	"Mongols":     "Mounted archer",
	"Persians":    "Cavalry",
	"Poles":       "Cavalry",
	"Portuguese":  "Naval and Gunpowder",
	"Saracens":    "Camel and Naval",
	"Sicilians":   "Infantry and Cavalry",
	"Slavs":       "Infantry and Siege",
	"Spanish":     "Gunpowder and Monk",
	"Tartars":     "Mounted archer",
	"Teutons":     "Infantry",
	"Turks":       "Gunpowder",
	"Vietnamese":  "Archer",
	"Vikings":     "Infantry and Naval",
}

var civStrengths = map[string][]string{
	"Armenians":   {"Drush", "Flush men-at-arms", "Fast Castle"},
	"Aztecs":      {"Drush", "Eagle Warrior rush"},
	"Bengalis":    {"Booming", "Elephant spam"},
	"Berbers":     {"Knight rush", "Booming", "Trash war"},
	"Bohemians":   {"Flush archer", "Tower rush", "Fast Castle"},
	"Britons":     {"Flush archer", "Crossbowmen rush"},
	"Bulgarians":  {"Drush", "Long Swordsmen rush", "Knight rush", "Krepost drop"},
	"Burgundians": {"Booming", "Knight rush", "Cavalier rush in Castle Age"},
	"Burmese":     {"Flush men-at-arms", "Fast Castle", "Cavalry spam"},
	"Byzantines":  {"Booming", "Turtle"},
	"Celts":       {"Drush", "Flush men-at-arms"},
	"Chinese":     {"Crossbowmen rush", "Knight rush"},
	"Cumans":      {"Scout rush", "Knight rush", "Steppe Lancer / Kipchak / Cav archer rush", "Booming"},
	"Dravidians":  {"Flush men-at-arms", "Flush archer", "Fast Castle", "Infantry spam"},
	"Ethiopians":  {"Flush archer", "Flush men-at-arms", "Crossbowmen rush"},
	"Franks":      {"Knight rush", "Scout rush", "Castle drop", "Fast Castle"},
	"Georgians":   {"Scout rush", "Knight rush", "Booming", "Turtle"},
	"Goths":       {"Drush", "Flush men-at-arms", "Fast Castle", "Infantry spam"},
	"Gurjaras":    {"Scout rush", "Camel Scout rush", "Booming", "Cavalry spam"},
	"Hindustanis": {"Booming", "Drush", "Scout rush"},
	"Huns":        {"Fast Castle", "Cavalry Archer rush", "Knight rush"},
	"Incas":       {"Tower rush", "Eagle Warrior rush", "Fast Castle"},
	"Italians":    {"Booming", "Fast Castle", "Galley rush"},
	"Japanese":    {"Flush men-at-arms", "Fast Castle"},
	"Khmer":       {"Fast Castle", "Flush archer", "Scout rush", "Scorpion and Elephant spam"},
	"Koreans":     {"Flush archer", "Turtle", "Booming"},
	"Lithuanians": {"Scout rush", "Fast Castle", "Knight Rush"},
	"Magyars":     {"Scout rush", "Trash war"},
	"Malay":       {"Elephant rush", "Naval spam", "Castle drop", "Trash war"},
	"Malians":     {"Flush men-at-arms", "Knight rush", "Booming"},
	"Mayans":      {"Flush archer", "Crossbowmen rush", "Archer and Eagle Warrior spam"},
	"Mongols":     {"Scout rush", "Mangudai and Siege spam", "Fast Castle"},
	"Persians":    {"Knight rush", "Persian Douche", "Booming"},
	"Poles":       {"Scout rush", "Flush men-at-arms", "Fast Castle", "Knight rush", "Trash wars"},
	"Portuguese":  {"Fast Castle", "Booming", "Gunpowder spam"},
	"Saracens":    {"Flush archer", "Booming", "Mameluke and Camel spam"},
	"Sicilians":   {"Flush men-at-arms", "Donjon drop", "Castle drop", "Booming", "Turtle"},
	"Slavs":       {"Knight rush", "Flush men-at-arms"},
	"Spanish":     {"Fast Castle", "Booming"},
	"Tartars":     {"Scout rush", "Knight rush", "Steppe Lancer / Cav Archer rush", "Cavalry archer spam"},
	"Teutons":     {"Booming", "Tower rush", "Knight rush"},
	"Turks":       {"Fast Imperial", "Fast Castle", "Booming"},
	"Vietnamese":  {"Fast Imperial", "Cavalry Archer rush"},
	"Vikings":     {"Fast Castle", "Knight rush", "Berserk / Longboat spam"},
}

var civWeaknesses = map[string][]string{
	"Armenians": {
		"Slow economy",
		"Missing many key Imperial Age upgrades",
		"Particularly weak to powerful late-game archers and siege",
	},
	"Aztecs": {
		"Weak late-game",
		"Gold reliant army",
		"Terrible late-game navy",
		"Weak defenses",
		"Mediocre archers",
	},
	"Bengalis": {
		"Missing many key Imperial Age upgrades",
		"Mediocre siege",
		"Mediocre cavalry",
		"Vulernable to early game rushes",
	},
	"Berbers": {
		"Weak defenses",
		"Weak siege",
		"Wood reliant army",
	},
	"Bohemians": {
		"Weak cavalry",
		"Mediocre archers",
		"Slow economy",
		"Mediocre navy",
	},
	"Britons": {
		"Weak cavalry",
		"Easily countered with Siege Rams",
		"Reliant on Trebuchets for siege",
		"Mediocre archer options outside of Longbowmen",
	},
	"Bulgarians": {
		"Weak archers",
		"Mediocre defenses",
		"Weak navy",
		"Mediocre economy",
	},
	"Burgundians": {
		"Weak archers",
		"Weak siege",
		"Mixed-bag unique unit",
		"Mediocre navy",
		"Needs to boom effectively to gain an advantage",
	},
	"Burmese": {
		"Terrible archers",
		"Need cavalry and siege to counter archers",
		"Mediocre late-game navy",
	},
	"Byzantines": {
		"Mediocre cavalry",
		"Mediocre siege",
		"Struggles to play offensively",
		"Unique unit weak vs cavalry and archers",
	},
	"Celts": {
		"Weak archers",
		"Weak defenses",
		"Mediocre cavalry",
		"Mediocre trash war options",
	},
	"Chinese": {
		"Difficult to play properly",
		"Easily disrupted",
		"Limited anti-siege options",
	},
	"Cumans": {
		"Weak defenses",
		"Weak navy",
		"Mediocre late-game",
		"Cavalry archers are their only viable archers",
	},
	"Dravidians": {
		"Weakest cavalry in the game",
		"Poor troop mobility",
	},
	"Ethiopians": {
		"Missing many key Imperial Age upgrades",
		"Weak late-game navy",
	},
	"Franks": {
		"Weak archers",
		"Terrible trash war options",
		"Strength fizzles out in the late gate",
	},
	"Georgians": {
		"Mediocre rushing options",
		"Weak navy",
		"Mediocre archers",
	},
	"Goths": {
		"Worst defenses in the game",
		"Weak to Siege Onagers",
		"Mediocre siege",
		"Mediocre archers",
		"Must constantly be on the offensive",
	},
	"Gurjaras": {
		"Terrible infantry",
		"Terrible trash war options",
		"Heavily reliant on camels",
	},
	"Hindustanis": {
		"Vulnerable to monk counter",
		"Weak knight line",
		"Weak defenses",
	},
	"Huns": {
		"Weak infantry",
		"Mediocre siege",
		"Terrible defenses",
		"Predictable end game spam (Heavy Cav Archers, Siege Rams, Trebs)",
	},
	"Incas": {
		"Heavily reliant on massing unique unit to counter cavalry",
		"Particularly weak against Aztecs and Mayans",
		"No significant military bonuses",
	},
	"Italians": {
		"Struggles to counter cavalry without unique unit",
		"Slow economy",
		"Weak siege",
		"Lack of bonuses until Imperial",
	},
	"Japanese": {
		"Weak cavalry",
		"Weak siege",
		"Weak late-game economy",
	},
	"Khmer": {
		"Weak infantry",
		"Civ bonus can be a double-edged sword",
		"Mediocre archers",
	},
	"Koreans": {
		"Slow economy",
		"Weak cavalry",
	},
	"Lithuanians": {
		"Mediocre siege",
		"Mediocre infantry",
	},
	"Magyars": {
		"Terrible infantry",
		"Weak defenses",
		"Mediocre siege",
	},
	"Malay": {
		"Generally, pretty terrible at land combat",
		"Worst cavalry in the game",
		"Weak infantry",
		"Limited to only Arbalesters for anti-infantry",
	},
	"Malians": {
		"Missing many key late-game blacksmith upgrades",
		"Weak late-game infantry and navy",
	},
	"Mayans": {
		"Unable to counter late-game heavy cavalry or gunpowder",
		"Terrible swordsmen line",
	},
	"Mongols": {
		"Weak mid game",
		"Limited tech tree outside of Mangudai, siege, and Hussars, Limited trash war options",
	},
	"Persians": {
		"Weak late-game",
		"Highly reliant on heavy cavalry",
		"Elephants are easily countered with monks",
	},
	"Poles": {
		"Weak to elephants",
		"Glass cannon cavalry weak to tougher/heavier cavalry",
		"Weak anti-cav",
		"Weak navy",
	},
	"Portuguese": {
		"Expensive navy",
		"Weak late-game cavalry and siege",
		"Slow units",
	},
	"Saracens": {
		"Weak late-game economy",
		"Lack of affordable, late-game anti-cav options",
		"Highly reliant on gold",
		"Terrible trash war options",
	},
	"Sicilians": {
		"Weak to early rushing",
		"Slow early-game economy",
		"Bad archers",
		"Weak late-game due to mediocre tech tree",
	},
	"Slavs": {
		"Weak to early rushing",
		"Poor defensive options",
		"Weak archers",
	},
	"Spanish": {
		"Weak early-game",
		"Terrible foot archer options",
		"Average siege",
		"Weak to Flush or Fast Castle",
	},
	"Tartars": {
		"Worst infantry in game",
		"Unable to counter cavalry or Eagle Warriors",
	},
	"Teutons": {
		"Slow early-game",
		"Mediocre cavalry and archer options",
		"Terrible trash war options",
	},
	"Turks": {
		"Weak early-game",
		"Susceptible to rushes",
		"Terrible trash war options",
		"Highly reliant on gold",
	},
	"Vietnamese": {
		"Terrible siege",
		"Weak infantry and cavalry",
	},
	"Vikings": {
		"Limited cavalry options",
		"Weak to Galley rush",
		"Weak to Onagers",
	},
}

func GetNewRandomCiv(civType string) *Civilization {
	var civ Civilization

	switch civType {
	case "archer":
		civ.Name = utils.SelectRandomArrayEle(archerCivs)
	case "cavalry":
		civ.Name = utils.SelectRandomArrayEle(cavCivs)
	case "infantry":
		civ.Name = utils.SelectRandomArrayEle(infantryCivs)
	default:
		civ.Name = utils.SelectRandomArrayEle(allCivs)
	}

	civ.Type = CivType[civ.Name]
	civ.Strengths = strings.Join(civStrengths[civ.Name], ", ")
	civ.Weaknesses = strings.Join(civWeaknesses[civ.Name], ", ")

	return &civ
}

func PrintCivOutput(civType string, civ *Civilization, s DiscordSession, m *discordgo.MessageCreate) {
	msg := fmt.Sprintf(
		`%v: %v || %v ||
---------------------------
Civ Strategies: || %v ||
Civ Weakness: || %v ||`, civ.Name, CivTypeToEmoji[civ.Type], civ.Type, civ.Strengths, civ.Weaknesses)
	_, err := s.ChannelMessageSend(m.ChannelID, msg)
	if err != nil {
		fmt.Printf("Error sending message to %v \n", msg)
	}
}

func GetThreeRandomCivs(s DiscordSession, m *discordgo.MessageCreate) {
	civsMap := make(map[string]bool)
	civNames := make([]string, 0, 3)

	for len(civsMap) < 3 {
		civ := GetNewRandomCiv("all")
		if !civsMap[civ.Name] {
			civsMap[civ.Name] = true
			civNames = append(civNames, civ.Name)
		}
	}

	msg := strings.Join(civNames, ", ")
	_, err := s.ChannelMessageSend(m.ChannelID, msg)
	if err != nil {
		fmt.Printf("Error sending message to %v \n", msg)
	}
}

// !info <civ>
func GetCivInfo(s DiscordSession, m *discordgo.MessageCreate) {
	if len(m.Content) < 6 {
		errMsg := "Invalid civ name"
		_, err := s.ChannelMessageSend(m.ChannelID, errMsg)
		if err != nil {
			fmt.Printf("Error sending message to %v \n", errMsg)
		}
	}

	civRaw := utils.FirstCharToUpper(m.Content[6:])
	civType := CivType[civRaw]
	if civType != "" {
		civStr := strings.Join(civStrengths[civRaw], ", ")
		civWeak := strings.Join(civWeaknesses[civRaw], ", ")
		msg := fmt.Sprintf(`
%v: %v || %v ||
---------------------------
Strengths: %v
Weaknesses: %v`, civRaw, CivTypeToEmoji[civType], civType, civStr, civWeak)
		_, err := s.ChannelMessageSend(m.ChannelID, msg)
		if err != nil {
			fmt.Printf("Error sending message to %v \n", msg)
		}
	} else {
		msg := fmt.Sprintf("Could not find civ: %v", civRaw)
		_, err := s.ChannelMessageSend(m.ChannelID, msg)
		if err != nil {
			fmt.Printf("Error sending message to %v \n", msg)
		}
	}
}

// !civstrat
func ListAllStrengths(s DiscordSession, m *discordgo.MessageCreate) {
	isCmdValid := utils.IsValidCmd(10, s, m)
	if !isCmdValid {
		return
	}

	civRaw := utils.FirstCharToUpper(m.Content[10:])
	civStrs := strings.Join(civStrengths[civRaw], ", ")
	if civStrs != "" {
		civType := CivType[civRaw]
		civEmojis := CivTypeToEmoji[civType]

		msg := fmt.Sprintf(`
	%v: %v
---------------------------
Strategies: %v`, civRaw, civEmojis, civStrs)
		_, err := s.ChannelMessageSend(m.ChannelID, msg)
		if err != nil {
			fmt.Printf("Error sending message to %v \n", msg)
		}
	}
}
