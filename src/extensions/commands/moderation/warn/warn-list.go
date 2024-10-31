package extensions

import (
	"fmt"

	"atomicmaya.me/digital-overdose-bot/src/common"
	database_utils "atomicmaya.me/digital-overdose-bot/src/db"
	"github.com/bwmarrin/discordgo"
)

func ListWarns(s *discordgo.Session, i *discordgo.InteractionCreate) {
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
	common.LogToServer(common.Log("Could not list warns: (ID: %v) because `%v`", user.UserValue(nil).ID, err), s)

	// member, err := s.GuildMember(*common.GuildID, user.UserValue(nil).ID)

	// if err != nil {
	// 	common.LogAndSend(fmt.Sprintf("[x] Could not retrieve warn list for (ID: %v) because `%v`", user.UserValue(nil).ID, err), s)
	// }

	// log.Printf("[+] Retrieving warns for member '%v#%v' (ID: %v)", member.User.Username, member.User.Discriminator, member.User.ID)

	// listWarns, err := getWarns(member.User.ID)
	// if err != nil {
	// 	log.Printf("[x] Could not retrieve Warns from database for member '%v#%v' (ID: %v)", member.User.Username, member.User.Discriminator, member.User.ID)
	// 	return
	// }

	// summary := ""
	// description := fmt.Sprintf("User `%v#%v` has received the following warnings:", member.User.Username, member.User.Discriminator)
	// for i, warn := range listWarns {
	// 	revoked := ""
	// 	if warn.Revoked {
	// 		revoked = " *(revoked)*"
	// 	}
	// 	summary += fmt.Sprintf("#%v: \"%v\" on <t:%v:f>%v\n", i+1, warn.WarnReason, warn.WarnTime.Unix(), revoked)

	// }
	// if len(listWarns) == 0 {
	// 	description = fmt.Sprintf("User `%v#%v` has received no warnings.", member.User.Username, member.User.Discriminator)
	// 	summary = "No warns recorded for this user."
	// }

	// embed := common.BuildEmbed(
	// 	fmt.Sprintf("Warnings for user `%v`.", member.User.Username),
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
	// 	common.LogAndSend("Warn - An error occured posting the embed.", s)
	// }
}

func getWarns(userID string) ([]database_utils.Warn, error) {
	// rows, err := (*database_utils.Database).Methods.ListWarns.Query(userID)
	// if err != nil {
	// 	return nil, err
	// }

	// defer rows.Close()

	listWarns := []database_utils.Warn{}
	// for rows.Next() {
	// 	i := database_utils.Warn{}
	// 	err := rows.Scan(&i.ID, &i.UserID, &i.WarnTime, &i.WarnReason, &i.Revoked)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	listWarns = append(listWarns, i)
	// }

	return listWarns, nil
}
