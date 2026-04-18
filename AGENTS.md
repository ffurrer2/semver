# AGENTS Guide for `semver`

## Project overview

A CLI tool for working with [Semantic Versions](https://semver.org/). Module path: `github.com/ffurrer2/semver/v2`. MIT-licensed.

## Directory layout

```text
cmd/semver/          CLI entrypoint (Cobra); one file per command, wired in init()
pkg/semver/          Public domain library (Parse, IsValid, CompareTo, Next*, Builder)
internal/pkg/cli/    Shared CLI plumbing: Apply (stream-per-item), Map (batch)
internal/pkg/app/    Build metadata/version info (linker-injected via ldflags)
internal/pkg/number/ Numeric parsing (ParseUint, MustParseUint) and generic CompareInt
internal/pkg/predicate/ Functional combinators (And, Or, Negate) using samber/lo
build/package/       Dockerfile + .goreleaser.yaml
tasks/               Task include files (GoTasks, GoreleaserTasks, LicensedTasks, etc.)
scripts/version      Shell script: tag/branch/dirty-aware version string generation
test/                Container structure tests (semver_container_test.yml)
.github/workflows/   CI, release, golangci-lint, CodeQL, actionlint, etc.
.github/linters/     golangci-lint config (.golangci.yml)
```

## Architecture

### CLI layer (`cmd/semver/`)

- `semver.go`: root command, custom usage template, wires stdin/stdout/stderr.
- Each subcommand (`compare`, `filter`, `format`, `next`, `sort`, `validate`, `version`) lives in its own file.
- `next` is a parent command; `major.go`, `minor.go`, `patch.go` are subcommands of `next`.
- Commands accepting variadic input follow a dual-input pattern:
  - CLI args present -> process args directly.
  - No args -> read stdin line-by-line via `cli.Apply` (per-item) or `cli.Map` (batch).
- Error handling: `semver.Parse(...)` -> `cmd.PrintErrf("error: %v\n", err)` -> `os.Exit(1)`.
- `validate` is the intentional exception: uses `semver.IsValid(...)` and exits `1` on first invalid input without printing an error message.
- `filter` composes predicates via `predicate.Or(...)` instead of nested if-chains. Flags (`--invalid`, `--releases`, `--pre-releases`, `--build-metadata`) are marked mutually exclusive where appropriate.
- `format` compiles Go templates once with `template.New(...).Funcs(sprig.FuncMap())` and executes per item. Template fields: `Major`, `Minor`, `Patch`, `PreRelease` (string), `BuildMetadata` (string).
- `sort` is the batch pattern: `cli.Map` collects all input, then `sort.Sort(semver.BySemVer(...))`. Supports `--reverse`.
- `version` supports `--json` and `--short` (mutually exclusive).

### Domain library (`pkg/semver/`)

- `SemVer` struct: `Major`, `Minor`, `Patch` (uint), `PreRelease`, `BuildMetadata` ([]string).
- Parsing is regex-based (`NamedGroupsPattern`), extracted via named capture groups.
- `CompareTo` compares major/minor/patch numerically, then prerelease precedence per spec. Build metadata does **not** affect ordering.
- `NextMajor/NextMinor/NextPatch`: if current version is prerelease, strip prerelease/build metadata without incrementing. Otherwise increment the relevant component. Overflow panics with exported errors (`ErrNextMajorUintOverflow`, etc.).
- `Builder`: fluent builder pattern returning immutable copies. `Build()` returns `(*SemVer, bool)`.
- `BySemVer`: implements `sort.Interface`.

### Internal packages

- `cli.Apply(args, reader, fn)`: iterates args or stdin lines, calls `fn` per item, then `os.Exit(0)`.
- `cli.Map(args, reader, fn)`: collects args or stdin lines into a slice, calls `fn` once with the full slice, then `os.Exit(0)`.
- `number.CompareInt[T constraints.Integer]`: generic three-way comparison.
- `predicate.And/Or/Negate`: generic higher-order predicate combinators using `lo.Reduce`.

## Build, test, and lint

Use [Task](https://taskfile.dev/) targets (not ad-hoc commands):

| Target                     | What it does                                                                       |
| -------------------------- | ---------------------------------------------------------------------------------- |
| `task build`               | Static binary to `artifacts/bin` (CGO_ENABLED=0, ldflags with version)             |
| `task test`                | `go test -race -covermode atomic ./...` + HTML coverage + container structure test |
| `task lint`                | golangci-lint + markdownlint + yamllint                                            |
| `task fmt`                 | `go mod edit -fmt` + gofumpt + goimports                                           |
| `task scan`                | Dockle + Trivy (image + filesystem) against `ghcr.io/ffurrer2/semver:latest`       |
| `task clean`               | Delete `artifacts/`                                                                |
| `task goreleaser:check`    | Validate GoReleaser config                                                         |
| `task goreleaser:build`    | Local GoReleaser build (`--single-target --snapshot`)                              |
| `task goreleaser:snapshot` | Full snapshot release (Docker images included)                                     |
| `task upgrade`             | Upgrade major + indirect Go dependencies via gomajor                               |
| `task version`             | Print versions of all external tools                                               |

Version strings are generated by `scripts/version` (uses git tags, branch names, commit hashes, dirty state).

## Testing conventions

- Tests use **Ginkgo v2 / Gomega** (`Describe`/`DescribeTable` style), not plain `testing.T` table tests.
- Suite files: `*_suite_test.go` in each package.
- `pkg/semver/semver_test.go` has extensive coverage; mirror its style and thoroughness when adding behavior.
- `pkg/semver/builder_test.go` covers the builder pattern.
- Container structure tests in `test/semver_container_test.yml` verify the Docker image (binary presence, permissions, OCI labels, command output).

## CI/CD

### Workflows (`.github/workflows/`)

| Workflow                | Trigger                                | Purpose                                                                             |
| ----------------------- | -------------------------------------- | ----------------------------------------------------------------------------------- |
| `ci.yml`                | push/PR to main                        | Build, test, GoReleaser snapshot, container structure test, Trivy + Anchore scans   |
| `release.yml`           | tag push `v*.*.*`                      | Full GoReleaser release (GitHub, Docker Hub, GHCR, Homebrew, Scoop, cosign signing) |
| `golangci-lint.yml`     | push/PR to main (Go/lint file changes) | golangci-lint with config from `.github/linters/.golangci.yml`                      |
| `codeql.yml`            | Code scanning                          | CodeQL analysis                                                                     |
| `dependency-review.yml` | PR                                     | Dependency review                                                                   |
| `actionlint.yml`        | Workflow file changes                  | Lint GitHub Actions workflows                                                       |
| `devskim.yml`           | Security                               | DevSkim security analysis                                                           |
| `markdownlint.yml`      | Markdown changes                       | Markdown linting                                                                    |
| `yamllint.yml`          | YAML changes                           | YAML linting                                                                        |
| `licensed.yml`          | Dependency changes                     | License compliance checking                                                         |

### Dependency management

- **Renovate** (`.github/renovate.json5`): auto-updates Go modules, GitHub Actions (pinned by SHA), pre-commit hooks, and `*_VERSION` env vars in workflows via custom regex manager.
- **Licensed** (`.licensed.yml`): tracks dependency licenses; allowed: `apache-2.0`, `bsd-3-clause`, `mit`.
- **Pre-commit** (`.pre-commit-config.yaml`): gitleaks for secret scanning.

## SemVer behavior that must not regress

- Parsing regex (`NamedGroupsPattern`) follows the official semver.org grammar exactly.
- `CompareTo` implements SemVer spec section 11 (precedence rules), including numeric vs. lexical prerelease comparison and set-size tiebreaking.
- `NextMajor/NextMinor/NextPatch` on a prerelease version strips prerelease/build metadata **without incrementing** the numeric component. This is intentional.
- Overflow in `Next*` panics with exported sentinel errors; tests assert this behavior.
- `Builder.Build()` validates via `IsValid()` and returns a bool; it does not panic.
- `BySemVer` sort order matches spec precedence (not string sorting).

## Coding guardrails

- **SPDX header**: every new source file must start with `// SPDX-License-Identifier: MIT` (or `# SPDX-License-Identifier: MIT` for YAML/shell).
- **Formatting**: gofumpt (with `--extra`) + goimports (local prefix: `github.com/ffurrer2/semver`). Run `task fmt` before committing.
- **Linting**: golangci-lint v2 config uses `default: all` with specific disables. Notable: depguard restricts imports to known dependencies. Max line length: 160 chars. No `math/rand` (use `math/rand/v2`).
- **Import ordering**: stdlib, then third-party, then local (`github.com/ffurrer2/semver/v2/...`), enforced by goimports.
- **Domain isolation**: `pkg/semver` must not import CLI, cobra, or internal packages (except `internal/pkg/number`). Keep it reusable as a library.
- **No `main` package dependencies leaking**: `internal/` packages are for CLI use only.
- **GitHub Actions pins**: all actions are pinned by full SHA with a comment showing the version tag. Renovate manages updates.
- **Go version**: read from `go.mod` (`go.mod` is the single source of truth; CI uses `go-version-file`).
