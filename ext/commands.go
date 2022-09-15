package ext

import (
	ext "atomicnicos.me/digital-overdose-bot/ext/commands"
	"github.com/bwmarrin/discordgo"
)

// List of command details for the user-facing side of the bot.
var Commands = []*discordgo.ApplicationCommand{
	{
		Name:        "list-purge-candidates",
		Description: "Lists all the people who would be affected by a purge.",
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
	"list-purge-candidates": ext.ListPurgeCandidates,
	"test-current-feature":  ext.TestCurrentFeature,
	//"is-user-admin":         ext.IsUserAdmin,
	//"test-dm-requester":     ext.TestDMRequester,
	"warn-user": ext.WarnUserTest,
}
