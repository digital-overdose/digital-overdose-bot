package common

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

// Checks whether the user having instantiated the command has sufficient rights to do so.
func CheckHasPermissions(i *discordgo.InteractionCreate, s *discordgo.Session, permission int64) (bool, error) {
	if (i.Member.Permissions & permission) != permission {
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Only Staff members (Community Manager, Community Volunteer, Helper) may use this command.",
				Flags:   1 << 6,
			},
		})

		formattedMessage := fmt.Sprintf("[👁] User '%v' (ID: %v) unsuccessfully used the command '%v'.", FormatUsername(i.Member.User), i.Member.User.ID, i.ApplicationCommandData().Name)
		LogAndSend(formattedMessage, s, *StaffChannelID)

		return false, err
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Processing",
			Flags:   1 << 6,
		},
	})

	log.Printf("[👁] User '%v' (ID: %v) successfully used the command '%v'.", FormatUsername(i.Member.User), i.Member.User.ID, i.ApplicationCommandData().Name)

	return true, nil
}
