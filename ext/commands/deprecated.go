package ext

import (
	"time"

	"atomicnicos.me/digital-overdose-bot/common"
	"github.com/bwmarrin/discordgo"
)

func WarnUserTest(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Processing",
		},
	})

	options := i.ApplicationCommandData().Options

	// Or convert the slice into a map
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	opt, ok := optionMap["user"]
	if !ok {
		return
	}

	m, _ := s.GuildMember(*common.GuildID, opt.UserValue(nil).ID)

	warnUsers([]discordgo.Member{*m}, map[string]time.Time{opt.UserValue(nil).ID: time.Now()}, s)

}
