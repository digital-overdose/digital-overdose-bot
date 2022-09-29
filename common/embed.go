package common

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

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
