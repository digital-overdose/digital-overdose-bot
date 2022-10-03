package common

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

// Helper function to help an embed without all of the annoying fields.
// One can simply provide a selection of fields and a footer on an "as needed" basis.
func BuildEmbed(
	title string,
	description string,
	fields []*discordgo.MessageEmbedField,
	footer *discordgo.MessageEmbedFooter) *discordgo.MessageEmbed {
	embed := &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Type:        discordgo.EmbedTypeRich,
		Title:       title,
		Description: description,
		Thumbnail:   &discordgo.MessageEmbedThumbnail{},
		Timestamp:   time.Now().Format(time.RFC3339),
	}
	if fields != nil {
		embed.Fields = fields
	}
	if footer != nil {
		embed.Footer = footer
	}

	return embed
}
