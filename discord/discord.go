package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/pkanti/v2/config"
	"log"
)

var session *discordgo.Session

func init() {
	var err error
	session, err = discordgo.New("Bot " + config.Config.Discord.Token)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Discord connected")
}
