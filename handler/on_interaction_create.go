package handler

import (
	"atomicmaya.me/digital-overdose-bot/extensions"

	"github.com/bwmarrin/discordgo"
)

func OnInteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if h, ok := extensions.CommandHandlers[i.ApplicationCommandData().Name]; ok {
		h(s, i)
	}

}
