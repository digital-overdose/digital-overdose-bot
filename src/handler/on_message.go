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
	if *common.CURRENTLY_DEV {
		//log.Printf("New message by %v (%v) in %v : %v\n> %v\n", m.Author.String(), m.Author.ID, m.ChannelID, m.GuildID, m.Content)
	}

	_, err := (*database_utils.Database).Methods.InsertUserEvent.Exec(database_utils.USER_MESSAGE_WRITE, time.Now(), m.Author.ID, common.FormatUsername(m.Author), nil, nil, fmt.Sprintf("Channel <#%v>", m.ChannelID))
	if err != nil {
		log.Printf("ERROR IN ONMESSAGE: %v", err)
	}
}
