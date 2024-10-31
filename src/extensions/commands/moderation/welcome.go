package extensions

import (
	"fmt"

	"atomicmaya.me/digital-overdose-bot/src/common"
	"github.com/bwmarrin/discordgo"
)

// Modifies the relevant roles and welcomes them into the server.
// Streamlines the manual process we were previously using.
func WelcomeUser(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if ok, _ := common.CheckHasPermissions(i, s, discordgo.PermissionManageRoles); !ok {
		return
	}

	if ok := common.ShouldExecutionBeSkippedIfDev(true); ok {
		return
	}

	options := i.ApplicationCommandData().Options
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	opt, ok := optionMap["user"]
	if !ok {
		return
	}

	member, err := s.GuildMember(*common.GuildID, opt.UserValue(nil).ID)
	if err != nil {
		common.LogToServer(common.Log("Failed to retrieve Member (ID: %v) from Guild %v", opt.UserValue(nil).ID, *common.GuildID), s)
	}

	userName := ""
	if member.User.Discriminator != "" { // Old username format
		userName = fmt.Sprintf("'%v#%v' (ID: %v)", member.User.Username, member.User.Discriminator, member.User.ID)
	} else {
		userName = fmt.Sprintf("'%v' (ID: %v)", member.User.Username, member.User.ID)
	}

	if *common.HumanRoleID != "" {
		common.LogToServer(common.Log("[+] Adding 'Human' role to %v", userName), s)
		err := s.GuildMemberRoleAdd(*common.GuildID, opt.UserValue(nil).ID, *common.HumanRoleID)

		if err != nil {
			common.LogToServer(common.Log("[x] Failed adding 'Human' role to %v", userName), s)
		}
	}

	if *common.MemberRoleID != "" {
		common.LogToServer(common.Log("[+] Adding 'Member' role to %v", userName), s)
		err := s.GuildMemberRoleAdd(*common.GuildID, opt.UserValue(nil).ID, *common.MemberRoleID)

		if err != nil {
			common.LogToServer(common.Log("[x] Failed adding 'Member' role to %v", userName), s)
		}
	}

	common.LogToServer(common.Log("[+] Removing 'Verification' role from %v", userName), s)
	err = s.GuildMemberRoleRemove(*common.GuildID, opt.UserValue(nil).ID, *common.VerificationRoleID)

	if err != nil {
		common.LogToServer(common.Log("[x] Failed removing 'Verification' role from %v", userName), s)
	}

	// TODO CONVERT TO EMBED

	if *common.MainChannelID != "" {
		formattedMessage := fmt.Sprintf(`Welcome <@%v>!
Feel free to introduce yourself to the community in the <#783109920240697414> section and grab some <#687232316061384779> and <#887783566916866069>.
Please remember the <#687239516800548894> and give us a shout if you need anything!`, opt.UserValue(nil).ID)

		common.LogToServer(common.Log("[+] Welcomed %v in main channel.", userName), s)

		_, err := s.ChannelMessageSend(*common.MainChannelID, formattedMessage)

		if err != nil {
			common.LogToServer(common.Log("[x] Failed to welcome %v in main channel.", userName), s)
		}
	}
}
