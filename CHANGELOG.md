<!-- SPDX-License-Identifier: MIT -->  

# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

### Changed

### Deprecated

### Removed

### Fixed

### Security

## [2.3.0] - 2023-12-29

### Added

- Extend format command with the [slim-sprig](https://go-task.github.io/slim-sprig/) library

### Changed

- Bumped `alpine` base image from 3.18.4 to 3.19.0
- Updated dependencies

## [2.2.0] - 2023-09-25

### Added

- `--json` option to version command to print the complete build info as json

### Changed

- Removed "built at" info from version command
- Set `org.opencontainers.image.created` label to `1970-01-01T00:00:00Z`
- Dependencies updated
- Updated `alpine` base image to 3.18.4

## [2.1.0] - 2023-07-10

### Changed

- Dependencies updated
- Updated `alpine` base image to 3.18.2

## [2.0.2] - 2022-09-25

### Changed

- Dependencies updated

## [2.0.1] - 2022-09-03

### Fixed

- Fixed version command output

## [2.0.0] - 2022-09-03

### Added

- `-r/--reverse` option to sort command
- Additional filter command flags

### Changed

- Changed the default behavior of the filter command. Pre-release versions and versions containing build metadata are no longer printed by default.

## [1.8.0] - 2022-08-25

### Changed

- Dependencies updated
- Updated `alpine` base image to 3.16.2

## [1.7.0] - 2022-04-24 [YANKED]

## [1.6.0] - 2022-04-02

### Changed

- Dependencies updated
- Updated `alpine` base image to 3.15.3

## [1.5.0] - 2022-03-21

### Changed

- Dependencies updated
- Updated `alpine` base image to 3.15.1

## [1.4.0] - 2022-01-30

### Added

- Docker images are now available in the `GitHub Container Registry`

### Changed

- Dependencies updated

## [1.3.0] - 2022-01-29

### Changed

- Base image changed to `distroless`
- Non-root user added to alpine-based image

## [1.2.0] - 2022-01-01

### Added

- Support for Darwin and Linux ARM64

### Changed

- Dependencies updated
- Updated `alpine` base image to 3.15.0

## [1.1.0] - 2021-04-17

### Changed

- Dependencies updated
- Updated `alpine` base image to 3.13.5
- Improved CI workflow

## [1.0.1] - 2020-07-02

### Fixed

- Fixed UNIX pipeline support for the sort command

## [1.0.0] - 2020-06-24

### Added

- Initial release of `semver`

[Unreleased]: https://github.com/ffurrer2/semver/compare/v2.2.0...HEAD
[2.2.0]: https://github.com/ffurrer2/semver/compare/v2.1.2...v2.2.0
[2.1.0]: https://github.com/ffurrer2/semver/compare/v2.0.2...v2.1.0
[2.0.2]: https://github.com/ffurrer2/semver/compare/v2.0.1...v2.0.2
[2.0.1]: https://github.com/ffurrer2/semver/compare/v2.0.0...v2.0.1
[2.0.0]: https://github.com/ffurrer2/semver/compare/v1.8.0...v2.0.0
[1.8.0]: https://github.com/ffurrer2/semver/compare/v1.7.0...v1.8.0
[1.7.0]: https://github.com/ffurrer2/semver/compare/v1.6.0...v1.7.0
[1.6.0]: https://github.com/ffurrer2/semver/compare/v1.5.0...v1.6.0
[1.5.0]: https://github.com/ffurrer2/semver/compare/v1.4.0...v1.5.0
[1.4.0]: https://github.com/ffurrer2/semver/compare/v1.3.0...v1.4.0
[1.3.0]: https://github.com/ffurrer2/semver/compare/v1.2.0...v1.3.0
[1.2.0]: https://github.com/ffurrer2/semver/compare/v1.1.0...v1.2.0
[1.1.0]: https://github.com/ffurrer2/semver/compare/v1.0.1...v1.1.0
[1.0.1]: https://github.com/ffurrer2/semver/compare/v1.0.0...v1.0.1
[1.0.0]: https://github.com/ffurrer2/semver/compare/c171518f...v1.0.0
