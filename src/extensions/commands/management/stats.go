package extensions

import (
	"fmt"

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

	common.LogAndSend(fmt.Sprintf("[⚠] Stats of `%s` executed by `%s`", *common.GuildID, i.User.ID), s)

}
