# DISCORD GO BOT

## HOW TO RUN

### WITH `./env/.env`

Run with `go run .`

### WITHOUT `.env`

Run with `go run . --guild <GUILD_ID> --token <TOKEN> --role <VERIFICATION_ROLE> --wall <VERIFICATION_CHANNEL_ID> --mod <MOD_ACTION_LOGS_CHANNEL_ID> --mod-thread <MOD_ACTION_LOGS_THREAD_ID> --debug <DEBUG_CHANNEL_ID> --human <HUMAN_ROLE_ID> --member <MEMBER_ROLE_ID> --main <MAIN_CHANNEL_ID> --staff <ABUSE_CHANNEL_ID> --upgrade <UPGRADE_RELEASE_PATH>`

## Template `./env/.env`

See [template.env](./env/template.env)

## TODO

### Features (Future)

- [ ] Cron job to check people that are muted.

- [x] On Join message.
- [ ] `/warn` ⇾ Warn a user for a behavior (user, reason)
- [ ] `/unwarn` ⇾ Remove an attributed warn.
- [ ] `/ban` ⇾ Ban a user for a behavior (user, reason) + DM.
- [ ] `/unban` ⇾ Unban someone
- [ ] `/stats` ⇾ Number of people interacting over 2 weeks. Channel usage. Keep message ID in a file? (Cron)
- [x] Refactor Database schema.

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
