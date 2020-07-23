// SPDX-License-Identifier: MIT
package app

var (
	version = "0.0.0"
	date    = "1970-01-01T00:00:00Z"
	commit  = "0000000000000000000000000000000000000000"
)

func BuildVersion() string {
	return version
}

func BuildDate() string {
	return date
}

func CommitHash() string {
	return commit
}
