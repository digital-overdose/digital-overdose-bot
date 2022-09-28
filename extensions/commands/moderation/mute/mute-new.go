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

	opt_duration, ok := optionMap["duration"]
	if !ok {
		return
	}
	duration, err := time.ParseDuration(opt_duration.StringValue())
	if err != nil {
		common.LogAndSend(fmt.Sprintf("[x] Could not parse duration '%v' because `%v`", opt_duration.StringValue(), err), s)
	}

	reason := ""
	opt_reason, reasonNotSet := optionMap["reason"]
	if reasonNotSet {
		reason = opt_reason.StringValue()
	}

	member, err := s.GuildMember(*common.GuildID, user.UserValue(nil).ID)

	if err != nil {
		common.LogAndSend(fmt.Sprintf("[x] Could not mute: (ID: %v) because `%v`", user.UserValue(nil).ID, err), s)
	}

	b := new(strings.Builder)
	json.NewEncoder(b).Encode(member.Roles)
	roles_str := b.String()

	log.Printf("MEMBERID %v, TIME %v, REASON %v, ROLES %v",
		member.User.ID,
		time.Now().Add(duration),
		reason,
		roles_str)

	_, err = (*database_utils.Database).Methods.InsertMute.Exec(member.User.ID, time.Now().Add(duration), reason, roles_str)

	if err != nil {
		common.LogAndSend(fmt.Sprintf("[x] Could not mute: '%v#%v' (ID: %v) because `%v`", member.User.Username, member.User.Discriminator, member.User.ID, err), s)
		return
	}

}
