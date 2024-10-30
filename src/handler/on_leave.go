package handler

import (
	"fmt"

	"atomicmaya.me/digital-overdose-bot/src/common"
	"github.com/bwmarrin/discordgo"
)

func OnLeave(s *discordgo.Session, i *discordgo.GuildMemberRemove) {
	if ok := common.ShouldExecutionBeSkippedIfDev(true); ok {
		return
	}

	common.LogAndSend(fmt.Sprintf("[+] User has left the server: '%v' (%v)", common.FormatUsername(i.Member.User), i.Member.User.ID), s)
	common.LogAndSend(fmt.Sprintf("[+] User had roles: %#v", i.Member.Roles), s)
}
