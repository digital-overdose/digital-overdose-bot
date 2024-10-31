package extensions

import (
	"atomicmaya.me/digital-overdose-bot/src/common"
	"github.com/bwmarrin/discordgo"
)

func LookupMember(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if ok, _ := common.CheckHasPermissions(i, s, discordgo.PermissionAdministrator); !ok {
		return
	}

	if ok := common.ShouldExecutionBeSkippedIfDev(true); ok {
		return
	}

	options := i.ApplicationCommandData().Options
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	userID := ""
	user, ok := optionMap["user"]
	if !ok {
		userID = optionMap["user_id"].StringValue()
		if !ok {
			return
		}
	} else {
		userID = user.UserValue(nil).ID
	}

	common.LogToServer(common.Log("[⚠] Lookup of `%s` executed by `%s`", userID, common.FormatUsername(i.User)), s)

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "LOOKUP:ACK",
			Flags:   1 << 6,
		},
	})
}
