package extensions

import (
	commands "atomicnicos.me/digital-overdose-bot/extensions/commands"
	commands_management "atomicnicos.me/digital-overdose-bot/extensions/commands/management"
	commands_moderation "atomicnicos.me/digital-overdose-bot/extensions/commands/moderation"
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
				Description: "`vx.x.x`: The version string of the release to be downloaded.`",
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
	"warn":                 commands_moderation.Warn,
	"unwarn":               commands_moderation.Unwarn,
	//"warn-user":            ext.WarnUserTest,
	//"is-user-admin":         ext.IsUserAdmin,
	//"test-dm-requester":     ext.TestDMRequester,
}
