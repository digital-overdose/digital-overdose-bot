/*
This package contains the core code for the Digital Overdose Discord Server's Management Bot.

Current features include:
- Manual Purge
*/
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"time"

	"atomicnicos.me/digital-overdose-bot/common"
	ext "atomicnicos.me/digital-overdose-bot/ext"
	"github.com/bwmarrin/discordgo"
)

var s *discordgo.Session

func init() {
	flag.Parse()
	common.LoadEnvOrFlags()
}

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
	f, err := os.OpenFile(fmt.Sprintf("log/%v-bot.log", time.Now().Format("2006-01-02-15-04-05")), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	mw := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mw)
}

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
