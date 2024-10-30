/*
This package contains the core code for the Digital Overdose Discord Server's Management Bot.

Current features include:
- Manual Purge
*/
package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"atomicmaya.me/digital-overdose-bot/src/common"
	cron "atomicmaya.me/digital-overdose-bot/src/cron"
	database_utils "atomicmaya.me/digital-overdose-bot/src/db"
	"atomicmaya.me/digital-overdose-bot/src/extensions"
	"atomicmaya.me/digital-overdose-bot/src/handler"
	"github.com/bwmarrin/discordgo"
	"github.com/go-co-op/gocron"
)

// The Discord session, used for state management.
var s *discordgo.Session

// Parses the flags provided in `argv`, then loads any overwrites from the .env file.
func init() {
	flag.Parse()
	common.LoadEnvOrFlags()
}

// Starts up the bot.
func init() {
	var err error

	s, err = discordgo.New("Bot " + *common.BotToken)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}

	s.Identify.Intents = discordgo.IntentsAllWithoutPrivileged | discordgo.IntentsGuildMembers
}

func init() {
	common.InitializeLogging()
}

func init() {
	err := errors.New("")
	database_utils.Database, err = database_utils.InitializeDatabase()
	if err != nil {
		log.Printf("DB INIT failed. ERR: %v", err)
		os.Exit(125)
	}
}

func init() {
	s.AddHandler(handler.OnReady)
	s.AddHandler(handler.OnInteractionCreate)
	s.AddHandler(handler.OnMessage)
	s.AddHandler(handler.OnJoin)
	s.AddHandler(handler.OnLeave)
}

// Initializes cron engine, and subsequently registers scheduled functions
func init() {
	schedulers := make([]*gocron.Scheduler, len(cron.CronJobs))

	for i := 0; i < len(cron.CronJobs); i++ {
		schedulers[i] = gocron.NewScheduler(time.UTC)
		job := cron.CronJobs[i]

		log.Printf("[--] Registered job '%v': '%v'", job.Name, job.CronString)
		schedulers[i].Cron(job.CronString).Do(func() {
			log.Printf("[+] Executing cron job '%v': '%v'", job.Name, job.CronString)
			job.Job(s, nil)
		})
		schedulers[i].StartAsync()
	}

	log.Print("[✓] Registering Jobs")

	// Start the cron scheduler in another thread.
	log.Print("[✓] cron handler Started")
}

// Entry point, loads system and performs clean-up
func main() {
	err := s.Open()

	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	log.Println("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(extensions.Commands))

	// Registers all of the commands in the designated server.
	for i, v := range extensions.Commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, *common.GuildID, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
		log.Printf("Added command '%v'\n", v.Name)
	}
	defer s.Close()

	// CTRL+C Signal Handler
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	log.Println("Press Ctrl+C to exit")

	common.LogAndSend(fmt.Sprintf(":robot::rotating_light: is currently running on version `%v%v", common.VERSION, func() string {
		if *common.CURRENTLY_DEV {
			return " (DEV VERSION)`"
		} else {
			return "`"
		}
	}()), s)

	<-stop

	// Unregisters the commands in the designated server.
	if *common.RemoveCommands {
		log.Println("Removing commands...")
		for _, v := range registeredCommands {
			s.ApplicationCommandDelete(s.State.User.ID, *common.GuildID, v.ID)
			if err != nil {
				log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
			}
			log.Printf("Removed command '%v'\n", v.Name)
		}
	}

	log.Printf("Gracefully shutting down.")
}
