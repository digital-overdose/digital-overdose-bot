package common

import (
	"flag"

	"github.com/bwmarrin/discordgo"
)

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

var (
	GuildID               = flag.String("guild", "", "Test guild ID. If not passed - bot registers commands globally")
	BotToken              = flag.String("token", "", "Bot access token")
	RemoveCommands        = flag.Bool("rmcmd", true, "Remove all commands after shutdowning")
	VerificationChannelID = "687238387463094317"
	VerificationRoleID    = "687228151096541185"
)
