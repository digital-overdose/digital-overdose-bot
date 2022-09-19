package ext

import (
	"log"

	"atomicnicos.me/digital-overdose-bot/common"
	"github.com/bwmarrin/discordgo"
)

func TestCurrentFeature(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if ok := common.ShouldExecutionBeSkippedIfDev(false); ok {
		return
	}

	// Triggered by user-interaction
	if i != nil {
		if ok, _ := common.CheckHasPermissions(i, s, discordgo.PermissionViewAuditLogs|discordgo.PermissionManageRoles); !ok {
			return
		}
	}

	log.Print("AAAA")

	_, _ = s.ChannelMessageSend("1021708685820559391", "TEST")
	//_, _ = s.ChannelMessageSend(*common.DebugChannelID, "It's been 1 minute since the last message, have another one!")
	//_, _ = s.ChannelMessageSend(*common.ModActionLogsChannelID, "TEST")
}
