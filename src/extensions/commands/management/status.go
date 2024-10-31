package extensions

import (
	"atomicmaya.me/digital-overdose-bot/src/common"
	"github.com/bwmarrin/discordgo"
)

func BotStatus(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if ok, _ := common.CheckHasPermissions(i, s, discordgo.PermissionAdministrator); !ok {
		return
	}

	if ok := common.ShouldExecutionBeSkippedIfDev(true); ok {
		return
	}

	// TODO

	common.LogToServer(common.Log("[⚠] Bot status info executed by `%s`", common.FormatUsername(i.User)), s)

}
