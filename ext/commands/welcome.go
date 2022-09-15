package ext

import (
	"fmt"
	"log"

	"atomicnicos.me/digital-overdose-bot/common"
	"github.com/bwmarrin/discordgo"
)

// Modifies the relevant roles and welcomes them into the server.
// Streamlines the manual process we were previously using.
func WelcomeUser(s *discordgo.Session, i *discordgo.InteractionCreate) {
	common.CheckHasPermissions(i, s, discordgo.PermissionViewAuditLogs|discordgo.PermissionManageRoles)

	options := i.ApplicationCommandData().Options
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	opt, ok := optionMap["user"]
	if !ok {
		return
	}

	usr_str := fmt.Sprintf("'%v#%v' (ID: %v)", opt.UserValue(nil).Username, opt.UserValue(nil).Discriminator, opt.UserValue(nil).ID)

	if *common.HumanRoleID != "" {
		log.Printf("[+] Adding 'Human' role to %v", usr_str)
		err := s.GuildMemberRoleAdd(*common.GuildID, opt.UserValue(nil).ID, *common.HumanRoleID)
		if err != nil {
			log.Printf("[x] Failed adding 'Human' role to %v", usr_str)
		}
	}

	if *common.MemberRoleID != "" {
		log.Printf("[+] Adding 'Member' role to %v", usr_str)
		err := s.GuildMemberRoleAdd(*common.GuildID, opt.UserValue(nil).ID, *common.MemberRoleID)
		if err != nil {
			log.Printf("[x] Failed adding 'Member' role to %v", usr_str)
		}
	}

	log.Printf("[+] Removing 'Verification' role from %v", usr_str)
	err := s.GuildMemberRoleRemove(*common.GuildID, opt.UserValue(nil).ID, *common.VerificationRoleID)
	if err != nil {
		log.Printf("[x] Failed removing 'Verification' role from %v", usr_str)
	}

	if *common.MainChannelID != "" {
		formatted_msg := fmt.Sprintf("Welcome <@%v>! Please remember the <#687239516800548894>, perhaps tell us something <#783109920240697414>, grab yourself some <#687232316061384779> and perhaps drop <#783110016076349450>!", opt.UserValue(nil).ID)

		log.Printf("[+] Welcomed %v in main channel.", usr_str)
		_, err := s.ChannelMessageSend(*common.MainChannelID, formatted_msg)

		if err != nil {
			log.Printf("[x] Failed to welcome %v in main channel.", usr_str)
		}
	}
}
