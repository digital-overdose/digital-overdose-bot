package handler

import (
	"fmt"

	"atomicmaya.me/digital-overdose-bot/src/common"
	"github.com/bwmarrin/discordgo"
)

func OnJoin(s *discordgo.Session, i *discordgo.GuildMemberAdd) {
	if ok := common.ShouldExecutionBeSkippedIfDev(true); ok {
		return
	}

	common.LogToServer(common.Log("[+] New user has joined the server: '%v#%v' (%v)", i.Member.User.Username, i.Member.User.Discriminator, i.Member.User.ID), s)

	i.Member.Roles = append(i.Member.Roles, *common.VerificationRoleID)

	publicEmbed := publicOnJoinEmbed(i.Member)
	_, err := s.ChannelMessageSendEmbed(*common.VerificationChannelID, publicEmbed)
	if err != nil {
		common.LogToServer(common.Log("[❌] OnJoin - An error occured posting the embed: \"%v\"", err), s)
	}

	_, err = s.ChannelMessageSend(*common.VerificationChannelID, fmt.Sprintf("<@%v> ^", i.Member.User.ID))
	if err != nil {
		common.LogToServer(common.Log("[❌] OnJoin - An error occured posting the embed: \"%v\"", err), s)
	}

	_, err = s.ChannelMessageSendEmbed(*common.PrivateModLogsChannelID, privateOnJoinEmbed(i.Member))
}

func publicOnJoinEmbed(member *discordgo.Member) *discordgo.MessageEmbed {
	return common.BuildEmbed(
		"Welcome to the Digital Overdose Discord!",
		fmt.Sprintf("Hey there <@%v>, and welcome! To get started please take the time to do the following:", member.User.ID),
		[]*discordgo.MessageEmbedField{
			{
				Name:   "Step 1",
				Value:  "Review <#789513212109258782> and tell us you agree and accept them :white_check_mark: in the <#687238387463094317> channel.",
				Inline: true,
			},
			{
				Name:   "Step 2",
				Value:  "Tell us a bit about yourself here!",
				Inline: true,
			},
			{
				Name:   "Step 3",
				Value:  "Wait. A staff member will get to you soon! :slight_smile:",
				Inline: true,
			},
		},
		nil,
	)
}

func privateOnJoinEmbed(member *discordgo.Member) *discordgo.MessageEmbed {
	return common.BuildEmbed(
		fmt.Sprintf("Member Joined - `%v`", member.User.ID),
		fmt.Sprintf("'%v#%v' (ID: %v)", member.User.Username, member.User.Discriminator, member.User.ID),
		nil,
		nil,
	)
}
