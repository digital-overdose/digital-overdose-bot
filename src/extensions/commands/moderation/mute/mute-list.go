package extensions

import (
	"fmt"
	"log"

	"atomicmaya.me/digital-overdose-bot/src/common"
	"github.com/bwmarrin/discordgo"
)

func ListMutes(s *discordgo.Session, i *discordgo.InteractionCreate) {
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

	member, err := s.GuildMember(*common.GuildID, user.UserValue(nil).ID)

	if err != nil {
		common.LogAndSend(fmt.Sprintf("[x] Could not retrieve mute list for (ID: %v) because `%v`", user.UserValue(nil).ID, err), s)
	}

	log.Printf("[+] Retrieving mutes for member '%v#%v' (ID: %v)", member.User.Username, member.User.Discriminator, member.User.ID)

	// listMutes, err := getMutes(member.User.ID)
	err = fmt.Errorf("DEPRECATED FUNCTION")
	if err != nil {
		log.Printf("[x] Could not retrieve Mutes from database for member '%v#%v' (ID: %v) because '%v'", member.User.Username, member.User.Discriminator, member.User.ID, err)
		return
	}

	// summary := ""
	// description := fmt.Sprintf("User `%v#%v` has received the following mutes:", member.User.Username, member.User.Discriminator)
	// for i, mute := range listMutes {
	// 	summary += fmt.Sprintf("#%v: \"%v\" on <t:%v:f>\n", i+1, mute.MuteReason, mute.MuteTime.Unix())
	// }

	// if len(listMutes) == 0 {
	// 	description = fmt.Sprintf("User `%v#%v` has received no mutes.", member.User.Username, member.User.Discriminator)
	// 	summary = "No mutes recorded for this user."
	// }

	// embed := common.BuildEmbed(
	// 	fmt.Sprintf("Mutes for user `%v`.", member.User.Username),
	// 	description,
	// 	[]*discordgo.MessageEmbedField{
	// 		{
	// 			Name:   "Summary",
	// 			Value:  summary,
	// 			Inline: false,
	// 		},
	// 	},
	// 	&discordgo.MessageEmbedFooter{Text: fmt.Sprintf("ID: %v", member.User.ID)},
	// )

	// _, err = s.ChannelMessageSendEmbed(i.ChannelID, embed)

	// if err != nil {
	// 	log.Printf("ERR: %v", err)
	// 	common.LogAndSend("Mute - An error occured posting the embed.", s)
	// }
}

// func getMutes(userID string) ([]database_utils.Mute, error) {
// 	rows, err := (*database_utils.Database).Methods.ListMutes.Query(userID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer rows.Close()

// 	listMutes := []database_utils.Mute{}
// 	for rows.Next() {
// 		i := database_utils.Mute{}
// 		err := rows.Scan(&i.ID, &i.UserID, &i.MuteTime, &i.MuteExpiration, &i.MuteReason, &i.Roles, &i.Revoked)
// 		if err != nil {
// 			return nil, err
// 		}
// 		listMutes = append(listMutes, i)
// 	}

// 	return listMutes, nil
// }
