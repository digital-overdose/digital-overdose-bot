package database_utils

type eventType struct {
	ID   int
	Name string
}

type EventType int

const (
	USER_SERVER_JOIN  EventType = 0x00
	USER_SERVER_LEAVE EventType = 0x01
	USER_ROLE_ADD     EventType = 0x02
	USER_ROLE_REMOVE  EventType = 0x03
	USER_MUTE_START   EventType = 0x04
	USER_MUTE_END     EventType = 0x05
	USER_WARN_START   EventType = 0x06
	USER_WARN_END     EventType = 0x07
	USER_BAN_START    EventType = 0x08
	USER_BAN_END      EventType = 0x09
	USER_RBAC_SUCCESS EventType = 0x0A
	USER_RBAC_FAIL    EventType = 0x0B

	USER_MESSAGE_WRITE  EventType = 0x10
	USER_MESSAGE_EDIT   EventType = 0x11
	USER_MESSAGE_DELETE EventType = 0x12

	USER_VC_JOIN         EventType = 0x20
	USER_VC_LEAVE        EventType = 0x21
	USER_VC_STREAM_START EventType = 0x22
	USER_VC__STREAM_END  EventType = 0x23

	SYSTEM_START  EventType = 0xA0
	SYSTEM_STOP   EventType = 0xA1
	SYSTEM_UPDATE EventType = 0xA2

	CRON_SYSTEM_START_LOG EventType = 0xB0
	CRON_SYSTEM_END_LOG   EventType = 0xC0

	CRON_MOD_START_CLEAN EventType = 0xD0
	CRON_MOD_END_CLEAN   EventType = 0xD2
	CRON_MOD_START_PRUNE EventType = 0xD1
	CRON_MOD_END_PRUNE   EventType = 0xE1
)

var eventTypes = map[int]string{
	0x00: "User joined Server",
	0x01: "User left Server",
	0x02: "User was added Role",
	0x03: "User was removed Role",
	0x04: "User was Muted",
	0x05: "User was Unmuted",
	0x06: "User was Warned",
	0x07: "User was Unwarned",
	0x08: "User was Banned",
	0x09: "User was Unbanned",
	0x0A: "User passed a permissions check",
	0x0B: "User failed a permissions check",

	0x10: "User wrote a message",
	0x11: "User edited a message",
	0x12: "User deleted a message",

	0x20: "User joined a voice channel",
	0x21: "User left a voice channel",
	0x22: "User started streaming in a voice channel",
	0x23: "User stopped streaming in a voice channel",

	0xA0: "Server started",
	0xA1: "Server shutting down",
	0xA2: "Server updating",

	0xB0: "Cron job started: 'Management: Log cycling'",
	0xC0: "Cron job ended: 'Management: Log cycling'",

	0xD0: "Cron job started: 'Automod: Verification Prune'",
	0xE0: "Cron job ended: 'Automod: Verification Prune'",

	0xD1: "Cron job started: 'Automod: Automated Unmute Check'",
	0xE1: "Cron job ended: 'Automod: Automated Unmute Check'",
}

func GetEventType(v int) eventType {
	return eventType{ID: v, Name: eventTypes[v]}
}
