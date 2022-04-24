<!-- SPDX-License-Identifier: MIT -->

# SemVer

[![CI](https://github.com/ffurrer2/semver/workflows/CI/badge.svg)](https://github.com/ffurrer2/semver/actions?query=workflow%3ACI)
[![Go Report Card](https://goreportcard.com/badge/github.com/ffurrer2/semver)](https://goreportcard.com/report/github.com/ffurrer2/semver)
[![MIT License](https://img.shields.io/github/license/ffurrer2/semver)](https://github.com/ffurrer2/semver/blob/main/LICENSE)
[![GitHub Release](https://img.shields.io/github/v/release/ffurrer2/semver?sort=semver)](https://github.com/ffurrer2/semver/releases/latest)

SemVer is a command-line utility for working with [Semantic Versions](https://semver.org/).

## Install

### Homebrew

```shell
brew install ffurrer2/tap/semver
```

### Scoop

```shell
scoop bucket add ffurrer2 https://github.com/ffurrer2/scoop-bucket
scoop install semver
```

### Build from source

```shell
go install github.com/ffurrer2/semver/cmd/semver@latest
```

## Usage

### help

```console
$ semver help
The semantic version utility

Usage:
  semver [command]

Available Commands:
  compare     Compare semantic versions
  filter      Filter semantic versions
  format      Format and print semantic versions
  help        Help about any command
  next        Increment semantic versions
  sort        Sort semantic versions
  validate    Validate semantic versions
  version     Print version information

Flags:
  -h, --help   help for semver

Use "semver [command] --help" for more information about a command.
```

### compare

```console
$ semver compare 1.0.0 1.0.0-alpha.1
1
```

### filter

```console
$ semver filter 1.0.0 1.0 v1.0.0 1.0.0-alpha.1
1.0.0
1.0.0-alpha.1
```

### format

```console
$ semver format 'v{{.Major}}.{{.Minor}}.{{.Patch}}' 1.0.0-alpha+001
v1.0.0
```

### next major/minor/patch

```console
$ semver next major 1.0.0-alpha+001
1.0.0
```

### sort

```console
$ semver sort 1.1.1 1.0.0 1.0.1
1.0.0
1.0.1
1.1.1
```

### validate

```console
$ semver validate 1.0.0-alpha+001
$ echo $?
0
```

### version

```console
$ semver version
semver version 1.0.0
commit: 224394b9d709277fe5e4965a6633535c597cd2ed
built at: 2020-06-18T20:22:25Z
```

## License

This project is licensed under the [MIT License](LICENSE).
