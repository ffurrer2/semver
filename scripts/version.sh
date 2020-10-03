#!/usr/bin/env bash
# SPDX-License-Identifier: MIT
set -euo pipefail

# ###################################################################################### #
# Prints the current project version                                                     #
#                                                                                        #
# Versioning follows Semantic Versioning 2.0.0 (https://semver.org):                     #
#   <version>           ::= <major>.<minor>.<patch>-<pre-release>+<build-metadata>       #
#   <major|minor|patch> ::= <numeric identifier>                                         #
#   <pre-release>       ::= <cleaned-branch-name>                                        #
#   <build-metadata>    ::= <yyyymmdd>.<commit-hash>                                     #
#                                                                                        #
# VERSION: 1.3.0                                                                         #
#                                                                                        #
# DEPENDENCIES:                                                                          #
#  - git                                                                                 #
#  - sed                                                                                 #
#  - tr                                                                                  #
#                                                                                        #
# EXAMPLES:                                                                              #
#  release version '1.0.0' tag:                 1.0.0                                    #
#  snapshot version 'develop' branch:           1.0.1-develop+20200101.42a4711           #
#  snapshot version 'hotfix/new-hotfix' branch: 1.0.1-hotfix-new-hotfix+20200101.42a4711 #
#                                                                                        #
# ###################################################################################### #

for cmd in {git,sed}; do
  if [[ ! -x "$(command -v "${cmd}")" ]]; then
    printf 'error: command not found: %s\n' "${cmd}" >&2
    exit 1
  fi
done

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
readonly SCRIPT_DIR

version() {
  # shellcheck source=./branch.sh
  source "${SCRIPT_DIR}/branch.sh"

  local version
  version="$(currentTag)"
  if [[ -z "${version}" ]]; then
    version="$(lastAccessibleTag)"
    version="$(nextPatchVersion "${version}")"

    local pre_release
    pre_release="$(branch)"
    # Remove leading zeros
    pre_release=$(printf '%s' "${pre_release}" | sed 's/^0*//')
    if [[ -n "${pre_release}" ]]; then
      version="${version}-${pre_release}"
    fi

    if [[ "$*" != "--short" ]]; then
      local commit_hash
      if [[ -n "$(git status --porcelain)" ]]; then
        commit_hash="dirty"
      else
        commit_hash="$(git rev-parse --short HEAD)"
      fi

      local build_metadata
      build_metadata="$(date -u +%Y%m%d).${commit_hash}"
      version="${version}+${build_metadata}"
    fi
  fi

  if ! isValid "${version}"; then
    printf 'error: invalid semantic version: %s\n' "${version}" >&2
    exit 1
  fi

  printf '%s' "${version}"
}

currentTag() {
  local tag
  tag="$(git describe --tags --exact-match 2>/dev/null)"
  printf '%s' "${tag}"
}

lastAccessibleTag() {
  local tag
  tag="$(git describe --tags --abbrev=0 2>/dev/null)"
  if [[ -z "${tag}" ]]; then
    tag="0.0.0"
  fi
  printf '%s' "${tag}"
}

nextPatchVersion() {
  local version="$1"
  IFS='.' read -r major minor patch _ <<<"${version}"
  local next_patch="$((patch + 1))"
  printf '%s.%s.%s' "${major}" "${minor}" "${next_patch}"
}

isValid() {
  local version="$1"
  if [[ "${version}" =~ ^(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)(-((0|[1-9][0-9]*|[0-9]*[a-zA-Z-][0-9a-zA-Z-]*)(\.(0|[1-9][0-9]*|[0-9]*[a-zA-Z-][0-9a-zA-Z-]*))*))?(\+([0-9a-zA-Z-]+(\.[0-9a-zA-Z-]+)*))?$ ]]; then
    return 0
  else
    return 1
  fi
}

if [[ "${BASH_SOURCE[0]}" == "$0" ]]; then
  version "$@"
fi
