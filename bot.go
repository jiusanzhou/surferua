package surferua

import (
	"math/rand"
	"time"
)

var botDBSize = 0
var botDB = []string{}

// Crawler and bot UA strings
type Bot struct {
	// The name of bot
	Name string

	// The version of bot
	Version string

	// The url of bot
	Url string
}

func (b *Bot) String() (s string) {
	return b.Name + "/" + b.Version + " (+" + b.Url + ")"
}

func NewBot() (string) {

	return botDB[rand.Intn(botDBSize)]
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
