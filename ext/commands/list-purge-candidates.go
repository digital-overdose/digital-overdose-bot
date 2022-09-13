package ext

import (
	"log"
	"time"

	"atomicnicos.me/go-bot/common"
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
	)

	log.Printf("[+] Starting Member Scan...")

	for listing {
		members, err := s.GuildMembers(g.ID, lastID, 1000)
		if err != nil {
			log.Panicf("Unable to get Members of Guild %v: %v", *common.GuildID, err)
		}

		log.Printf("\tGot %v Members from Guild %v!", len(members), g.ID)

		for _, m := range members {
			for _, r := range m.Roles {
				if r == common.VerificationRoleID {
					candidates = append(candidates, *m)
				}
			}
		}

		if len(members) < 1000 {
			listing = false
		} else {
			lastID = members[len(members)-1].User.ID
		}
	}

	log.Printf("[+] Finished Member Scan...\n")

	log.Printf("[+] Got %v Deletion Candidates from Guild %v!\n", len(candidates), g.ID)

	log.Printf("[+] Scanning verification channel messages.")

	for messagesProcessed < messageLimit {
		messages, err := s.ChannelMessages(common.VerificationChannelID, 100, lastMessageID, "", "")
		if err != nil {
			log.Panicf("Unable to get Messages of Guild %v: %v", *common.GuildID, err)
		}

		messagesProcessed += len(messages)
		lastMessageID = messages[len(messages)-1].ID

		log.Printf("\tGot %v messages (%v/%v)!", len(messages), messagesProcessed, messageLimit)

		for _, msg := range messages {
			if _, ok := membersThatMessaged[msg.Author.ID]; !ok {
				membersThatMessaged[msg.Author.ID] = msg.Timestamp
			}
		}
	}

	log.Printf("[+] Got %v members having posted a message.", len(membersThatMessaged))

	log.Printf("[+] Starting Candidate Filtering...")

	now := time.Now()
	kickDate := now.Add(-7 * 24 * time.Hour)
	warnDate := now.Add(-5 * 24 * time.Hour)

	for _, c := range candidates {
		if _, ok := membersThatMessaged[c.User.ID]; ok {
			if membersThatMessaged[c.User.ID].Before(kickDate) {
				candidatesKick = append(candidatesKick, c)
			} else if membersThatMessaged[c.User.ID].Before(warnDate) {
				candidatesWarn = append(candidatesWarn, c)
			}
		} else if c.JoinedAt.Before(kickDate) {
			candidatesKick = append(candidatesKick, c)
		} else if c.JoinedAt.Before(warnDate) {
			candidatesWarn = append(candidatesWarn, c)
		}
	}
	log.Printf("[+] Finished Candidate Filtering...")

	for _, kick := range candidatesKick {
		log.Printf("User %v will be kicked.", kick.User.Username)
	}
	for _, warn := range candidatesWarn {
		log.Printf("User %v will be warned.", warn.User.Username)
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Pong",
		},
	})
}
