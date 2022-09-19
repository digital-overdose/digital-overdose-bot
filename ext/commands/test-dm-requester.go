package ext

import (
	"atomicnicos.me/digital-overdose-bot/common"
	"github.com/bwmarrin/discordgo"
)

func TestDMRequester(s *discordgo.Session, i *discordgo.InteractionCreate) {
	common.CheckHasPermissions(i, s, discordgo.PermissionViewAuditLogs|discordgo.PermissionManageRoles)

	dmChannel, _ := s.UserChannelCreate(i.Member.User.ID)
	_, _ = s.ChannelMessageSend(dmChannel.ID, "SOME BLOODY FUCKING MESSAGE")

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Sent you a DM :wink:",
		},
	})
}
