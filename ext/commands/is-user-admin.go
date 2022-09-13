package ext

import (
	"log"

	"atomicnicos.me/digital-overdose-bot/common"
	"github.com/bwmarrin/discordgo"
)

func IsUserAdmin(s *discordgo.Session, i *discordgo.InteractionCreate) {
	ok, err := common.HasPermissions(i, s, discordgo.PermissionViewAuditLogs|discordgo.PermissionManageRoles)
	if err != nil {
		log.Println("Error checking permissions.")
	}

	if !ok {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "You're not STAFF",
			},
		})
	} else {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "You're STAFF, yaaay",
			},
		})
	}
}
