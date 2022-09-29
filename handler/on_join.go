package handler

import (
	"fmt"
	"time"

	"atomicnicos.me/digital-overdose-bot/common"
	"github.com/bwmarrin/discordgo"
)

func OnJoin(s *discordgo.Session, i *discordgo.GuildMemberAdd) {
	if ok := common.ShouldExecutionBeSkippedIfDev(true); ok {
		return
	}

	common.LogAndSend(fmt.Sprintf("[+] New user has joined the server: '%v#%v' (%v)", i.Member.User.Username, i.Member.User.Discriminator, i.Member.User.ID), s)

	embed := &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Type:        discordgo.EmbedTypeRich,
		Title:       "Welcome to the Digital Overdose Discord!",
		Description: fmt.Sprintf("Hey there <@%v>, and welcome! To get started please take the time to do the following:", i.Member.User.ID),
		Fields: []*discordgo.MessageEmbedField{
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
		Thumbnail: &discordgo.MessageEmbedThumbnail{},
		Timestamp: time.Now().Format(time.RFC3339),
	}

	// TODO SEND THEM A PING
	_, err := s.ChannelMessageSendEmbed(*common.VerificationChannelID, embed)

	if err != nil {
		common.LogAndSend("OnJoin - An error occured posting the embed.", s)
	}

	// TODO SEND EMBED TO MOD LOGS
}
