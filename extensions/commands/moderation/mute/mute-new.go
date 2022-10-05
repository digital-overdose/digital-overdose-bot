package extensions

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"atomicnicos.me/digital-overdose-bot/common"
	database_utils "atomicnicos.me/digital-overdose-bot/db"
	"github.com/bwmarrin/discordgo"
)

func Mute(s *discordgo.Session, i *discordgo.InteractionCreate) {
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

	opt_duration, durationIsSet := optionMap["duration"]
	var (
		duration time.Duration
		err      error
	)
	if !durationIsSet {
		duration, err = time.ParseDuration("2540400h")
	} else {
		duration, err = time.ParseDuration(opt_duration.StringValue())
	}

	if err != nil {
		common.LogAndSend(fmt.Sprintf("[x] Could not parse duration '%v' because `%v`", opt_duration.StringValue(), err), s)
	}

	reason := ""
	opt_reason, reasonIsSet := optionMap["reason"]
	if reasonIsSet {
		reason = opt_reason.StringValue()
	} else {
		reason = "Reason not provided."
	}

	member, err := s.GuildMember(*common.GuildID, user.UserValue(nil).ID)
	if err != nil {
		common.LogAndSend(fmt.Sprintf("[x] Could not mute: (ID: %v) because `%v`", user.UserValue(nil).ID, err), s)
	}

	// TODO Check whether user is already muted.

	b := new(strings.Builder)
	json.NewEncoder(b).Encode(member.Roles)
	roles_str := b.String()

	roles := member.Roles
	for _, r := range roles {
		err := s.GuildMemberRoleRemove(*common.GuildID, member.User.ID, r)
		if err != nil {
			common.LogAndSend(fmt.Sprintf("[x] Could not remove role '%v' from '%v#%v' (ID: %v) because `%v`", r, member.User.Username, member.User.Discriminator, member.User.ID, err), s)
		}
	}

	err = s.GuildMemberRoleAdd(*common.GuildID, member.User.ID, *common.MuteRoleID)
	if err != nil {
		common.LogAndSend(fmt.Sprintf("[x] Could not add role 'Timeout' from '%v#%v' (ID: %v) because `%v`", member.User.Username, member.User.Discriminator, member.User.ID, err), s)
	}

	_, err = (*database_utils.Database).Methods.InsertMute.Exec(member.User.ID, time.Now(), time.Now().Add(duration), reason, roles_str)

	if err != nil {
		common.LogAndSend(fmt.Sprintf("[x] Could not mute: '%v#%v' (ID: %v) because `%v`", member.User.Username, member.User.Discriminator, member.User.ID, err), s)
		return
	}

	ActiveMutesRegistered += 1

	public_embed := buildPublicInsertMuteEmbed(member, i.Member)
	_, err = s.ChannelMessageSendEmbed(i.ChannelID, public_embed)

	if err != nil {
		log.Printf("Mute - Could not post embed because: %v", err)
	}

	private_embed := buildPrivateAInsertMuteEmbed(member, i.Member, reason, roles)
	_, err = s.ChannelMessageSendEmbed(*common.PrivateModLogsChannelID, private_embed)

	if err != nil {
		log.Printf("Mute - Could not post embed because: %v", err)
	}
}

func buildPublicInsertMuteEmbed(target *discordgo.Member, moderator *discordgo.Member) *discordgo.MessageEmbed {
	return common.BuildEmbed(
		fmt.Sprintf("Member '%v#%v' has been muted.", target.User.Username, target.User.Discriminator),
		fmt.Sprintf("Responsible moderator: '%v#%v'", moderator.User.Username, moderator.User.Discriminator),
		nil,
		nil,
	)
}

func buildPrivateAInsertMuteEmbed(target *discordgo.Member, moderator *discordgo.Member, reason string, roles []string) *discordgo.MessageEmbed {
	muteCount, _ := database_utils.GetTotalMutesCount()
	roleString := ""
	for i, r := range roles {
		roleString += fmt.Sprintf("<@%v>", r)
		if i != len(roles)-1 {
			roleString += ", "
		}
	}

	return common.BuildEmbed(
		fmt.Sprintf("Mute | Case %v", muteCount),
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
			{
				Name:   "Roles Removed",
				Value:  roleString,
				Inline: false,
			},
		},
		&discordgo.MessageEmbedFooter{Text: fmt.Sprintf("ID: %v", target.User.ID)},
	)
}
