package handler

import (
	"fmt"
	"time"

	"atomicmaya.me/digital-overdose-bot/src/common"
	database_utils "atomicmaya.me/digital-overdose-bot/src/db"
	"github.com/bwmarrin/discordgo"
)

func OnReady(s *discordgo.Session, r *discordgo.Ready) {
	common.Log("Running in version: v%v", common.VERSION)
	common.Log("Logged in as: '%v'", common.FormatUsername(s.State.User))

	_, err := (*database_utils.Database).Methods.InsertOpsEvent.Exec(database_utils.SYSTEM_START, time.Now(), fmt.Sprintf("Running in version: v%v", common.VERSION))
	if err != nil {
		common.Log("ERROR IN ONREADY: %v", err)
	}
}
