package extensions

import (
	"fmt"
	"log"

	"atomicmaya.me/digital-overdose-bot/common"
	database_utils "atomicmaya.me/digital-overdose-bot/db"
	"github.com/bwmarrin/discordgo"
)

func Unwarn(s *discordgo.Session, i *discordgo.InteractionCreate) {
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
		common.LogAndSend(fmt.Sprintf("[x] Could not remove warn: (ID: %v) because `%v`", user.UserValue(nil).ID, err), s)
	}

	which, whichSet := optionMap["which"]

	log.Printf("[+] Retrieving warns for member '%v#%v' (ID: %v)", member.User.Username, member.User.Discriminator, member.User.ID)

	listWarns, err := getWarns(member.User.ID)
	if err != nil {
		log.Printf("[x] Could not retrieve Warns from database for member '%v#%v' (ID: %v)", member.User.Username, member.User.Discriminator, member.User.ID)
		return
	}

	if len(listWarns) == 0 {
		s.ChannelMessageSend(i.ChannelID, fmt.Sprintf("Member '%v#%v' (ID: %v) has no warnings.", member.User.Username, member.User.Discriminator, member.User.ID))
		return

	}

	//log.Printf("WHICHSET %v WHICH %v", whichSet, which)
	var rowID int
	if whichSet {
		if which.IntValue() > 0 && int(which.IntValue()) < len(listWarns) {
			rowID = listWarns[int(which.IntValue())-1].ID
		} else {
			s.ChannelMessageSend(i.ChannelID, "You have selected an out-of-bounds warning. Please try again.")
			return
		}
	} else {
		rowID = listWarns[len(listWarns)-1].ID
	}

	_, err = (*database_utils.Database).Methods.RemoveWarn.Exec(rowID)

	if err != nil {
		common.LogAndSend(fmt.Sprintf("[x] Could not unwarn: '%v#%v' (ID: %v) because `%v`", member.User.Username, member.User.Discriminator, member.User.ID, err), s)
		return
	}

	common.LogAndSend(fmt.Sprintf("[+] Successfully unwarned: '%v#%v' (ID: %v)", member.User.Username, member.User.Discriminator, member.User.ID), s)
}
