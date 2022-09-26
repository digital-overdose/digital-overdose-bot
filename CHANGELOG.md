# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [Released]

## [0.1.6] - 2022-09-26

### Added
- Version file that exports the current version number.
- Welcome message (on join).
- Handler folder for all the event handlers.

### Changed
- File structure of the "extensions" (the various commands) to differentiate semantically. (moderation actions, management, etc)

### Removed
- Duplicate logging initialization in `main.go`.
- Various debug functions/experiments that are now sunset.

## `0.1.5-hotfix` - 2022-09-22
### Fixes
- Issue #5: [Link to issue](https://github.com/digital-overdose/digital-overdose-bot/issues/5s)
### Changed
- How the `cron` scheduler is provisioned, as the current version would mix up memory addresses. TL;DR: Can't run multiple tasks in one scheduler. And one needs to dereference the provided job.

## [0.1.5] - 2022-09-21
### Fixes
- Issue #3: [Link to issue](https://github.com/digital-overdose/digital-overdose-bot/issues/3)
- Issue #4: [Link to issue](https://github.com/digital-overdose/digital-overdose-bot/issues/4)

### Added
- Self-upgrade capability for the bot, via a command. (reliant on systemd restart feature).
- Log pagination at midnight every day.

### Changed
- Target for verification purge channel

## [0.1.4] - 2022-09-17
### Fixes
- Issue #2: [Link to issue](https://github.com/digital-overdose/digital-overdose-bot/issues/2)

### Changed
- `purge-verification`: Cleanliness by editing counter messages instead of reposting them.
- `purge-verification`: Message in debug channel to indicate that the challenge was launched by the cron.
- `welcome`: Changes how the user string is determined.
- Powershell / CMD / sh build scripts to something more... clean.

## `0.1.3-hotfix` - 2022-09-16
### Fixes
- Issue #1: [Link to issue](https://github.com/digital-overdose/digital-overdose-bot/issues/1)

### Changed
- Restoring changes to the "Time Last Seen" calculation to take into account the most recent of the last messages, and not be overwritten.
- Slight changes to build scripts.

## [0.1.3] - 2022-09-15
### Added
- Cron feature.
- Cron job for verification purge.
- Verification! (`/welcome <user>`)
- Build scripts!
- Build folder!
- RBAC infringement message in staff chat!

### Changed
- Renamed list-purge-candidates to purge-verification to stay consistent with reality.
- Made logging to files optional (for contexts where no log files can be created)
- Role-based access control! Now it actually works!

## [0.1.2] - 2022-09-15
### Added
- CHANGELOG.md.
- Relevant comments.
- Build artifact to `.gitignore`.

### Changed
- Simplified extension code.
- Migrated flag system to `common` module.

### Removed
- Syntactic sugar.

## [0.1.1] - 2022-09-15
### Added
- Environment variable usage.
- Dynamic flags usage.
- Writing to the moderation-action-logs channel if it is defined.

## Changed
- Migrated dual lane logging to `common` package.

## [0.1.0] - 2022-09-13
### Added 
- Kick functionality.
- Warn message functionality.
- DM user functionality.
- Logging.

### Changed
- Repository location.

## [0.0.2] - 2022-09-13
### Added
- Permission check feature.
- `common` module for generic code.

### Changed
- Migrated commands code to `ext` modules (extensions, supposed to be plug and play).

## [0.0.1] - 2022-09-12
### Added
- Initial run code.

[Unreleased]: https://github.com/digital-overdose/digital-overdose-bot/compare/v0.1.6...HEAD
[0.1.6]: https://github.com/digital-overdose/digital-overdose-bot/compare/v0.1.5...v0.1.6
[0.1.5]: https://github.com/digital-overdose/digital-overdose-bot/compare/v0.1.4...v0.1.5
[0.1.4]: https://github.com/digital-overdose/digital-overdose-bot/compare/v0.1.3...v0.1.4
[0.1.3]: https://github.com/digital-overdose/digital-overdose-bot/compare/v0.1.2...v0.1.3
[0.1.2]: https://github.com/digital-overdose/digital-overdose-bot/compare/v0.1.1...v0.1.2
[0.1.1]: https://github.com/digital-overdose/digital-overdose-bot/compare/v0.1.0...v0.1.1
[0.1.0]: https://github.com/digital-overdose/digital-overdose-bot/compare/v0.0.2...v0.1.0
[0.0.2]: https://github.com/digital-overdose/digital-overdose-bot/compare/v0.0.1...v0.0.2
[0.0.1]: https://github.com/digital-overdose/digital-overdose-bot/releases/tag/v0.0.1

