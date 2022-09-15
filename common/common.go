package common

import (
	"flag"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	GuildID                = flag.String("guild", "", "The ID of the Discord Server")
	BotToken               = flag.String("token", "", "Bot access token")
	RemoveCommands         = flag.Bool("rmcmd", true, "Remove all commands after shutdowning")
	VerificationRoleID     = flag.String("role", "", "The role that targets the role given to new people in the server.")
	VerificationChannelID  = flag.String("wall", "", "The channel where new members aim to verify.")
	ModActionLogsChannelID = flag.String("mod", "", "The channel where the bots actions are then published.")
	DebugChannelID         = flag.String("debug", "", "The channel to write debug to.")
	DebugChannelWarning    = false

	// Extensions
	HumanRoleID    = flag.String("human", "", "The role assigned to humans.")
	MemberRoleID   = flag.String("member", "", "The role assigned to members.")
	MainChannelID  = flag.String("main", "", "The main channel.")
	StaffChannelID = flag.String("staff", "", "The staff channel.")
)

// Writes the message to the application log and the debug channel, if it is set.
func LogAndSend(message string, s *discordgo.Session, optional ...string) {
	log.Print(message)

	if len(optional) > 0 && len(optional[0]) > 0 {
		_, _ = s.ChannelMessageSend(optional[0], message)
	} else if *DebugChannelID != "" {
		_, _ = s.ChannelMessageSend(*DebugChannelID, message)
	} else if !DebugChannelWarning {
		log.Print("DEBUG_CHANNEL_ID / --debug was not defined, skipping all debug message forwarding.")
	}
}

// Attempts to load an environment file, and if it exists, overwrites any flags set through argv with it's contents.
// Checks whether any essential flags are missing.
func LoadEnvOrFlags() {
	env, err := godotenv.Read()

	if err != nil {
		log.Println("Error loading .env file, relying on flags.")
	} else {
		_ = godotenv.Load()

		tokens := []string{"GUILD", "TOKEN", "VERIFICATION_ROLE_ID", "VERIFICATION_CHANNEL_ID", "MOD_ACTION_CHANNEL_ID", "DEBUG_CHANNEL_ID"}
		references := []*string{GuildID, BotToken, VerificationRoleID, VerificationChannelID, ModActionLogsChannelID, DebugChannelID}

		// EXTENSIONS
		tokens = append(tokens, "HUMAN_ROLE_ID", "MEMBER_ROLE_ID", "MAIN_CHANNEL_ID", "STAFF_CHANNEL_ID")
		references = append(references, HumanRoleID, MemberRoleID, MainChannelID, StaffChannelID)

		if len(tokens) != len(references) {
			log.Fatalf("Mismatched Environment flags.")
		}

		for i := 0; i < len(tokens); i++ {
			if env[tokens[i]] != "" {
				*references[i] = env[tokens[i]]
			}
		}
	}

	missing := ""
	if *GuildID == "" {
		missing += "GUILD_ID / --guild, "
	}
	if *BotToken == "" {
		missing += "BOT_TOKEN / --token, "
	}
	if *VerificationRoleID == "" {
		missing += "VERIFICATION_ROLE_ID / --role, "
	}
	if *VerificationChannelID == "" {
		missing += "VERIFICATION_CHANNEL_ID / --wall, "
	}

	if len(missing) > 0 {
		log.Fatalf("One or more parameters missing: %v", missing[0:len(missing)-2])
	}
}
