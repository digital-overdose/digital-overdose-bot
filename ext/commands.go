package ext

import (
	ext "atomicnicos.me/digital-overdose-bot/ext/commands"
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
	"purge-verification":   ext.PurgeVerification,
	"welcome":              ext.WelcomeUser,
	"test-current-feature": ext.TestCurrentFeature,
	//"warn-user":            ext.WarnUserTest,
	//"is-user-admin":         ext.IsUserAdmin,
	//"test-dm-requester":     ext.TestDMRequester,
}
