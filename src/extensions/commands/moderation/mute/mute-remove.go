package extensions

import (
	"encoding/json"
	"fmt"

	"atomicmaya.me/digital-overdose-bot/src/common"
	"github.com/bwmarrin/discordgo"
)

func UnmuteManual(s *discordgo.Session, i *discordgo.InteractionCreate) {
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
	common.LogToServer(common.Log("Could not unmute: (ID: %v) because `%v`", user.UserValue(nil).ID, err), s)

	// member, err := s.GuildMember(*common.GuildID, user.UserValue(nil).ID)
	// if err != nil {
	// 	common.LogAndSend(fmt.Sprintf("[❌] Could not unmute: (ID: %v) because `%v`", user.UserValue(nil).ID, err), s)
	// }

	// rows, err := (*database_utils.Database).Methods.GetCurrentMute.Query(member.User.ID)
	// if err != nil {
	// 	common.LogAndSend(fmt.Sprintf("[❌] ERROR: Automated Unmute - '%v'", err), s)
	// 	return
	// }

	// defer rows.Close()

	// for rows.Next() {
	// 	i := database_utils.Mute{}
	// 	err := rows.Scan(&i.ID, &i.UserID, &i.MuteTime, &i.MuteExpiration, &i.MuteReason, &i.Roles, &i.Revoked)
	// 	if err != nil {
	// 		common.LogAndSend(fmt.Sprintf("[❌] ERROR: Unmute - '%v'", err), s)
	// 		return
	// 	}

	// 	restoreUser(s, fmt.Sprint(i.UserID), i.Roles, i.ID, false)
	// }
}

func UnmuteAutomated(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// This command can abort early, saving CPU processing power aka. money.
	if ActiveMutesRegistered == 0 {
		common.LogToServer(common.Log("[🕑] No recent mutes, cron job aborting."), s)
		return
	}

	// counts, err := (*database_utils.Database).Methods.CountActiveMutes.Query()
	// if err != nil {
	// 	common.LogAndSend(fmt.Sprintf("[❌] ERROR: Automated Unmute - '%v'", err), s)
	// 	return
	// }

	// defer counts.Close()

	// list_counts := []database_utils.Count{}
	// for counts.Next() {
	// 	i := database_utils.Count{}
	// 	err := counts.Scan(&i.Count)

	// 	if err != nil {
	// 		common.LogAndSend(fmt.Sprintf("[❌] ERROR: Automated Unmute - '%v'", err), s)
	// 		return
	// 	}
	// 	list_counts = append(list_counts, i)
	// }

	// ActiveMutesRegistered = list_counts[0].Count
	// if ActiveMutesRegistered == 0 {
	// 	return
	// }

	// rows, err := (*database_utils.Database).Methods.ActiveMutes.Query()
	// if err != nil {
	// 	common.LogAndSend(fmt.Sprintf("[❌] ERROR: Automated Unmute - '%v'", err), s)
	// 	return
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	i := database_utils.Mute{}
	// 	err := rows.Scan(&i.ID, &i.UserID, &i.MuteTime, &i.MuteExpiration, &i.MuteReason, &i.Roles, &i.Revoked)
	// 	if err != nil {
	// 		common.LogAndSend(fmt.Sprintf("[❌] ERROR: Automated Unmute - '%v'", err), s)
	// 		return
	// 	}

	// 	if i.MuteExpiration.Before(time.Now()) {
	// 		restoreUser(s, fmt.Sprint(i.UserID), i.Roles, i.ID, true)
	// 	}
	// }
}

func restoreUser(s *discordgo.Session, userID string, roles_encoded string, mute_case int, automated bool) {
	var roles []string
	err := json.Unmarshal([]byte(roles_encoded), &roles)
	if err != nil {
		common.LogToServer(common.Log("[❌] ERROR: Unmute - Roles Unmarshaling - '%v'", err), s)
	}

	err = s.GuildMemberRoleRemove(*common.GuildID, userID, *common.MuteRoleID)
	if err != nil {
		common.LogToServer(common.Log("[❌] ERROR: Unmute - Removing 'Timeout' Role - '%v'", err), s)
	}

	for _, r := range roles {
		err = s.GuildMemberRoleAdd(*common.GuildID, userID, r)
		if err != nil {
			common.LogToServer(common.Log("[❌] ERROR: Unmute - Adding Role `%v` - '%v'", r, err), s)
		}
	}

	// _, err = (*database_utils.Database).Methods.RemoveMute.Exec(userID)
	// if err != nil {
	// 	common.LogAndSend(fmt.Sprintf("[❌] ERROR: Unmute - Clearing Database - '%v'", err), s)
	// }

	// common.SendEmbed(s, *common.PrivateModLogsChannelID, buildPrivateMuteRemoveEmbed(
	// 	userID,
	// 	automated,
	// 	mute_case,
	// 	roles,
	// ))
}

func buildPrivateMuteRemoveEmbed(userID string, automated bool, mute_case int, roles []string) *discordgo.MessageEmbed {
	roleString := ""
	for i, r := range roles {
		roleString += fmt.Sprintf("<@%v>", r)
		if i != len(roles)-1 {
			roleString += ", "
		}
	}

	auto := ""
	if automated {
		auto = "Timer expired"
	} else {
		auto = "Manual"
	}

	return common.BuildEmbed(
		fmt.Sprintf("Unmute | Case %v", mute_case),
		fmt.Sprintf("Method: %v", auto),
		[]*discordgo.MessageEmbedField{
			{
				Name: "Target", Value: fmt.Sprintf("<@%v> (ID: %v)", userID, userID), Inline: false,
			},
			{
				Name:   "Roles Added",
				Value:  roleString,
				Inline: false,
			},
		},
		&discordgo.MessageEmbedFooter{Text: fmt.Sprintf("ID: %v", userID)},
	)
}
