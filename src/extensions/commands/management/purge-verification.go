package extensions

import (
	"fmt"
	"log"
	"time"

	"atomicmaya.me/digital-overdose-bot/src/common"

	"github.com/bwmarrin/discordgo"
)

// Purges the Verification channel of lurkers (> 7 days) or warns them (> 5 days)
func PurgeVerification(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if ok := common.ShouldExecutionBeSkippedIfDev(true); ok {
		return
	}
	// If triggered by user-interaction
	if i != nil {
		if ok, _ := common.CheckHasPermissions(i, s, discordgo.PermissionManageRoles); !ok {
			return
		}
	} else {
		common.LogAndSend(":robot: :rotating_light: `/purge-verification` triggered by cron.", s)
	}

	var (
		listing           = true // Loop controller.
		lastID            = ""   // Last member ID that was mapped (used as offset).
		messagesProcessed = 0    // The number of messages processed.
		lastMessageID     = ""   // Last message ID processed (used as offset).
		messageLimit      = 1000 // How many messages to test.

		candidates       []discordgo.Member   = []discordgo.Member{}   // The members with the verification role.
		candidatesKick   []discordgo.Member   = []discordgo.Member{}   // The ^ members that have exceeded the 7-day non-verification limit.
		candidatesWarn   []discordgo.Member   = []discordgo.Member{}   // The ^ members that are nearing the 7-day non-verification limit (days 5 and 6)
		timeLastSeen     map[string]time.Time = map[string]time.Time{} // When the ^ members were last seen (shortest time between joining or last message time)
		timeLastMessaged map[string]time.Time = map[string]time.Time{} // When the ^ members last messaged
	)

	// Retrieve server information
	g, err := s.State.Guild(*common.GuildID)
	if err != nil {
		log.Panicf("Unable to get Guild %v: %v", *common.GuildID, err)
	}

	common.LogAndSend("[+] Starting Member Scan...", s)

	// Retrieves a list of all the server members.
	for listing {
		members, err := s.GuildMembers(g.ID, lastID, 1000)
		if err != nil {
			log.Panicf("Unable to get Members of Guild %v: %v", *common.GuildID, err)
		}

		common.LogAndSend(fmt.Sprintf("\tGot %v Members from Guild %v!", len(members), g.ID), s)

		for _, m := range members {
			for _, r := range m.Roles {
				if r == *common.VerificationRoleID {
					candidates = append(candidates, *m)
					if _, ok := timeLastSeen[m.User.ID]; !ok {
						timeLastSeen[m.User.ID] = m.JoinedAt
					}
				}
			}
		}

		if len(members) < 1000 {
			listing = false
		} else {
			lastID = members[len(members)-1].User.ID
		}
	}

	common.LogAndSend("[+] Finished Member Scan...", s)

	common.LogAndSend(fmt.Sprintf("[+] Got %v Deletion Candidates from Guild %v!\n", len(candidates), g.ID), s)

	common.LogAndSend("[+] Scanning verification channel messages.", s)

	// Processes the last 1000 messages (in batches of 100) in reverse-chronological order.
	// Adds them to the list of the members having messages (read: interacted).
	counterMessageID := ""
	for messagesProcessed < messageLimit {
		messages, err := s.ChannelMessages(*common.VerificationChannelID, 100, lastMessageID, "", "")
		if err != nil {
			log.Panicf("Unable to get Messages of Guild %v: %v", *common.GuildID, err)
		}

		messagesProcessed += len(messages)
		lastMessageID = messages[len(messages)-1].ID

		if messagesProcessed == 100 {
			counterMessageID = common.LogAndSend(fmt.Sprintf("\tGot %v messages (%v/%v)!", len(messages), messagesProcessed, messageLimit), s)
		} else {
			common.LogAndSend(fmt.Sprintf("\tGot %v messages (%v/%v)!", len(messages), messagesProcessed, messageLimit), s, "")
			if *common.DebugChannelID != "" {
				s.ChannelMessageEdit(*common.DebugChannelID, counterMessageID, fmt.Sprintf("\tGot %v messages (%v/%v)!", len(messages), messagesProcessed, messageLimit))
			}
		}

		for _, msg := range messages {
			if _, ok := timeLastMessaged[msg.Author.ID]; !ok {
				timeLastMessaged[msg.Author.ID] = msg.Timestamp
			}
		}
	}

	for k, v := range timeLastMessaged {
		timeLastSeen[k] = v
	}

	common.LogAndSend(fmt.Sprintf("[+] Got %v members having posted a message.", len(timeLastSeen)), s)

	common.LogAndSend("[∨] Starting Candidate Filtering...", s)

	now := time.Now()
	kickDate := now.Add(-7 * 24 * time.Hour)
	warnDate := now.Add(-5 * 24 * time.Hour)

	// Filters the list in order to determine which users are eligible for kick, and which ones for warn.
	for _, c := range candidates {
		if timeLastSeen[c.User.ID].Before(kickDate) {
			candidatesKick = append(candidatesKick, c)
		} else if timeLastSeen[c.User.ID].Before(warnDate) {
			candidatesWarn = append(candidatesWarn, c)
		}
	}
	common.LogAndSend("[+] Finished Candidate Filtering...", s)

	// DMs and Kicks the candidates.
	for _, candidate := range candidatesKick {
		common.LogAndSend(fmt.Sprintf("[----] User %v will be kicked. Last interaction: <t:%v:f>", candidate.User.Username, timeLastSeen[candidate.User.ID].Unix()), s)

		sendDMToUser(candidate, s)
		kickUser(candidate, s)
	}

	// Generate the warn message for the warned candidates.
	for _, candidate := range candidatesWarn {
		common.LogAndSend(fmt.Sprintf("[----] User %v will be warned. Last interaction: <t:%v:f>", candidate.User.Username, timeLastSeen[candidate.User.ID].Unix()), s)
	}

	if len(candidatesWarn) > 0 {
		warnUsers(candidatesWarn, timeLastSeen, s)
	}

	// Writes a report to the specified mod-action-logs channel.
	sendModActionLogsMessage(candidatesKick, candidatesWarn, s)
	common.LogAndSend("[✓] Done!", s)
}

