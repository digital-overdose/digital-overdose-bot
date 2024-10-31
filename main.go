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
		common.Log("DB INIT failed. ERR: %v", err)
		os.Exit(125)
	}
}

func init() {
	s.AddHandler(handler.OnReady)
	s.AddHandler(handler.OnInteractionCreate)
	s.AddHandler(handler.OnMessage)
	s.AddHandler(handler.OnMessageUpdate)
	s.AddHandler(handler.OnMessageDelete)
	s.AddHandler(handler.OnJoin)
	s.AddHandler(handler.OnLeave)
}

// Initializes cron engine, and subsequently registers scheduled functions
func init() {
	schedulers := make([]*gocron.Scheduler, len(cron.CronJobs))

	for i := 0; i < len(cron.CronJobs); i++ {
		schedulers[i] = gocron.NewScheduler(time.UTC)
		job := cron.CronJobs[i]

		common.Log("[--] Registered job '%v': '%v'", job.Name, job.CronString)
		schedulers[i].Cron(job.CronString).Do(func() {
			common.Log("[+] Executing cron job '%v': '%v'", job.Name, job.CronString)
			job.Job(s, nil)
		})
		schedulers[i].StartAsync()
	}

	common.Log("[✓] Registering Jobs")

	// Start the cron scheduler in another thread.
	common.Log("[✓] cron handler Started")
}

// Entry point, loads system and performs clean-up
func main() {
	err := s.Open()

	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	common.Log("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(extensions.Commands))

	// Registers all of the commands in the designated server.
	for i, v := range extensions.Commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, *common.GuildID, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
		common.Log("Added command '%v'\n", v.Name)
	}
	defer s.Close()

	// CTRL+C Signal Handler
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	common.Log("Press Ctrl+C to exit")

	common.LogToServer(fmt.Sprintf(":robot::rotating_light: is currently running on version `%v%v", common.VERSION, func() string {
		if *common.CURRENTLY_DEV {
			return " (DEV VERSION)`"
		} else {
			return "`"
		}
	}()), s)

	<-stop

	_, err = (*database_utils.Database).Methods.InsertOpsEvent.Exec(database_utils.SYSTEM_STOP, time.Now(), "Graceful exit.")

	// Unregisters the commands in the designated server.
	if *common.RemoveCommands {
		common.Log("Removing commands...")
		for _, v := range registeredCommands {
			s.ApplicationCommandDelete(s.State.User.ID, *common.GuildID, v.ID)
			if err != nil {
				log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
			}
			common.Log("Removed command '%v'\n", v.Name)
		}
	}

	common.Log("Gracefully shutting down.")
}
