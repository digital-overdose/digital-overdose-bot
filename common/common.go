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
)

// Checks whether the user having instantiated the command has sufficient rights to do so.
func HasPermissions(i *discordgo.InteractionCreate, s *discordgo.Session, permission int64) (bool, error) {
	if (i.Member.Permissions & permission) != permission {
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "You don't have permission to use this command!",
				Flags:   1 << 6,
			},
		})
		return false, err
	}
	return true, nil
}

// Writes the message to the application log and the debug channel, if it is set.
func LogAndSend(message string, s *discordgo.Session) {
	log.Print(message)

	if *DebugChannelID != "" {
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
		if env["GUILD"] != "" {
			*GuildID = env["GUILD"]
		}
		if env["TOKEN"] != "" {
			*BotToken = env["TOKEN"]
		}
		if env["VERIFICATION_ROLE_ID"] != "" {
			*VerificationRoleID = env["VERIFICATION_ROLE_ID"]
		}
		if env["VERIFICATION_CHANNEL_ID"] != "" {
			*VerificationChannelID = env["VERIFICATION_CHANNEL_ID"]
		}
		if env["MOD_ACTION_CHANNEL_ID"] != "" {
			*ModActionLogsChannelID = env["MOD_ACTION_CHANNEL_ID"]
		}
		if env["DEBUG_CHANNEL_ID"] != "" {
			*DebugChannelID = env["DEBUG_CHANNEL_ID"]
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
