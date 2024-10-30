package extensions

import (
	commands "atomicmaya.me/digital-overdose-bot/src/extensions/commands"
	commands_management "atomicmaya.me/digital-overdose-bot/src/extensions/commands/management"
	commands_moderation "atomicmaya.me/digital-overdose-bot/src/extensions/commands/moderation"
	commands_moderation_ban "atomicmaya.me/digital-overdose-bot/src/extensions/commands/moderation/ban"
	commands_moderation_mute "atomicmaya.me/digital-overdose-bot/src/extensions/commands/moderation/mute"
	commands_moderation_warn "atomicmaya.me/digital-overdose-bot/src/extensions/commands/moderation/warn"
	"github.com/bwmarrin/discordgo"
)

// List of command details for the user-facing side of the bot.
var Commands = []*discordgo.ApplicationCommand{
	{
		Name:        "purge-verification",
		Description: "Purges all members that have failed to verify in time (Scheduled).",
	},
	{
		Name:        "status",
		Description: "Provides status information about the bot.",
	},
	{
		Name:        "welcome",
		Description: "Assigns roles, and welcomes the user to the server!",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "user",
				Description: "User to be welcomed",
				Required:    true,
			},
		},
	},
	{
		Name:        "lookup",
		Description: "Provides any information stored on the system about the specified user.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "user_id",
				Description: "(IF NOT MEMBER) ID of the User to lookup",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "user",
				Description: "(IF MEMBER) User to lookup",
				Required:    false,
			},
		},
	},
	{
		Name:        "stats",
		Description: "Gets server activity stats",
	},
	{
		Name:        "test-current-feature",
		Description: "Tests whatever feature I'm currently trying out.",
	},
	{
		Name:        "upgrade",
		Description: "Upgrades the bot to a specified release.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "version",
				Description: "`x.x.x`: The version string of the release to be downloaded.`",
				Required:    true,
			},
		},
	},
	{
		Name:        "warn",
		Description: "Warns a user.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "user",
				Description: "The user to be warned.",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "reason",
				Description: "The reason for the warn.",
				Required:    true,
			},
		},
	},
	{
		Name:        "unwarn",
		Description: "Unwarns a user.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "user",
				Description: "The user to have their last warning removed.",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "which",
				Description: "Which of the warns should be removed.",
				Required:    false,
			},
		},
	},
	{
		Name:        "list-warns",
		Description: "Lists the warns a user has gotten.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "user",
				Description: "The user to list the warnings of.",
				Required:    true,
			},
		},
	},
	{
		Name:        "mute",
		Description: "Mutes a user.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "user",
				Description: "The user to be muted.",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "duration",
				Description: "DEFAULT=-1 // The duration of the mute (e.g. '1.5h', '2h45m', valid time units are 'm', 'h', 'd'.)",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "reason",
				Description: "The reason for the mute.",
				Required:    false,
			},
		},
	},
	{
		Name:        "unmute",
		Description: "Unmutes a user (manually).",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "user",
				Description: "The user to be unmuted.",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "reason",
				Description: "The reason for the unmute.",
				Required:    false,
			},
		},
	},
	{
		Name:        "ban",
		Description: "Bans a user.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "user",
				Description: "The user to be banned.",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "reason",
				Description: "The reason for the ban.",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "duration",
				Description: "DEFAULT=-1 // The duration of the ban. (e.g. '1.5h', '2h45m', valid time units are 'm', 'h', 'd'.)",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionNumber,
				Name:        "delete_days",
				Description: "DEFAULT=0 // Number of days of messages to delete",
				Required:    false,
			},
		},
	},
	{
		Name:        "unban",
		Description: "Unbans a user",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "user_id",
				Description: "The ID of the User to be unbanned.",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "reason",
				Description: "The reason for the unban.",
				Required:    true,
			},
		},
	},
	/*{
		Name:        "is-user-admin",
		Description: "Checks whether the user has the Manage Server permission",
	},
	{
		Name:        "test-dm-requester",
		Description: "Sends a DM to the person requesting the command.",
	},*/
	/*{
		Name:        "warn-user",
		Description: "Warns the person requesting the command.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "user",
				Description: "User to be warned",
				Required:    false,
			},
		},
	},*/
}

// Command to bot function map.
var CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"purge-verification":   commands_management.PurgeVerification,
	"welcome":              commands_moderation.WelcomeUser,
	"lookup":               commands_management.LookupMember,
	"stats":                commands_management.ServerStats,
	"bot-status":           commands_management.BotStatus,
	"bot-upgrade":          commands_management.BotUpgrade,
	"mute":                 commands_moderation_mute.Mute,
	"unmute":               commands_moderation_mute.UnmuteManual,
	"warn":                 commands_moderation_warn.Warn,
	"unwarn":               commands_moderation_warn.Unwarn,
	"ban":                  commands_moderation_ban.Ban,
	"unban":                commands_moderation_ban.Unban,
	"test-current-feature": commands.TestCurrentFeature,
	// "list-warns":           commands_moderation_warn.ListWarns, // Replaced by "lookup"
	// "list-mutes":           commands_moderation_mute.ListMutes, // Replaced by "lookup"
	// "reason-ban":           commands_moderation_ban.ReasonBan, // Replaced by "lookup"
}
