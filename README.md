# DISCORD GO BOT

## HOW TO RUN

### WITH `.env`

Run with `go run .`

### WITHOUT `.env`

Run with `go run . --guild <GUILD_ID> --token <BOT_TOKEN> --role
 <VERIFICATION_ROLE_ID> --wall <VERIFICATION_CHANNEL_ID> --mod <MOD_ACTION_LOGS_CHANNEL_ID> --debug <DEBUG_CHANNEL_ID>`

## DEMO `.env`

```txt
GUILD=
TOKEN=
VERIFICATION_ROLE_ID=
VERIFICATION_CHANNEL_ID=
MOD_ACTION_CHANNEL_ID=
DEBUG_CHANNEL_ID=

# ADDITIONAL FEATURE: welcome.go
HUMAN_ROLE_ID=
MEMBER_ROLE_ID=
MAIN_CHANNEL_ID=

# ADDITIONAL FEATURE: Abuse Warning
STAFF_CHANNEL_ID=
```

## TODO

- [x] Ability to Cron
- [x] Message mod-logs
- [x] Implement DM
- [x] Implement warn
- [x] Implement kick
- [x] Fix timeout
