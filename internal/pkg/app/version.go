// SPDX-License-Identifier: MIT

package app

import (
	"runtime/debug"
)

const unknown = "unknown"

var (
	version string
	date    string
	commit  string
)

func BuildVersion() string {
	if version != "" {
		return version
	}
	info, ok := debug.ReadBuildInfo()
	if ok {
		return info.Main.Version
	}
	return unknown
}

func BuildDate() string {
	if date != "" {
		return date
	}
	info, ok := debug.ReadBuildInfo()
	if ok {
		for _, v := range info.Settings {
			if v.Key == "vcs.time" {
				return v.Value
			}
		}
	}
	return unknown
}

func CommitHash() string {
	if commit != "" {
		return commit
	}
	info, ok := debug.ReadBuildInfo()
	if ok {
		for _, v := range info.Settings {
			if v.Key == "vcs.revision" {
				return v.Value
			}
		}
	}
	return unknown
}

func TreeState() string {
	info, ok := debug.ReadBuildInfo()
	if ok {
		for _, v := range info.Settings {
			if v.Key == "vcs.modified" {
				if v.Value == "true" {
					return "dirty"
				}
				return "clean"
			}
		}
	}
	return unknown
}
