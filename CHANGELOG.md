# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [Released]

## `0.1.4` - 2022-09-17
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

[Unreleased]: https://github.com/digital-overdose/digital-overdose-bot/compare/v0.1.3...HEAD
[0.1.3]: https://github.com/digital-overdose/digital-overdose-bot/compare/v0.1.2...v0.1.3
[0.1.2]: https://github.com/digital-overdose/digital-overdose-bot/compare/v0.1.1...v0.1.2
[0.1.1]: https://github.com/digital-overdose/digital-overdose-bot/compare/v0.1.0...v0.1.1
[0.1.0]: https://github.com/digital-overdose/digital-overdose-bot/compare/v0.0.2...v0.1.0
[0.0.2]: https://github.com/digital-overdose/digital-overdose-bot/compare/v0.0.1...v0.0.2
[0.0.1]: https://github.com/digital-overdose/digital-overdose-bot/releases/tag/v0.0.1

