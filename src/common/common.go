package common

import (
	"flag"
	"fmt"
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
	ModActionLogsThreadID  = flag.String("mod-thread", "", "The channel where the bots' repetitive actions are published.")
	DebugChannelID         = flag.String("debug", "", "The channel to write debug to.")
	DebugChannelWarning    = false

	// Extensions
	HumanRoleID              = flag.String("human", "", "The role assigned to humans.")
	MemberRoleID             = flag.String("member", "", "The role assigned to members.")
	MainChannelID            = flag.String("main", "", "The main channel.")
	StaffChannelID           = flag.String("staff", "", "The staff channel.")
	UpgradeReleaseURL        = flag.String("upgrade", "", "The baseline path for bot upgrades.")
	MuteRoleID               = flag.String("mute", "", "The role assigned to members that are on timeout.")
	PrivateModLogsChannelID  = flag.String("private-mod", "", "The channel where all the mod events are logged to.")
	PrivateChatLogsChannelID = flag.String("private-chat", "", "The channel where all the chat events are logged to.")

	CURRENTLY_DEV = flag.Bool("dev", false, "Whether or not the bot is run from a dev environment.")
)

// A bypass function, to avoid collisions between a temporarily running dev environment and a permanently running server.
func ShouldExecutionBeSkippedIfDev(shouldSkipIfDev bool) bool {
	// The function should always run in prod.
	if !*CURRENTLY_DEV {
		return false
	} else {
		// Should only run if explicitly told that it can.
		return *CURRENTLY_DEV == shouldSkipIfDev
	}
}

// Attempts to load an environment file, and if it exists, overwrites any flags set through argv with it's contents.
// Checks whether any essential flags are missing.
func LoadEnvOrFlags() {
	env, err := godotenv.Read("./env/.env")

	if err != nil {
		log.Println("Error loading .env file, relying on flags.")
	} else {
		_ = godotenv.Load()

		tokens := []string{"GUILD", "TOKEN", "VERIFICATION_ROLE_ID", "VERIFICATION_CHANNEL_ID", "MOD_ACTION_CHANNEL_ID", "MOD_ACTION_THREAD_ID", "DEBUG_CHANNEL_ID"}
		references := []*string{GuildID, BotToken, VerificationRoleID, VerificationChannelID, ModActionLogsChannelID, ModActionLogsThreadID, DebugChannelID}

		// EXTENSIONS
		tokens = append(tokens, "HUMAN_ROLE_ID", "MEMBER_ROLE_ID", "MAIN_CHANNEL_ID", "STAFF_CHANNEL_ID", "UPGRADE_RELEASE_PATH", "MUTE_ROLE_ID", "PRIVATE_MOD_LOGS_CHANNEL_ID", "PRIVATE_CHAT_LOGS_CHANNEL_ID")
		references = append(references, HumanRoleID, MemberRoleID, MainChannelID, StaffChannelID, UpgradeReleaseURL, MuteRoleID, PrivateModLogsChannelID, PrivateChatLogsChannelID)

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

func SendMessage(s *discordgo.Session, channel string, message string) {
	if channel != "" {
		s.ChannelMessageSend(channel, message)
	} else {
		log.Printf("[❌] ERROR - Could not send message to channel because it is undefined.")
	}
}

func SendEmbed(s *discordgo.Session, channel string, message *discordgo.MessageEmbed) {
	if channel != "" {
		s.ChannelMessageSendEmbed(channel, message)
	} else {
		log.Printf("[❌] ERROR - Could not send message to channel because it is undefined.")
	}
}

func FormatUsername(u *discordgo.User) string {
	if u.Discriminator == "0" {
		return u.Username
	} else {
		return fmt.Sprintf("%s#%s", u.Username, u.Discriminator)
	}
}
