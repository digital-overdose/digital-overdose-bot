package extensions

import (
	"atomicmaya.me/digital-overdose-bot/src/common"
	"github.com/bwmarrin/discordgo"
)

func ServerStats(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if ok, _ := common.CheckHasPermissions(i, s, discordgo.PermissionAdministrator); !ok {
		return
	}

	if ok := common.ShouldExecutionBeSkippedIfDev(true); ok {
		return
	}

	// TODO FINISH

	common.LogToServer(common.Log("[⚠] Stats of `%s` executed by `%s`", *common.GuildID, common.FormatUsername(i.User)), s)

}
