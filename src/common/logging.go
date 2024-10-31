package common

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	mw             io.Writer
	logFile        *os.File
	canWriteToFile bool
)

// Creates the logging infrastructure for the program.
// Initializes logging to a ./log/<datetime>.log file (if it can).
func InitializeLogging() {
	log.SetFlags(log.LUTC)
	log.SetOutput(os.Stdout)

	// Checks whether or not the program has sufficient rights to create files.
	canWriteToFile = true
	if _, err := os.Stat("./log/"); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir("./log/", os.ModePerm)
		if err != nil {
			canWriteToFile = false
			log.Println(err)
		}
	}

	// Creates a MultiWriter (basically an output bifurcator) and assigns a series of output streams depending on write rights.
	mw = io.MultiWriter()
	if canWriteToFile {
		logFile, err := os.OpenFile(fmt.Sprintf("log/%v-bot.log", time.Now().UTC().Format("2006-01-02-15-04-05Z")), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("Error opening file: %v", err)
		}
		mw = io.MultiWriter(os.Stdout, logFile)
	} else {
		mw = io.MultiWriter(os.Stdout)
	}
	log.SetOutput(mw)
}

// Allows for log pagination, ie. the closing and opening of a new log file after a specific action.
// The objective is to reduce the amount of data a running process can push to a single file, to simplify parsing later on.
func PaginateLog(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if !canWriteToFile {
		return
	}

	LogToServer(":robot: :rotating_light: `log-pagination` triggered by cron.", s)

	log.SetOutput(os.Stdout)
	err := logFile.Close()
	if err != nil {
		log.Printf("Error closing file: %v", err)
	}

	logFile, err = os.OpenFile(fmt.Sprintf("log/%v-bot.log", time.Now().UTC().Format("2006-01-02-15-04-05Z")), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	mw = io.MultiWriter( /* os.Stdout, */ logFile)
	log.SetOutput(mw)
}

func Log(format string, a ...any) string {
	v := []any{time.Now().UTC().Format("2006-01-02 15:04:05Z")}
	v = append(v, a...)
	log.Printf("[%s] "+format, v...)
	// fmt.Printf("[%s] "+format+"\n", v...)
	return fmt.Sprintf("[%s] "+format, v...)
}

// Writes the message to the debug channel, if it is set.
func LogToServer(message string, s *discordgo.Session, nonDefaultChannelID ...string) string {
	if len(nonDefaultChannelID) > 0 && len(nonDefaultChannelID[0]) > 0 {
		msg, _ := s.ChannelMessageSend(nonDefaultChannelID[0], message)
		return msg.ID
	} else if len(nonDefaultChannelID) > 0 && len(nonDefaultChannelID[0]) == 0 {
		return ""
	} else if *DebugChannelID != "" {
		msg, _ := s.ChannelMessageSend(*DebugChannelID, message)
		return msg.ID
	} else if !DebugChannelWarning {
		Log("DEBUG_CHANNEL_ID / --debug was not defined, skipping all debug message forwarding.")
	}

	return ""
}
