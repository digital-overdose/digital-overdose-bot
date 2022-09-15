package ext

import (
	"atomicnicos.me/digital-overdose-bot/common"
	"github.com/bwmarrin/discordgo"
)

func TestCurrentFeature(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Processing",
		},
	})

	_, _ = s.ChannelMessageSend(*common.VerificationChannelID, "TEST")
	_, _ = s.ChannelMessageSend(*common.DebugChannelID, "TEST")
	_, _ = s.ChannelMessageSend(*common.ModActionLogsChannelID, "TEST")
}
