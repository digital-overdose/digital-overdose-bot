package handler

import (
	"atomicmaya.me/digital-overdose-bot/src/common"
	"github.com/bwmarrin/discordgo"
)

func OnLeave(s *discordgo.Session, i *discordgo.GuildMemberRemove) {
	if ok := common.ShouldExecutionBeSkippedIfDev(true); ok {
		return
	}

	common.LogToServer(common.Log("[+] User has left the server: '%v' (%v)", common.FormatUsername(i.Member.User), i.Member.User.ID), s)
	common.LogToServer(common.Log("[+] User had roles: %#v", i.Member.Roles), s)
}
