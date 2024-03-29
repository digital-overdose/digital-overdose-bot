package extensions

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"atomicmaya.me/digital-overdose-bot/common"
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

	// Checks whether or not the source for binaries is set.
	if ok := *common.UpgradeReleaseURL != ""; !ok {
		log.Printf("[x] Release URL validation missing.")
		return
	}

	// Retrieves the user-supplied options.
	options := i.ApplicationCommandData().Options
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	// Retrieves the version from the user-supplied options.
	opt_version, ok := optionMap["version"]
	if !ok {
		return
	}

	version := opt_version.StringValue()

	common.LogAndSend(fmt.Sprintf("[⚠] Upgrade to version `%v` starting now!", version), s)
	common.LogAndSend(fmt.Sprintf("---- URL: %v", fmt.Sprintf(*common.UpgradeReleaseURL, version, version)), s)

	// Downloads the binary.
	resp, err := grab.Get(".", fmt.Sprintf(*common.UpgradeReleaseURL, version, version))
	if err != nil {
		common.LogAndSend(fmt.Sprintf("[x] Error downloading the release: %v", err), s)
		return
	}

	// Gets the downloaded filename.
	exe := resp.Filename

	common.LogAndSend(fmt.Sprintf("[⇑] Successfully downloaded: %v", exe), s)

	// Windows binaries specific code.
	ext := ""
	if runtime.GOOS == "windows" {
		ext = ".exe"
	}

	// Renames the current executable to an old version
	err = os.Rename(fmt.Sprintf("./digital-overdose-bot%v", ext), fmt.Sprintf("./digital-overdose-bot.old%v", ext))

	if err != nil {
		common.LogAndSend(fmt.Sprintf("[x] Error renaming the old executable: %v", err), s)
		return
	}

	// Renames the new executable to the correct path.
	err = os.Rename(fmt.Sprintf("./%v", exe), fmt.Sprintf("./digital-overdose-bot%v", ext))
	if err != nil {
		common.LogAndSend(fmt.Sprintf("[x] Error renaming the new executable: %v", err), s)
		return
	}

	common.LogAndSend("[+] Renamed executable to expected pattern.", s)

	// Modifies the permissions so that systemd can start it with the same rights.
	os.Chmod(fmt.Sprintf("./digital-overdose-bot%v", ext), 0755)

	// Removes the old version of the binary.
	os.Remove(fmt.Sprintf("./digital-overdose-bot.old%v", ext))
	common.LogAndSend("[+] Removed the outdated executable.", s)

	common.LogAndSend("[+] Killing current bot, `systemd` will restart it in 10 seconds.", s)
	os.Exit(42)
}
