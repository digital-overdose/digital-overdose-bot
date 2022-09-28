package extensions

import (
	"fmt"
	"log"
	"time"

	"atomicnicos.me/digital-overdose-bot/common"
	"github.com/bwmarrin/discordgo"

	database_utils "atomicnicos.me/digital-overdose-bot/db"
)

func Warn(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if ok, _ := common.CheckHasPermissions(i, s, discordgo.PermissionKickMembers); !ok {
		return
	}

	if ok := common.ShouldExecutionBeSkippedIfDev(false); ok {
		return
	}

	options := i.ApplicationCommandData().Options
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	user, ok := optionMap["user"]
	if !ok {
		return
	}

	reason, ok := optionMap["reason"]
	if !ok {
		return
	}

	member, err := s.GuildMember(*common.GuildID, user.UserValue(nil).ID)

	if err != nil {
		common.LogAndSend(fmt.Sprintf("[x] Could not warn: (ID: %v) because `%v`", user.UserValue(nil).ID, err), s)
	}

	_, err = (*database_utils.Database).Methods.InsertWarn.Exec(member.User.ID, time.Now(), reason.StringValue())

	if err != nil {
		common.LogAndSend(fmt.Sprintf("[x] Could not warn: '%v#%v' (ID: %v) because `%v`", member.User.Username, member.User.Discriminator, member.User.ID, err), s)
		return
	}

	rows, err := (*database_utils.Database).Methods.ListWarns.Query(member.User.ID)
	if err != nil {
		return
	}

	defer rows.Close()

	listWarns, err := getWarns(member.User.ID)
	if err != nil {
		log.Printf("[x] Could not retrieve Warns from database for member '%v#%v' (ID: %v)", member.User.Username, member.User.Discriminator, member.User.ID)
		return
	}

	embed := &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Type:        discordgo.EmbedTypeRich,
		Title:       fmt.Sprintf("User `%v` has been warned.", member.User.Username),
		Description: fmt.Sprintf("<@%v> - You have been warned by `%v#%v` for the following reason: '%v'", member.User.ID, i.Member.User.Username, i.Member.User.Discriminator, reason.StringValue()),
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Statistics",
				Value:  fmt.Sprintf("User '%v#%v' now has `%v` warnings.", member.User.Username, member.User.Discriminator, len(listWarns)),
				Inline: true,
			},
		},
		Footer:    &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("ID: %v", member.User.ID)},
		Thumbnail: &discordgo.MessageEmbedThumbnail{},
		Timestamp: time.Now().Format(time.RFC3339),
	}

	_, err = s.ChannelMessageSendEmbed(i.ChannelID, embed)

	if err != nil {
		common.LogAndSend("Warn - An error occured posting the embed.", s)
	}

	dmChannel, err := s.UserChannelCreate(member.User.ID)
	if err != nil {
		common.LogAndSend(fmt.Sprintf("Could not DM user %v", member.User.ID), s)
		return
	}

	_, err = s.ChannelMessageSendEmbed(dmChannel.ID,
		&discordgo.MessageEmbed{
			Author:      &discordgo.MessageEmbedAuthor{},
			Type:        discordgo.EmbedTypeRich,
			Title:       "You have received a warning.",
			Description: fmt.Sprintf("You have been warned by `%v#%v` for the following reason: '%v'", i.Member.User.Username, i.Member.User.Discriminator, reason.StringValue()),
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Statistics",
					Value:  fmt.Sprintf("This is warn number `%v`.", len(listWarns)),
					Inline: true,
				},
			},
			Footer:    &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("ID: %v", member.User.ID)},
			Thumbnail: &discordgo.MessageEmbedThumbnail{},
			Timestamp: time.Now().Format(time.RFC3339),
		})

	if err != nil {
		common.LogAndSend("Warn - An error occured DM'ing the embed.", s)
	}

	common.LogAndSend(fmt.Sprintf("[👮] Member '%v#%v' (ID: %v) has been warned by `%v#%v` (ID: %v).", member.User.Username, member.User.Discriminator, member.User.ID, i.Member.User.Username, i.Member.User.Discriminator, i.Member.User.ID), s)
}
