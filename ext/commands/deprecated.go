package ext

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

func WarnUserTest(s *discordgo.Session, i *discordgo.InteractionCreate) {
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

	warnUsers([]string{opt.UserValue(nil).ID}, map[string]time.Time{opt.UserValue(nil).ID: time.Now()}, s)
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Processing",
		},
	})

}
