package ext

import (
	"fmt"
	"log"
	"time"

	"atomicnicos.me/digital-overdose-bot/common"

	"github.com/bwmarrin/discordgo"
)

func ListPurgeCandidates(s *discordgo.Session, i *discordgo.InteractionCreate) {
	ok, err := common.HasPermissions(i, s, discordgo.PermissionViewAuditLogs|discordgo.PermissionManageRoles)
	if err != nil {
		log.Println("Error checking permissions.")
	}

	if !ok {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "You're not STAFF",
			},
		})
		return
	} else {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Processing",
			},
		})
	}

	g, err := s.State.Guild(*common.GuildID)
	if err != nil {
		log.Panicf("Unable to get Guild %v: %v", *common.GuildID, err)
	}

	var (
		listing                                  = true
		lastID                                   = ""
		messagesProcessed                        = 0
		lastMessageID                            = ""
		messageLimit                             = 1000
		candidates          []discordgo.Member   = []discordgo.Member{}
		candidatesKick      []discordgo.Member   = []discordgo.Member{}
		candidatesWarn      []discordgo.Member   = []discordgo.Member{}
		membersThatMessaged map[string]time.Time = map[string]time.Time{}
		timeLastSeen        map[string]time.Time = map[string]time.Time{}
		candidatesWarnIDs   []string             = []string{}
	)

	common.LogAndSend("[+] Starting Member Scan...", s)

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

	for messagesProcessed < messageLimit {
		messages, err := s.ChannelMessages(*common.VerificationChannelID, 100, lastMessageID, "", "")
		if err != nil {
			log.Panicf("Unable to get Messages of Guild %v: %v", *common.GuildID, err)
		}

		messagesProcessed += len(messages)
		lastMessageID = messages[len(messages)-1].ID

		common.LogAndSend(fmt.Sprintf("\tGot %v messages (%v/%v)!", len(messages), messagesProcessed, messageLimit), s)

		for _, msg := range messages {
			if _, ok := membersThatMessaged[msg.Author.ID]; !ok {
				membersThatMessaged[msg.Author.ID] = msg.Timestamp
			}
		}
	}

	for k, v := range membersThatMessaged {
		timeLastSeen[k] = v
	}

	common.LogAndSend(fmt.Sprintf("[+] Got %v members having posted a message.", len(membersThatMessaged)), s)

	common.LogAndSend("[∨] Starting Candidate Filtering...", s)

	now := time.Now()
	kickDate := now.Add(-7 * 24 * time.Hour)
	warnDate := now.Add(-5 * 24 * time.Hour)

	for _, c := range candidates {
		if timeLastSeen[c.User.ID].Before(kickDate) {
			candidatesKick = append(candidatesKick, c)
		} else if timeLastSeen[c.User.ID].Before(warnDate) {
			candidatesWarn = append(candidatesWarn, c)
			candidatesWarnIDs = append(candidatesWarnIDs, c.User.ID)
		}
	}
	common.LogAndSend("[+] Finished Candidate Filtering...", s)

	for _, kick := range candidatesKick {
		common.LogAndSend(fmt.Sprintf("[----] User %v will be kicked.", kick.User.Username), s)
		kickUser(kick.User.ID, kick.User.Username, s)
	}

	for _, warn := range candidatesWarn {
		common.LogAndSend(fmt.Sprintf("[----] User %v will be warned.", warn.User.Username), s)
	}
	warnUsers(candidatesWarnIDs, timeLastSeen, s)

	sendModActionLogsMessage(candidatesKick, candidatesWarn, s)
	common.LogAndSend("[✓] Done!", s)
}

func warnUsers(userIDs []string, timeLastSeen map[string]time.Time, s *discordgo.Session) (bool, error) {
	users := "`Not seen in server since`   |   User\n"
	for _, user := range userIDs {
		users += fmt.Sprintf("<t:%v:f>   -   %v\n", timeLastSeen[user].Unix(), fmt.Sprintf("<@%v>,", user))
	}

	formatted_msg := fmt.Sprintf(":warning: **Verification** :warning:\n\nThe following users have yet to verify!\n\n%v\nPlease verify! In order to verify:\n\t1) ✅ Please accept the rules by posting your acceptance here in the verification channel\n\t2) Please tell us a bit about yourself.\n\nFailure to verify will result in you being kicked in less than 2 days.", users)

	_, err := s.ChannelMessageSend(*common.VerificationChannelID, formatted_msg)

	return err != nil, err
}

func dmUser(userID string, s *discordgo.Session) (bool, error) {
	dmChannel, err := s.UserChannelCreate(userID)
	if err != nil {
		common.LogAndSend(fmt.Sprintf("Could not DM user %v", userID), s)
		return err != nil, err
	}

	_, err = s.ChannelMessageSend(dmChannel.ID, "*This message is sent to you by the Digital Overdose Server's Management Bot.*\n\nHello!\nWe wish to inform you that you have been automatically pruned from the server due to not completing the steps required for verification.\n\nIf you wish to rejoin (and this time complete verification, it only takes 2 minutes), you may use this link: https://discord.gg/jaNuDB95Yd\n\nHave a nice day!")

	return err != nil, err
}

func kickUser(userID string, username string, s *discordgo.Session) (bool, error) {
	_, _ = dmUser(userID, s)
	err := s.GuildMemberDeleteWithReason(*common.GuildID, userID, fmt.Sprintf("User %v failed to verify in time!", username))
	if err != nil {
		return err != nil, err
	}
	return err != nil, err
}

func sendModActionLogsMessage(kickedUsers []discordgo.Member, warnedUsers []discordgo.Member, s *discordgo.Session) {
	if len(*common.ModActionLogsChannelID) == 0 {
		log.Printf("Channel ModActionLogs not set, skipping message.")
		return
	}
	kickedUsernames := ""

	for _, u := range kickedUsers {
		if len(kickedUsernames) == 0 {
			kickedUsernames = fmt.Sprintf("%v", u.User.Username)
		} else {
			kickedUsernames = fmt.Sprintf("%v,\n%v", kickedUsernames, u.User.Username)
		}
	}

	formatted_msg := fmt.Sprintf("*Verification Pruning Report*\n\n**%v** users kicked for failing to verify:\n```%v```\n**%v** additional users reminded to verify.", len(kickedUsers), kickedUsernames, len(warnedUsers))

	_, _ = s.ChannelMessageSend(*common.ModActionLogsChannelID, formatted_msg)
}
