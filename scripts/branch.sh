#!/usr/bin/env bash
# SPDX-License-Identifier: MIT
set -euo pipefail

for cmd in {git,tr}; do
  if [[ ! -x "$(command -v "${cmd}")" ]]; then
    printf "error: command not found: %s\n" "${cmd}" >&2
    exit 1
  fi
done

# ###################################################################################### #
# Prints the cleaned name of the current Git branch                                      #
#                                                                                        #
# VERSION: 1.1.0                                                                         #
#                                                                                        #
# DEPENDENCIES:                                                                          #
#  - git                                                                                 #
#  - tr                                                                                  #
#                                                                                        #
# EXAMPLES:                                                                              #
#  master => master                                                                      #
#  feature/my-feature => feature-my-feature                                              #
#  poc/Other_Feature => poc-other-feature                                                #
#                                                                                        #
# ###################################################################################### #
branch() {
  local branch_name
  # Try to determine branch name from Git repository
  branch_name="$(git rev-parse --abbrev-ref HEAD)"
  # shellcheck disable=SC2154
  if [[ "${branch_name}" == "HEAD" ]]; then
    # Try to determine branch name from GITHUB_REF environment variable (GitHub Actions)
    if [[ -n "${GITHUB_REF:-}" ]]; then
      branch_name="${GITHUB_REF/refs\/heads\//}"
    # Try to determine branch name from BRANCH_NAME environment variable (Jenkins)
    elif [[ -n "${BRANCH_NAME:-}" ]]; then
      branch_name="${BRANCH_NAME}"
    else
      printf 'error: detached HEAD\n' >&2
      exit 1
    fi
  fi

  # Clean branch name
  branch_name="${branch_name//[^-a-zA-Z0-9]/-}"
  branch_name="$(printf '%s' "${branch_name}" | tr '[:upper:]' '[:lower:]')"

  printf '%s' "${branch_name}"
}

if [[ "${BASH_SOURCE[0]}" == "$0" ]]; then
  branch "$@"
fi
