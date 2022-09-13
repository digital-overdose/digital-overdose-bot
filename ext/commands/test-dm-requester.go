package ext

import (
	"log"

	"atomicnicos.me/go-bot/common"
	"github.com/bwmarrin/discordgo"
)

func TestDMRequester(s *discordgo.Session, i *discordgo.InteractionCreate) {
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
		return
	}

	dmChannel, _ := s.UserChannelCreate(i.Member.User.ID)
	_, _ = s.ChannelMessageSend(dmChannel.ID, "SOME BLOODY FUCKING MESSAGE")

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Sent you a DM :wink:",
		},
	})
}
