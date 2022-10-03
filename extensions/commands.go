package extensions

import (
	commands "atomicnicos.me/digital-overdose-bot/extensions/commands"
	commands_management "atomicnicos.me/digital-overdose-bot/extensions/commands/management"
	commands_moderation "atomicnicos.me/digital-overdose-bot/extensions/commands/moderation"
	commands_moderation_ban "atomicnicos.me/digital-overdose-bot/extensions/commands/moderation/ban"
	commands_moderation_mute "atomicnicos.me/digital-overdose-bot/extensions/commands/moderation/mute"
	commands_moderation_warn "atomicnicos.me/digital-overdose-bot/extensions/commands/moderation/warn"
	"github.com/bwmarrin/discordgo"
)

// List of command details for the user-facing side of the bot.
var Commands = []*discordgo.ApplicationCommand{
	{
		Name:        "purge-verification",
		Description: "Purges all members that have failed to verify in time (Scheduled).",
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
				Description: "The duration of the mute. Is a duration string: '1.5h' or '2h45m'. Valid time units are 'm', 'h'.",
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
		},
	},
	{
		Name:        "list-mutes",
		Description: "Lists the mutes a user has gotten.",
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
		},
	},
	{
		Name:        "unban",
		Description: "Unbans a user",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "user",
				Description: "The user to be unbanned.",
				Required:    true,
			},
		},
	},
	{
		Name:        "reason-ban",
		Description: "Lists the reason that a user was banned for.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "user_id",
				Description: "The ID of the user to get the ban warning for.",
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
	"test-current-feature": commands.TestCurrentFeature,
	"upgrade":              commands_management.UpgradeBot,
	"warn":                 commands_moderation_warn.Warn,
	"unwarn":               commands_moderation_warn.Unwarn,
	"list-warns":           commands_moderation_warn.ListWarns,
	"ban":                  commands_moderation_ban.Ban,
	"unban":                commands_moderation_ban.Unban,
	"reason-ban":           commands_moderation_ban.ReasonBan,
	"mute":                 commands_moderation_mute.Mute,
	"unmute":               commands_moderation_mute.UnmuteManual,
	"list-mutes":           commands_moderation_mute.ListMutes,
}
