package ext

import (
	"log"

	"atomicnicos.me/digital-overdose-bot/common"
	"github.com/bwmarrin/discordgo"
)

func TestCurrentFeature(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Triggered by user-interaction
	if i != nil {
		_, err := common.HasPermissions(i, s, discordgo.PermissionViewAuditLogs|discordgo.PermissionManageRoles)
		if err != nil {
			log.Println("Error checking permissions.")
		}

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Processing",
			},
		})
	}

	//_, _ = s.ChannelMessageSend(*common.VerificationChannelID, "TEST")
	_, _ = s.ChannelMessageSend(*common.DebugChannelID, "Functional")
	//_, _ = s.ChannelMessageSend(*common.ModActionLogsChannelID, "TEST")
}
