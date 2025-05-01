package data

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/rjhoppe/aoe-bot/utils"
)

var StratToCivs = map[string][]string{
	"flush archer":      {"Bohemenians", "Britons", "Dravidians", "Ethiopians", "Khmer", "Koreans", "Mayans", "Saracens"},
	"flush men-at-arms": {"Celts", "Dravidians", "Ethiopians", "Goths", "Japanese", "Malians", "Poles", "Sicilians", "Slavs"},
	"drush":             {"Armenians", "Aztecs", "Bulgarians", "Celts", "Goths"},
	"scout rush":        {"Cumans", "Franks", "Georgians", "Gurjaras", "Hindustanis", "Khmer", "Lithuanians", "Magyars", "Mongols", "Poles", "Tartars"},
	"tower rush":        {"Bohemians", "Incas", "Teutons"},
	"fast castle":       {"Armenians", "Bohemians", "Burmese", "Dravidians", "Franks", "Goths", "Huns", "Incas", "Italians", "Japanese", "Khmer", "Lithuanians", "Mongols", "Poles", "Portuguese", "Spanish", "Turks", "Vikings"},
	"knight rush":       {"Berbers", "Bulgarians", "Burgudians", "Chinese", "Cumans", "Franks", "Georgians", "Huns", "Lithuanians", "Malians", "Persians", "Poles", "Slavs", "Tartars", "Teutons", "Vikings"},
	"booming":           {"Bengalis", "Berbers", "Burgudians", "Byzantines", "Cumans", "Georgians", "Gurjaras", "Hindustanis", "Italians", "Koreans", "Malians", "Persians", "Portuguese", "Saracens", "Sicilians", "Spanish", "Teutons", "Turks"},
	"turtle":            {"Byzanties", "Georgians", "Sicilians", "Koreans"},
	"crossbow rush":     {"Britons", "Chinese", "Ethiopians", "Mayans"},
	"trash war":         {"Berbers", "Magyars", "Malay", "Poles"},
	"castle drop":       {"Franks", "Malay", "Sicilians", "Malay"},
}

