// SPDX-License-Identifier: MIT

package app

import (
	"encoding/json"
	"runtime/debug"

	"github.com/samber/lo"
)

const unknown = "unknown"

var version string

func Version() string {
	if version != "" {
		return version
	}

	info, ok := debug.ReadBuildInfo()
	if !ok {
		return unknown
	}

	return info.Main.Version
}

func CommitHash() string {
	vcsRevision, ok := readBuildSetting("vcs.revision")
	if !ok {
		return unknown
	}

	return vcsRevision
}

func TreeState() string {
	vcsModified, ok := readBuildSetting("vcs.modified")
	if !ok {
		return unknown
	}

	if vcsModified == "true" {
		return "dirty"
	}

	return "clean"
}

func BuildInfoJSON() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "{}"
	}

	marshal, err := json.Marshal(info)
	if err != nil {
		panic(err)
	}

	return string(marshal)
}

func readBuildSetting(key string) (string, bool) {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "", false
	}

	bs, ok := lo.Find(info.Settings, func(v debug.BuildSetting) bool { return v.Key == key })
	if !ok {
		return "", false
	}

	return bs.Value, true
}
