package cron

import (
	"atomicnicos.me/digital-overdose-bot/common"
	ext "atomicnicos.me/digital-overdose-bot/ext/commands"
	"github.com/bwmarrin/discordgo"
)

type CronJob struct {
	Name       string
	CronString string
	Job        func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

var CronJobs = []*CronJob{
	//{
	//	Name:       "Test Job",
	//	CronString: "*/1 * * * *",
	//	Job:        ext.TestCurrentFeature,
	//},
	{
		Name:       "Automod: Verification Prune",
		CronString: "0 12 */1 * *",
		Job:        ext.PurgeVerification,
	},
	{
		Name:       "Management: Log cycling",
		CronString: "0 0 * * *",
		Job:        common.PaginateLog,
	},
}
