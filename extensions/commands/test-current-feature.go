package extensions

import (
	"atomicnicos.me/digital-overdose-bot/common"
	"github.com/bwmarrin/discordgo"
)

func TestCurrentFeature(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if ok := common.ShouldExecutionBeSkippedIfDev(false); ok {
		return
	}

	/*

		embed := &discordgo.MessageEmbed{
			Author:      &discordgo.MessageEmbedAuthor{},
			Color:       0xAD7EC2,
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

		_, err := s.ChannelMessageSendEmbed(*common.DebugChannelID, embed)

		if err != nil {
			common.LogAndSend("OnJoin - An error occured posting the embed.", s)
		}*/
}
