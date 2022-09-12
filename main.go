package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	GuildID               = flag.String("guild", "", "Test guild ID. If not passed - bot registers commands globally")
	BotToken              = flag.String("token", "", "Bot access token")
	RemoveCommands        = flag.Bool("rmcmd", true, "Remove all commands after shutdowning")
	VerificationChannelID = "687238387463094317"
	VerificationRoleID    = "687228151096541185"
)
var s *discordgo.Session

func init() { flag.Parse() }

func init() {
	var err error

	s, err = discordgo.New("Bot " + *BotToken)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}

	s.Identify.Intents = discordgo.IntentsAllWithoutPrivileged | discordgo.IntentsGuildMembers
}

var (
	dmPermission                   = false
	defaultMemberPermissions int64 = discordgo.PermissionManageServer

	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "list-purge-candidates",
			Description: "Lists all the people who would be affected by a purge.",
		},
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"list-purge-candidates": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			g, err := s.State.Guild(*GuildID)
			if err != nil {
				log.Panicf("Unable to get Guild %v: %v", *GuildID, err)
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
					log.Panicf("Unable to get Members of Guild %v: %v", *GuildID, err)
				}

				log.Printf("\tGot %v Members from Guild %v!", len(members), g.ID)

				for _, m := range members {
					for _, r := range m.Roles {
						if r == VerificationRoleID {
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
				messages, err := s.ChannelMessages(VerificationChannelID, 100, lastMessageID, "", "")
				if err != nil {
					log.Panicf("Unable to get Messages of Guild %v: %v", *GuildID, err)
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
		},
	}
)

func init() {
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}

func main() {

	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	err := s.Open()

	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	log.Println("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))

	for i, v := range commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, *GuildID, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	defer s.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	log.Printf("Gracefully shutting down.")
}
