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

	logAndSend("[+] Starting Member Scan...", i.ChannelID, s)

	for listing {
		members, err := s.GuildMembers(g.ID, lastID, 1000)
		if err != nil {
			log.Panicf("Unable to get Members of Guild %v: %v", *common.GuildID, err)
		}

		logAndSend(fmt.Sprintf("\tGot %v Members from Guild %v!", len(members), g.ID), i.ChannelID, s)

		for _, m := range members {
			for _, r := range m.Roles {
				if r == common.VerificationRoleID {
					candidates = append(candidates, *m)
				}
			}
			if _, ok := timeLastSeen[m.User.ID]; !ok {
				timeLastSeen[m.User.ID] = m.JoinedAt
			}
		}

		if len(members) < 1000 {
			listing = false
		} else {
			lastID = members[len(members)-1].User.ID
		}
	}

	logAndSend("[+] Finished Member Scan...", i.ChannelID, s)

	logAndSend(fmt.Sprintf("[+] Got %v Deletion Candidates from Guild %v!\n", len(candidates), g.ID), i.ChannelID, s)

	logAndSend("[+] Scanning verification channel messages.", i.ChannelID, s)

	for messagesProcessed < messageLimit {
		messages, err := s.ChannelMessages(common.VerificationChannelID, 100, lastMessageID, "", "")
		if err != nil {
			log.Panicf("Unable to get Messages of Guild %v: %v", *common.GuildID, err)
		}

		messagesProcessed += len(messages)
		lastMessageID = messages[len(messages)-1].ID

		logAndSend(fmt.Sprintf("\tGot %v messages (%v/%v)!", len(messages), messagesProcessed, messageLimit), i.ChannelID, s)

		for _, msg := range messages {
			if _, ok := membersThatMessaged[msg.Author.ID]; !ok {
				membersThatMessaged[msg.Author.ID] = msg.Timestamp
			}
		}
	}

	for k, v := range membersThatMessaged {
		timeLastSeen[k] = v
	}

	logAndSend(fmt.Sprintf("[+] Got %v members having posted a message.", len(membersThatMessaged)), i.ChannelID, s)

	logAndSend("[+] Starting Candidate Filtering...", i.ChannelID, s)

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
	logAndSend("[+] Finished Candidate Filtering...", i.ChannelID, s)

	for _, kick := range candidatesKick {
		logAndSend(fmt.Sprintf("User %v will be kicked.", kick.User.Username), i.ChannelID, s)
		kickUser(kick.User.ID, kick.User.Username, i.ChannelID, s)
	}

	for _, warn := range candidatesWarn {
		logAndSend(fmt.Sprintf("User %v will be warned.", warn.User.Username), i.ChannelID, s)
	}
	warnUsers(candidatesWarnIDs, timeLastSeen, s)
	logAndSend("Done!", i.ChannelID, s)
}

func warnUsers(userIDs []string, timeLastSeen map[string]time.Time, s *discordgo.Session) (bool, error) {
	users := "`Not seen in server since`   |   User\n"
	for _, user := range userIDs {
		users += fmt.Sprintf("<t:%v:f>   -   %v\n", timeLastSeen[user].Unix(), fmt.Sprintf("<@%v>,", user))
	}

	formatted_msg := fmt.Sprintf(":warning: **Verification** :warning:\n\nThe following users have yet to verify!\n\n%v\nPlease verify! In order to verify:\n\t1) ✅ Please accept the rules verbally in the verification channel\n\t2) Please tell us a bit about yourself.\n\nFailure to verify will result in you being kicked in less than 2 days.", users)

	//_, err := s.ChannelMessageSend(common.VerificationChannelID, formatted_msg)
	_, err := s.ChannelMessageSend("1018835444173131836", formatted_msg)

	return err != nil, err
}

func dmUser(userID string, channelID string, s *discordgo.Session) (bool, error) {
	dmChannel, err := s.UserChannelCreate(userID)
	if err != nil {
		logAndSend(fmt.Sprintf("Could not DM user %v", userID), channelID, s)
		return err != nil, err
	}

	_, err = s.ChannelMessageSend(dmChannel.ID, "*This message is sent to you by the Digital Overdose Server's Management Bot.*\n\nHello!\nWe wish to inform you that you have been automatically pruned from the server due to not completing the steps required for verification.\n\nIf you wish to rejoin (and this time complete verification, it only takes 2 minutes), you may use this link: https://discord.gg/jaNuDB95Yd\n\nHave a nice day!")

	return err != nil, err
}

func kickUser(userID string, username string, channelID string, s *discordgo.Session) (bool, error) {
	_, _ = dmUser(userID, channelID, s)
	err := s.GuildMemberDeleteWithReason(*common.GuildID, userID, fmt.Sprintf("User %v failed to verify in time!", username))
	if err != nil {
		return err != nil, err
	}
	return err != nil, err
}

func logAndSend(message string, channelID string, s *discordgo.Session) {
	log.Print(message)

	_, _ = s.ChannelMessageSend(channelID, message)
}
