package handler

import (
	"log"

	"atomicnicos.me/digital-overdose-bot/common"
	"github.com/bwmarrin/discordgo"
)

func OnReady(s *discordgo.Session, r *discordgo.Ready) {
	log.Printf("Running in version: v%v", common.VERSION)
	log.Printf("Logged in as: '%v#%v'", s.State.User.Username, s.State.User.Discriminator)
}
