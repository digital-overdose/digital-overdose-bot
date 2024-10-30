package extensions

import (
	"fmt"

	"atomicmaya.me/digital-overdose-bot/src/common"
	"github.com/bwmarrin/discordgo"
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

	err := fmt.Errorf("DEPRECATED FUNCTION")
	common.LogAndSend(fmt.Sprintf("Could not warn: (ID: %v) because `%v`", user.UserValue(nil).ID, err), s)

	// reason := ""
	// opt_reason, reasonIsSet := optionMap["reason"]
	// if reasonIsSet {
	// 	reason = opt_reason.StringValue()
	// } else {
	// 	reason = "No reason provided."
	// }

	// member, err := s.GuildMember(*common.GuildID, user.UserValue(nil).ID)

	// if err != nil {
	// 	common.LogAndSend(fmt.Sprintf("[x] Could not warn: (ID: %v) because `%v`", user.UserValue(nil).ID, err), s)
	// }

	// _, err = (*database_utils.Database).Methods.InsertWarn.Exec(member.User.ID, time.Now(), reason)

	// if err != nil {
	// 	common.LogAndSend(fmt.Sprintf("[x] Could not warn: '%v#%v' (ID: %v) because `%v`", member.User.Username, member.User.Discriminator, member.User.ID, err), s)
	// 	return
	// }

	// listWarns, err := getWarns(member.User.ID)
	// if err != nil {
	// 	log.Printf("[x] Could not retrieve Warns from database for member '%v#%v' (ID: %v)", member.User.Username, member.User.Discriminator, member.User.ID)
	// 	return
	// }

	// numberOfWarns := 0
	// for _, w := range listWarns {
	// 	if !w.Revoked {
	// 		numberOfWarns += 1
	// 	}
	// }

	// public_embed := buildPublicEmbed(member, i.Member, reason)
	// _, err = s.ChannelMessageSendEmbed(i.ChannelID, public_embed)
	// if err != nil {
	// 	common.LogAndSend("Warn - An error occured posting the embed.", s)
	// }

	// dmChannel, err := s.UserChannelCreate(member.User.ID)
	// if err != nil {
	// 	common.LogAndSend(fmt.Sprintf("Could not DM user %v", member.User.ID), s)
	// 	return
	// }

	// dm_embed := buildDMEmbed(member, i.Member, reason, numberOfWarns)
	// _, err = s.ChannelMessageSendEmbed(dmChannel.ID, dm_embed)
	// if err != nil {
	// 	common.LogAndSend("Warn - An error occured DM'ing the embed.", s)
	// }

	// private_embed := buildPrivateEmbed(member, i.Member, reason, numberOfWarns)
	// _, err = s.ChannelMessageSendEmbed(*common.PrivateModLogsChannelID, private_embed)
	// if err != nil {
	// 	common.LogAndSend("Warn - An error occured posting the embed to the private channel.", s)
	// }

	// common.LogAndSend(fmt.Sprintf("[👮] Member '%v#%v' (ID: %v) has been warned by `%v#%v` (ID: %v).", member.User.Username, member.User.Discriminator, member.User.ID, i.Member.User.Username, i.Member.User.Discriminator, i.Member.User.ID), s)
}

func buildPublicEmbed(target *discordgo.Member, moderator *discordgo.Member, reason string) *discordgo.MessageEmbed {
	return common.BuildEmbed(
		fmt.Sprintf("User `%v` has been warned.", target.User.Username),
		fmt.Sprintf("<@%v> - You have been warned by `%v#%v` for the following reason: '%v'", target.User.ID, moderator.User.Username, moderator.User.Discriminator, reason),
		nil,
		nil,
	)
}

func buildDMEmbed(target *discordgo.Member, moderator *discordgo.Member, reason string, nth int) *discordgo.MessageEmbed {
	return common.BuildEmbed(
		"You have received a warning in Digital Overdose",
		fmt.Sprintf("You have been warned by `%v#%v` for the following reason: '%v'", moderator.User.Username, moderator.User.Discriminator, reason),
		[]*discordgo.MessageEmbedField{
			{
				Name:   "Statistics",
				Value:  fmt.Sprintf("This is warn number `%v`.", nth),
				Inline: true,
			},
		},
		&discordgo.MessageEmbedFooter{Text: fmt.Sprintf("ID: %v", target.User.ID)},
	)
}

func buildPrivateEmbed(target *discordgo.Member, moderator *discordgo.Member, reason string, nth int) *discordgo.MessageEmbed {
	// warnCount, _ := database_utils.GetTotalWarnsCount()
	return common.BuildEmbed(
		fmt.Sprintf("Warn | Case %v", 0),
		fmt.Sprintf("Reason: \"%v\"", reason),
		[]*discordgo.MessageEmbedField{
			{
				Name: "Target", Value: fmt.Sprintf("%v#%v (ID: %v)\n<@%v>", target.User.Username, target.User.Discriminator, target.User.ID, target.User.ID), Inline: true,
			},
			{
				Name:   "Responsible Moderator",
				Value:  fmt.Sprintf("%v#%v (ID: %v)\n<@%v>", moderator.User.Username, moderator.User.Discriminator, moderator.User.ID, moderator.User.ID),
				Inline: true,
			},
		},
		&discordgo.MessageEmbedFooter{Text: fmt.Sprintf("ID: %v", target.User.ID)},
	)
}
