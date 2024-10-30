package handler

import (
	"fmt"
	"log"
	"time"

	"atomicmaya.me/digital-overdose-bot/src/common"
	database_utils "atomicmaya.me/digital-overdose-bot/src/db"
	"github.com/bwmarrin/discordgo"
)

func OnMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	log.Println("ONMESSAGE")
	channel, err := s.Channel(m.ChannelID)
	if err != nil {
		log.Printf("ERROR IN ONMESSAGE: %v", err)
	}

	_, err = (*database_utils.Database).Methods.InsertUserEvent.Exec(
		database_utils.USER_MESSAGE_WRITE,
		time.Now(),
		m.Author.ID, common.FormatUsername(m.Author),
		channel.ID, fmt.Sprintf("#%v (<#%v>)", channel.Name, channel.ID),
		fmt.Sprintf("ID %v\nContent (%d): '%v'", m.ID, len(m.Content), common.EncodeMessage(m.Content, 100)),
	)
	if err != nil {
		log.Printf("ERROR IN ONMESSAGE: %v", err)
	}
}

func OnMessageUpdate(s *discordgo.Session, m *discordgo.MessageUpdate) {
	log.Println("ONMESSAGEUPDATE")
	channel, err := s.Channel(m.ChannelID)
	if err != nil {
		log.Printf("ERROR IN ONMESSAGEUPDATE: %v", err)
	}

	_, err = (*database_utils.Database).Methods.InsertUserEvent.Exec(
		database_utils.USER_MESSAGE_UPDATE,
		time.Now(),
		m.Author.ID, common.FormatUsername(m.Author),
		channel.ID, fmt.Sprintf("#%v (<#%v>)", channel.Name, channel.ID),
		fmt.Sprintf("ID %v\nContent (%d): '%v'", m.ID, len(m.Content), common.EncodeMessage(m.Content, 100)),
	)
	if err != nil {
		log.Printf("ERROR IN ONMESSAGEUPDATE: %v", err)
	}
}
