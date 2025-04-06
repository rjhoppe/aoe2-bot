package data

type BotConfig struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
}

type Civilization struct {
	Name       string
	Type       string
	Strengths  string
	Weaknesses string
}

var CivTypeToEmoji = map[string]string{
	"Infantry and Monk":      "⚔️✝️",
	"Elephant and Naval":     "🐘🚢",
	"Cavalry and Naval":      "🐴🚢",
	"Gunpowder and Monk":     "🔫✝️",
	"Foot archer":            "🦶🏹",
	"Infantry and Cavalry":   "⚔️🐴",
	"Cavalry":                "🐴",
	"Monk and Elephant":      "✝️🐘",
	"Defensive":              "🏰",
	"Infantry and Siege":     "⚔️⚙️",
	"Cavalry and Defensive":  "🐴🏰",
	"Infantry":               "⚔️",
	"Cavalry and Camel":      "🐴🐫",
	"Camel and Gunpowder":    "🐫🔫",
	"Foot archers and Naval": "🦶🏹🚢",
	"Siege and Elephant":     "⚙️🐘",
	"Defensive and Naval":    "🏰🚢",
	"Cavalry and Monk":       "🐴✝️",
	"Mounted archer":         "🏇🏹",
	"Naval and Gunpowder":    "🚢🔫",
	"Camel and Naval":        "🐫🚢",
	"Gunpowder":              "🔫",
	"Archer":                 "🏹",
	"Infantry and Naval":     "⚔️🚢",
}