// Generates and posts the warn message.
func warnUsers(candidatesWarn []discordgo.Member, timeLastSeen map[string]time.Time, s *discordgo.Session) (bool, error) {
	formatted_users := "`Not seen in server since`   |   User\n"
	for _, candidate := range candidatesWarn {
		formatted_users += fmt.Sprintf("<t:%v:f>   -   %v\n", timeLastSeen[candidate.User.ID].Unix(), fmt.Sprintf("<@%v>", candidate.User.ID))
	}

	formatted_msg := fmt.Sprintf(":warning: **Verification** :warning:\n\nThe following users have yet to verify!\n\n%v\nPlease verify! In order to verify:\n\t1) ✅ Please accept the rules by posting your acceptance here in the verification channel\n\t2) Please tell us a bit about yourself.\n\nFailure to verify will result in you being kicked in less than 2 days.", formatted_users)

	_, err := s.ChannelMessageSend(*common.VerificationChannelID, formatted_msg)

	return err != nil, err
}

// Sends an explicative DM to a specified user, with a rejoin link.
func sendDMToUser(candidateKick discordgo.Member, s *discordgo.Session) (bool, error) {
	dmChannel, err := s.UserChannelCreate(candidateKick.User.ID)
	if err != nil {
		common.LogAndSend(fmt.Sprintf("Could not DM user %v", candidateKick.User.ID), s)
		return err != nil, err
	}

	_, err = s.ChannelMessageSend(dmChannel.ID, "*This message is sent to you by the Digital Overdose Server's Management Bot.*\n\nHello!\nWe wish to inform you that you have been automatically pruned from the server due to not completing the steps required for verification.\n\nIf you wish to rejoin (and this time complete verification, it only takes 2 minutes), you may use this link: https://discord.gg/jaNuDB95Yd\n\nHave a nice day!")

	return err != nil, err
}

// Kicks a user.
func kickUser(candidateKick discordgo.Member, s *discordgo.Session) (bool, error) {
	err := s.GuildMemberDeleteWithReason(*common.GuildID, candidateKick.User.ID, fmt.Sprintf("User %v failed to verify in time!", candidateKick.User.Username))
	if err != nil {
		return err != nil, err
	}
	return err != nil, err
}

// Writes an itemized message in #mod-action-logs.
func sendModActionLogsMessage(candidatesKick []discordgo.Member, warnedUsers []discordgo.Member, s *discordgo.Session) {
	if len(*common.ModActionLogsThreadID) == 0 {
		log.Printf("Channel ModActionLogs not set, skipping message.")
		return
	}
	formattedUsers := ""

	for _, candidate := range candidatesKick {
		if len(formattedUsers) == 0 {
			formattedUsers = fmt.Sprintf("%v", candidate.User.Username)
		} else {
			formattedUsers = fmt.Sprintf("%v\n%v", formattedUsers, candidate.User.Username)
		}
	}

	if len(formattedUsers) != 0 {
		formattedUsers = fmt.Sprintf(":\n```\n%v\n```\n", formattedUsers)
	} else {
		formattedUsers = ".\n"
	}

	formattedMessage := fmt.Sprintf("*Verification Pruning Report*\n\n**%v** users kicked for failing to verify%v**%v** additional users reminded to verify.", len(candidatesKick), formattedUsers, len(warnedUsers))

	_, _ = s.ChannelMessageSend(*common.ModActionLogsThreadID, formattedMessage)
}
