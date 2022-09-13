package ext

import (
	ext "atomicnicos.me/digital-overdose-bot/ext/commands"
	"github.com/bwmarrin/discordgo"
)

var Commands = []*discordgo.ApplicationCommand{
	{
		Name:        "list-purge-candidates",
		Description: "Lists all the people who would be affected by a purge.",
	},
	/*{
		Name:        "is-user-admin",
		Description: "Checks whether the user has the Manage Server permission",
	},*/
	/*{
		Name:        "test-dm-requester",
		Description: "Sends a DM to the person requesting the command.",
	},*/
}

var CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"list-purge-candidates": ext.ListPurgeCandidates,
	//"is-user-admin":         ext.IsUserAdmin,
	//"test-dm-requester":     ext.TestDMRequester,
}
