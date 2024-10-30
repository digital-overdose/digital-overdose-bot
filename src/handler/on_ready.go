package handler

import (
	"fmt"
	"log"
	"time"

	"atomicmaya.me/digital-overdose-bot/src/common"
	database_utils "atomicmaya.me/digital-overdose-bot/src/db"
	"github.com/bwmarrin/discordgo"
)

func OnReady(s *discordgo.Session, r *discordgo.Ready) {
	log.Printf("Running in version: v%v", common.VERSION)
	log.Printf("Logged in as: '%v#%v'", s.State.User.Username, s.State.User.Discriminator)

	_, err := (*database_utils.Database).Methods.InsertOpsEvent.Exec(database_utils.SYSTEM_START, time.Now(), fmt.Sprintf("Running in version: v%v", common.VERSION))
	if err != nil {
		log.Printf("ERROR IN ONREADY: %w", err)
	}
}
