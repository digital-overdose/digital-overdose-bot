package handler

import (
	"database/sql"
	"fmt"
	"time"

	"atomicmaya.me/digital-overdose-bot/src/common"
	"atomicmaya.me/digital-overdose-bot/src/db"
	"atomicmaya.me/digital-overdose-bot/src/db/db_events"
	"atomicmaya.me/digital-overdose-bot/src/db/db_instr"
	"github.com/bwmarrin/discordgo"
)

func OnMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	channel, err := s.Channel(m.ChannelID)
	if err != nil {
		common.Log("ERROR IN ONMESSAGE: %v", err) // TODO BETTER LOGGING
	}

	db.DatabaseAddToQueue(&db_instr.DatabaseInstruction{
		InstrType: db_instr.INSERT_USER_EVENT,
		Params: []any{
			db_events.USER_MESSAGE_WRITE,
			time.Now(),
			m.Author.ID, common.FormatUsername(m.Author),
			channel.ID, fmt.Sprintf("#%v (<#%v>)", channel.Name, channel.ID),
			fmt.Sprintf("ID %v\nContent (%d): '%v'", m.ID, len(m.Content), common.EncodeMessage(m.Content, 50)),
		},
	})

	response := make(chan *sql.Rows)
	db.RegisterResponseAwaiter(response)

	res <- response

	if err != nil {
		common.Log("ERROR IN ONMESSAGE: %v", err) // TODO BETTER LOGGING
	}
}

func OnMessageUpdate(s *discordgo.Session, m *discordgo.MessageUpdate) {
	channel, err := s.Channel(m.ChannelID)
	if err != nil {
		common.Log("ERROR IN ONMESSAGEUPDATE: %v", err) // TODO BETTER LOGGING
	}

	_, err = (*database_utils.Database).Methods.InsertUserEvent.Exec(
		db_events.USER_MESSAGE_UPDATE,
		time.Now(),
		m.Author.ID, common.FormatUsername(m.Author),
		channel.ID, fmt.Sprintf("#%v (<#%v>)", channel.Name, channel.ID),
		fmt.Sprintf("ID %v\nContent (%d): '%v'", m.ID, len(m.Content), common.EncodeMessage(m.Content, 50)),
	)
	if err != nil {
		common.Log("ERROR IN ONMESSAGEUPDATE: %v", err) // TODO FIX
	}
}

func OnMessageDelete(s *discordgo.Session, m *discordgo.MessageDelete) {
	channel, err := s.Channel(m.ChannelID)
	if err != nil {
		common.Log("ERROR IN ONMESSAGEDELETE: %v", err) // TODO BETTER LOGGING
	}

	message := m.BeforeDelete
	if message == nil {
		_, err = (*database_utils.Database).Methods.InsertUserEvent.Exec(
			database_utils.USER_MESSAGE_DELETE,
			time.Now(),
			"-1", "UNKNOWN",
			channel.ID, fmt.Sprintf("#%v (<#%v>)", channel.Name, channel.ID),
			fmt.Sprintf("CONTENT UNKNOWN - ID %v/%v/%v", m.GuildID, m.ChannelID, m.ID),
		)
	} else {
		_, err = (*database_utils.Database).Methods.InsertUserEvent.Exec(
			database_utils.USER_MESSAGE_DELETE,
			time.Now(),
			message.Author.ID, common.FormatUsername(message.Author),
			channel.ID, fmt.Sprintf("#%v (<#%v>)", channel.Name, channel.ID),
			fmt.Sprintf("ID %v\nContent (%d): '%v'", message.ID, len(message.Content), common.EncodeMessage(message.Content, 50)),
		)
	}

	if err != nil {
		common.Log("ERROR IN ONMESSAGEDELETE: %v", err) // TODO BETTER LOGGING
	}
}
