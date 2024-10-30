package cron

import (
	commands_management "atomicmaya.me/digital-overdose-bot/src/extensions/commands/management"
	commands_moderation_mute "atomicmaya.me/digital-overdose-bot/src/extensions/commands/moderation/mute"

	"atomicmaya.me/digital-overdose-bot/src/common"
	"github.com/bwmarrin/discordgo"
)

type CronJob struct {
	Name       string
	CronString string
	Job        func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

// Stores all of the automated commands which will be executed by the server on a scheduled basis.
var CronJobs = []*CronJob{
	// Verification pruning happens every day at 12:00 UTC.
	{
		Name:       "Automod: Verification Prune",
		CronString: "0 12 */1 * *",
		Job:        commands_management.PurgeVerification,
	},
	// Checking whether a user should be unmuted happens every 10 minutes.
	{
		Name:       "Automod: Automated Unmute Check",
		CronString: "*/10 * * * *",
		Job:        commands_moderation_mute.UnmuteAutomated,
	},
	// Cycling the log file happens every day at 00:00 UTC.
	{
		Name:       "Management: Log cycling",
		CronString: "0 0 * * *",
		Job:        common.PaginateLog,
	},
}