var StrategiesInfo = map[string]Strategy{
	"flush archer": {
		Name:        "Feudal Archer Rush",
		Emoji:       "🏹💀",
		Description: "Raid your opponent's economy as soon as possible with a small group of Archers",
		Pros:        "Archers are great at harrassing and countering infantry. This rush is effective at keeping your opponent away from their resources.",
		Cons:        "Slower than a Feudal Men-At-Arms Rush or a Scout Rush",
		Tips:        "For max efficiency, don't start your raid until you've massed at least 5-6 archers and researched Fletching",
	},
	"flush men-at-arms": {
		Name:        "Feudal Men-At-Arms Rush",
		Emoji:       "🛡️⚔️",
		Description: "Start your aggression early by sending your Militia to your enemy base while advancing to the Feudal Age and then immediately upgrading to Men-At-Arms",
		Pros:        "A very fast rush that can catch your enemy off guard and cripple their economy",
		Cons:        "Countered by walling strategies to limit economic impact. It can also be difficult to micromanage your infantry effectively.",
		Tips:        "The goal of this rush is to disrupt your opponent's economy. Killing villagers is a plus, but not necessarily the goal.",
	},
	"scout rush": {
		Name:        "Scout Rush",
		Emoji:       "🐴👣",
		Description: "Harrass your opponent with highly mobile Scout Cavalry as soon as possible to disrupt your enemy's economy",
		Pros:        "You are able to reach your enemy's base very quickly with Scout Cavalry. Scouts are mobile enough to chase down villagers and retreat from danger.",
		Cons:        "Effectively countered by walling and Spearmen",
		Tips:        "Always keep your Scouts out of the range of the enemy's Town Center",
	},
	"drush": {
		Name:        "Dark Age Rush",
		Emoji:       "⚔️👣",
		Description: "Create 3 Militia as soon as possible in the Dark Age and use them to distract your opponent",
		Pros:        "Fastest rush available to most civs. Limited resource investment. Can easily springboard from a Drush into a Fast Castle.",
		Cons:        "If opponent fends off your Drush quickly, you will be playing catch-up from an economic standpoint. Microing your Milita and developing your economy is difficult. Militia are not great raiding units.",
		Tips:        "In team games, it is preferred for one of the flanking players to initiate a drush as they are closer to an opponent. Make sure you scout your opponent's base in advance. The objective of a Drush is to disrupt an economy, not necessarily kill villagers (although that is a plus).",
	},
	"tower rush": {
		Name:        "Tower Rush",
		Emoji:       "🚧🗼",
		Description: "Disrupt your opponent's access to their resources by placing a tower in a strategic location in their base.",
		Pros:        "Incredibly disruptive to your opponent if properly executed.",
		Cons:        "Complicated build pattern that requires practice to execute correctly. Requires a big investment of villagers (~5). Failure to execute will result in a huge economic setback.",
		Tips:        "Scout your opponent's base well to identify ideal tower locations. The second tower you build will be to defend the first tower.",
	},
	"fast castle": {
		Name:        "Fast Castle Age",
		Emoji:       "⏩🏰",
		Description: "Reach the Castle Age as fast as possible.",
		Pros:        "This is a very flexible stratgey that can be parlayed into a number of late game strategies (Booming, Knight Rush, Crossbow Rush, etc.)",
		Cons:        "This strategy only works when you are reasonably protected early game as you will be very susceptible to early rushing.",
		Tips:        "Your Town Center(s) must always be producing villagers until you have around 100. Make sure you always have a large enough food income (Farms and/or Fishing Ships) to maintain this",
	},
	"knight rush": {
		Name:        "Knight Rush",
		Emoji:       "🐴🛡️",
		Description: "Reach the Castle Age as fast as possible and then begin producing knights to attack an opponent",
		Pros:        "Knights are a very powerful and effective raiding unit due to high mobility, damage, and HP. You can easily expand your economy with Town Centers after creating a few knights.",
		Cons:        "No military defenses against enemy rushes until after you reach the Castle Age and train knights",
		Tips:        "This strategy only works when you are reasonably protected early game as you will be very susceptible to early rushing.",
	},
	"booming": {
		Name:        "Booming",
		Emoji:       "🤑📈",
		Description: "Reach the Castle Age as fast as possible and then quickly expand your economy using multiple Town Centers",
		Pros:        "This builds the foundation for an extremely strong economy which will allow you to create a large military",
		Cons:        "No military defenses against early game rushes",
		Tips:        "Booming is tough to pull off unless you are protected by teammates or by a closed-off map (Black Forest)",
	},
	"turtle": {
		Name:        "Turtle",
		Emoji:       "🛡️🐢",
		Description: "Massing of armies and/or economies behind fortifications.",
		Pros:        "Effectively counters many early and mid game rushes.",
		Cons:        "Turtling by nature grants the turtler very little map control outside of the player's base, thus making long-term resource access more difficult. Late game siege is difficult to defend against.",
		Tips:        "This strategy is much more successfully executed on closed-off maps (ex: Black Forest) than open maps (Arabia). ",
	},
	"crossbow rush": {
		Name:        "Crossbow Rush",
		Emoji:       "❌🏹",
		Description: "Rush your enemy early with Crossbowmen right after reaching the Castle Age.",
		Pros:        "Very powerful rush that is effective against players attemtping to boom. Cheaper than a knight rush and can be micromanged more easily.",
		Cons:        "Can be countered by walls (turtling), cavalry, and skirmishers",
		Tips:        "Begin building your Archers in the Feudal Age (8-10 will suffice). Immediately research Bodkin Arrows from the Blacksmith and the Crossbowmen upgrade upon advancing to the Castle Age before attacking.",
	},
	"trash war": {
		Name:        "Trash War",
		Emoji:       "🗑️⚔️",
		Description: "A late game strategy/situation when gold is scarce and the remaining players rely on cheap units with little or no gold cost.",
		Pros:        "Some civilizations have bonuses that give them a great advantage in a trash war scenario. These advantages are typically quite difficult to counter without units that require gold.",
		Cons:        "This is not a viable strategy in team games due to trade. The strategy is also highly reliant on reaching the endgame.",
		Tips:        "Players relying on waging an effective trash war will need to survive until their opponent's gold reserves run low. It would also help if you limit your opponent's access to relics to reduce/eliminate gold trickle.",
	},
	"castle drop": {
		Name:        "Castle Drop",
		Emoji:       "⤵️🏯",
		Description: "Build a Castle inside or right next to your enemy's base as soon as you reach Castle Age",
		Pros:        "Powerful mechanism to take map control and cripple your opponent's economy. Also gives you the opportunit to build your unique units next to or inside your enemy's base.",
		Cons:        "High risk, high reward strategy. Failure to execute will result in a significant setback. Can be countered by walling and rushing.",
		Tips:        "Prioritize gathering stone early (~3-5 villagers on stone by Feudal). You will need around 7-10 villagers to build your castle effectively. Be sure to scout your opponent's base before moving your villagers, identifying any obstacles or ideal places to build your Castle.",
	},
}

// !strat!
func FormatStratOutput(strat string, s *discordgo.Session, m *discordgo.MessageCreate) {
	isCmdValid := utils.IsValidCmd(8, s, m)
	if !isCmdValid {
		return
	}

	stratRaw := strings.ToLower(m.Content[7:])
	strategy := StrategiesInfo[stratRaw]

	msg := fmt.Sprintf(
		`%v: %v
---------------------------
Description: %v
Pros: %v
Cons: %v
Tips: %v`, strategy.Name, strategy.Emoji, strategy.Description, strategy.Pros, strategy.Cons, strategy.Tips)
	_, err := s.ChannelMessageSend(m.ChannelID, msg)
	if err != nil {
		fmt.Printf("Error sending message to %v \n", msg)
	}
}

// !stratlist
func ListAllStrats(s *discordgo.Session, m *discordgo.MessageCreate) {
	strats := utils.GetAllKeys(StratToCivs)
	msg := fmt.Sprintf("%v", strats)
	_, err := s.ChannelMessageSend(m.ChannelID, msg)
	if err != nil {
		fmt.Printf("Error sending message to %v \n", msg)
	}
}

// !stratcivs
// "**!stratcivs <STRATEGY>** -> Returns all the civs that can employ a specified strategy effectively"
func CivsForStratOutput(strat string, s *discordgo.Session, m *discordgo.MessageCreate) {
	isCmdValid := utils.IsValidCmd(12, s, m)
	if !isCmdValid {
		return
	}

	stratRaw := strings.ToLower(m.Content[11:])
	civs := strings.Join(StratToCivs[stratRaw], ", ")

	msg := fmt.Sprintf("%v: %v", stratRaw, civs)
	_, err := s.ChannelMessageSend(m.ChannelID, msg)
	if err != nil {
		fmt.Printf("Error sending message to %v \n", msg)
	}
}
