package ext

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"atomicnicos.me/digital-overdose-bot/common"
	"github.com/bwmarrin/discordgo"
	"github.com/cavaliergopher/grab/v3"
)

func UpgradeBot(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if ok, _ := common.CheckHasPermissions(i, s, discordgo.PermissionAdministrator); !ok {
		return
	}

	if ok := common.ShouldExecutionBeSkippedIfDev(true); ok {
		return
	}

	if ok := *common.UpgradeReleaseURL != ""; !ok {
		log.Printf("[x] Release URL validation missing.")
		return
	}

	options := i.ApplicationCommandData().Options
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	opt, ok := optionMap["version"]
	if !ok {
		return
	}

	version := opt.StringValue()

	common.LogAndSend(fmt.Sprintf("[⚠] Upgrade to version `%v` starting now!", version), s)
	common.LogAndSend(fmt.Sprintf("---- URL: %v", fmt.Sprintf(*common.UpgradeReleaseURL, version, version)), s)

	resp, err := grab.Get(".", fmt.Sprintf(*common.UpgradeReleaseURL, version, version))
	if err != nil {
		common.LogAndSend(fmt.Sprintf("[x] Error downloading the release: %v", err), s)
		return
	}

	exe := resp.Filename

	common.LogAndSend(fmt.Sprintf("[⇑] Successfully downloaded: %v", exe), s)

	ext := ""
	if runtime.GOOS == "windows" {
		ext = ".exe"
	}

	err = os.Rename(fmt.Sprintf("./digital-overdose-bot%v", ext), fmt.Sprintf("./digital-overdose-bot.old%v", ext))

	if err != nil {
		common.LogAndSend(fmt.Sprintf("[x] Error renaming the old executable: %v", err), s)
		return
	}

	err = os.Rename(fmt.Sprintf("./%v", exe), fmt.Sprintf("./digital-overdose-bot%v", ext))
	if err != nil {
		common.LogAndSend(fmt.Sprintf("[x] Error renaming the new executable: %v", err), s)
		return
	}

	common.LogAndSend(fmt.Sprintf("[+] Renamed executable to expected pattern"), s)

	os.Remove(fmt.Sprintf("./digital-overdose-bot.old%v", ext))
	common.LogAndSend(fmt.Sprintf("[+] Removed the outdated executable."), s)

	common.LogAndSend(fmt.Sprintf("[+] Killing current bot, `systemd` will restart it in 10 seconds.."), s)
	os.Exit(42)
}
