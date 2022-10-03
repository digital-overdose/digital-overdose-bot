# DISCORD GO BOT

## HOW TO RUN

### WITH `.env`

Run with `go run .`

### WITHOUT `.env`

Run with `go run . --guild <GUILD_ID> --token <TOKEN> --role <VERIFICATION_ROLE> --wall <VERIFICATION_CHANNEL_ID> --mod <MOD_ACTION_LOGS_CHANNEL_ID> --mod-thread <MOD_ACTION_LOGS_THREAD_ID> --debug <DEBUG_CHANNEL_ID> --human <HUMAN_ROLE_ID> --member <MEMBER_ROLE_ID> --main <MAIN_CHANNEL_ID> --staff <ABUSE_CHANNEL_ID> --upgrade <UPGRADE_RELEASE_PATH>`

## DEMO `.env`

```txt
GUILD=
TOKEN=
VERIFICATION_ROLE_ID=
VERIFICATION_CHANNEL_ID=
MOD_ACTION_CHANNEL_ID=
MOD_ACTION_THREAD_ID=
DEBUG_CHANNEL_ID=

# ADDITIONAL FEATURE: welcome.go
HUMAN_ROLE_ID=
MEMBER_ROLE_ID=
MAIN_CHANNEL_ID=

# ADDITIONAL FEATURE: Abuse Warning
STAFF_CHANNEL_ID=

# ADDITIONAL FEATURE: upgrade.go
UPGRADE_RELEASE_PATH=https://github.com/digital-overdose/digital-overdose-bot/releases/download/v%v/digital-overdose-bot-v%v-linux-amd64

# ADDITIONAL FEATURE: Fully Comprehensive Mod Logs
PRIVATE_MOD_LOGS_CHANNEL_ID=
PRIVATE_CHAT_LOGS_CHANNEL_ID=

# ADDITIONAL FEATURE: Mutes (mute-new.go, mute-remove.go, mute-list.go)
MUTE_ROLE_ID=
```

## TODO

### Features (Future)

- [ ] Cron job to check people that are muted.

- [x] On Join message.
- [ ] `/warn` ⇾ Warn a user for a behavior (user, reason)
- [ ] `/unwarn` ⇾ Remove an attributed warn.
- [ ] `/ban` ⇾ Ban a user for a behavior (user, reason) + DM.
- [ ] `/unban` ⇾ Unban someone
- [ ] `/stats` ⇾ Number of people interacting over 2 weeks. Channel usage. Keep message ID in a file? (Cron)

#### Reqs

- Small DBMS, sqlite?

### Features (Important ⇾ Critical)

- [x] Ability to download new binary
- [x] Ability to restart on a specific binary
- [x] Ability to list running instances of bot
- [x] Ability to stop a specific instance from reacting to commands (for testing)

### Done

- [x] Ability to Cron
- [x] Message mod-logs
- [x] Implement DM
- [x] Implement warn
- [x] Implement kick
- [x] Fix timeout
