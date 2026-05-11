# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project

A CLI tool and Go library for working with [Semantic Versions](https://semver.org/).
Module path: `github.com/ffurrer2/semver/v2`. Requires Go 1.26+. MIT-licensed.

## Commands

Build, test, and lint use [Task](https://taskfile.dev/) (not Make or raw `go` commands):

```shell
task build                # Static binary → artifacts/bin/ (CGO_ENABLED=0, ldflags)
task test                 # go test -race -covermode atomic ./... + coverage HTML
task lint                 # golangci-lint + markdownlint + yamllint
task fmt                  # go mod edit -fmt + gofumpt --extra + goimports
task golangcilint:lint    # golangci-lint only
task clean                # rm -rf artifacts/
```

Run a single test package:
```shell
go test -race -v ./pkg/semver/...
```

Run a single Ginkgo test by label or focus:
```shell
go test -race -v ./pkg/semver/... -ginkgo.focus="NextMajor"
```

golangci-lint config lives at `.github/linters/.golangci.yml`.

## Architecture

```
cmd/semver/           CLI (Cobra). One file per subcommand, wired in init().
pkg/semver/           Public library: Parse, IsValid, CompareTo, Next*, Builder, BySemVer.
internal/pkg/cli/     CLI plumbing: Apply (per-item stdin) and Map (batch stdin).
internal/pkg/app/     Version info (linker-injected via ldflags).
internal/pkg/number/  Generic numeric parsing and comparison.
internal/pkg/predicate/ Functional combinators (And, Or, Negate) using samber/lo.
```

Key patterns:
- Subcommands use a dual-input pattern: args present → process directly; no args → read stdin via `cli.Apply` or `cli.Map`.
- `pkg/semver` must remain a standalone library — no imports of CLI, Cobra, or internal packages (except `internal/pkg/number`).

## Coding Requirements

- Every new source file must begin with `// SPDX-License-Identifier: MIT` (or `#` variant for YAML/shell).
- Formatting: gofumpt (with `--extra`) + goimports (local prefix: `github.com/ffurrer2/semver`). Run `task fmt` before committing.
- Import order: stdlib, then third-party, then local (`github.com/ffurrer2/semver/v2/...`).
- Max line length: 160 chars.
- Linting uses `default: all` with specific disables. depguard strictly restricts imports:
  - **Allowed**: stdlib, `github.com/ffurrer2/semver/v2`, `github.com/spf13/cobra`, `github.com/samber/lo`, `github.com/onsi` (Ginkgo/Gomega), `github.com/go-task/slim-sprig/v3`
  - **Denied**: `math/rand` (use `math/rand/v2`), `github.com/pkg/errors` (use stdlib `errors`), `golang.org/x/net/context` (use stdlib `context`)
  - Adding a new dependency requires updating the depguard allowlist in `.github/linters/.golangci.yml`.

## Testing

- Tests use **Ginkgo v2 / Gomega** (`Describe`/`DescribeTable`), not plain table-driven tests.
- Each package has a `*_suite_test.go` file that bootstraps Ginkgo.
- When adding tests, match the style in `pkg/semver/semver_test.go`.

## SemVer Spec Invariants

- `NextMajor/NextMinor/NextPatch` on a prerelease version strips prerelease/build metadata **without incrementing** the numeric component.
- `CompareTo` implements SemVer spec section 11 (precedence rules). Build metadata does not affect ordering.
- Overflow in `Next*` panics with exported sentinel errors.

## Release Pipeline

- GoReleaser (`build/package/.goreleaser.yaml`) builds static multi-platform binaries (darwin/linux/windows, amd64/arm64).
- Docker images published to GHCR and Docker Hub with multi-tag strategy (latest, major, major.minor, full version).
- Distribution: Homebrew tap (`ffurrer2/tap`), Scoop bucket, GitHub releases with cosign signatures.
- Version string injected via ldflags into `internal/pkg/app.version` — computed by `scripts/version` script from git tags.
- CI runs `task test` + `task lint` + GoReleaser snapshot on PRs; full release triggers on `v*.*.*` tags.
