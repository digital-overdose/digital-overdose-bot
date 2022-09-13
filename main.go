package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"atomicnicos.me/go-bot/common"
	"atomicnicos.me/go-bot/ext"
	"github.com/bwmarrin/discordgo"
)

var s *discordgo.Session

func init() { flag.Parse() }

func init() {
	var err error

	s, err = discordgo.New("Bot " + *common.BotToken)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}

	s.Identify.Intents = discordgo.IntentsAllWithoutPrivileged | discordgo.IntentsGuildMembers
}

var (
	dmPermission                   = false
	defaultMemberPermissions int64 = discordgo.PermissionManageServer
	commands                       = ext.Commands
	commandHandlers                = ext.CommandHandlers
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
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, *common.GuildID, v)
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
