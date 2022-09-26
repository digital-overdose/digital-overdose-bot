package handler

import (
	"atomicnicos.me/digital-overdose-bot/common"
	"github.com/bwmarrin/discordgo"
)

func OnMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if *common.CURRENTLY_DEV {
		//log.Printf("New message by %v (%v) in %v : %v\n> %v\n", m.Author.String(), m.Author.ID, m.ChannelID, m.GuildID, m.Content)
	}
}
