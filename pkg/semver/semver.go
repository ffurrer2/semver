// SPDX-License-Identifier: MIT

package semver

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/ffurrer2/semver/v2/internal/pkg/number"
)

const (
	MaxMajor           = ^uint(0)
	MaxMinor           = ^uint(0)
	MaxPatch           = ^uint(0)
	NamedGroupsPattern = `^(?P<major>0|[1-9]\d*)\.(?P<minor>0|[1-9]\d*)\.(?P<patch>0|[1-9]\d*)` +
		`(?:-(?P<prerelease>(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?` +
		`(?:\+(?P<buildmetadata>[0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`
)

var (
	ErrNextMajorUintOverflow = errors.New("next major version overflows uint")
	ErrNextMinorUintOverflow = errors.New("next minor version overflows uint")
	ErrNextPatchUintOverflow = errors.New("next patch version overflows uint")
	semverRegexp             *regexp.Regexp
)

func init() {
	semverRegexp = regexp.MustCompile(NamedGroupsPattern)
}

type InvalidSemVerError string

func (s InvalidSemVerError) Error() string {
	return "invalid semantic version: " + string(s)
}

type SemVer struct {
	Major         uint     `json:"major"`
	Minor         uint     `json:"minor"`
	Patch         uint     `json:"patch"`
	PreRelease    []string `json:"preRelease"`
	BuildMetadata []string `json:"buildMetadata"`
}

type BySemVer []SemVer

func (s BySemVer) Len() int {
	return len(s)
}

func (s BySemVer) Less(i, j int) bool {
	return s[i].CompareTo(s[j]) == -1
}

func (s BySemVer) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func Parse(s string) (*SemVer, error) {
	matches := semverRegexp.FindAllStringSubmatch(s, -1)
	if matches == nil {
		return nil, InvalidSemVerError(s)
	}
	groupNames := semverRegexp.SubexpNames()
	semver := &SemVer{}
	for groupIdx, group := range matches[0] {
		name := groupNames[groupIdx]
		switch name {
		case "major":
			semver.Major = number.MustParseUint(group)
		case "minor":
			semver.Minor = number.MustParseUint(group)
		case "patch":
			semver.Patch = number.MustParseUint(group)
		case "prerelease":
			semver.PreRelease = splitDotSeparatedString(group)
		case "buildmetadata":
			semver.BuildMetadata = splitDotSeparatedString(group)
		}
	}
	return semver, nil
}

func MustParse(s string) *SemVer {
	semver, err := Parse(s)
	if err != nil {
		panic(err)
	}
	return semver
}

func IsValid(s string) bool {
	return semverRegexp.MatchString(s)
}

func (s *SemVer) SetMajor(major uint) {
	s.Major = major
}

func (s *SemVer) SetMinor(minor uint) {
	s.Minor = minor
}

func (s *SemVer) SetPatch(patch uint) {
	s.Patch = patch
}

func (s *SemVer) SetPreRelease(preRelease []string) {
	s.PreRelease = preRelease
}

func (s *SemVer) SetBuildMetadata(buildMetadata []string) {
	s.BuildMetadata = buildMetadata
}

func (s *SemVer) NextMajor() *SemVer {
	if s.IsPreRelease() {
		return &SemVer{
			Major:         s.Major,
			Minor:         s.Minor,
			Patch:         s.Patch,
			PreRelease:    []string{},
			BuildMetadata: []string{},
		}
	}
	if s.Major == MaxMajor {
		panic(ErrNextMajorUintOverflow)
	}
	return &SemVer{
		Major:         s.Major + 1,
		Minor:         0,
		Patch:         0,
		PreRelease:    []string{},
		BuildMetadata: []string{},
	}
}

func (s *SemVer) NextMinor() *SemVer {
	if s.IsPreRelease() {
		return &SemVer{
			Major:         s.Major,
			Minor:         s.Minor,
			Patch:         s.Patch,
			PreRelease:    []string{},
			BuildMetadata: []string{},
		}
	}
	if s.Minor == MaxMinor {
		panic(ErrNextMinorUintOverflow)
	}
	return &SemVer{
		Major:         s.Major,
		Minor:         s.Minor + 1,
		Patch:         0,
		PreRelease:    []string{},
		BuildMetadata: []string{},
	}
}

func (s *SemVer) NextPatch() *SemVer {
	if s.IsPreRelease() {
		return &SemVer{
			Major:         s.Major,
			Minor:         s.Minor,
			Patch:         s.Patch,
			PreRelease:    []string{},
			BuildMetadata: []string{},
		}
	}
	if s.Patch == MaxPatch {
		panic(ErrNextPatchUintOverflow)
	}
	return &SemVer{
		Major:         s.Major,
		Minor:         s.Minor,
		Patch:         s.Patch + 1,
		PreRelease:    []string{},
		BuildMetadata: []string{},
	}
}

func (s *SemVer) IsValid() bool {
	return semverRegexp.MatchString(s.String())
}

func (s *SemVer) IsRelease() bool {
	return s.IsNotPreRelease() && s.HasNoBuildMetadata()
}

func (s *SemVer) IsNotPreRelease() bool {
	return len(s.PreRelease) == 0
}

func (s *SemVer) IsPreRelease() bool {
	return len(s.PreRelease) > 0
}

func (s *SemVer) HasNoBuildMetadata() bool {
	return len(s.BuildMetadata) == 0
}

func (s *SemVer) HasBuildMetadata() bool {
	return len(s.BuildMetadata) > 0
}

func (s *SemVer) CompareTo(o SemVer) int {
	// Major, minor, and patch versions are always compared numerically.
	if res := number.CompareInt(s.Major, o.Major); res != 0 {
		return res
	} else if res := number.CompareInt(s.Minor, o.Minor); res != 0 {
		return res
	} else if res := number.CompareInt(s.Patch, o.Patch); res != 0 {
		return res
	}
	// => Major, minor, and patch are equal and pre-release identifiers are not equal

	// When major, minor, and patch are equal, a pre-release version has lower precedence than a normal version.
	switch {
	case s.IsNotPreRelease() && o.IsNotPreRelease():
		return 0
	case s.IsNotPreRelease():
		return 1
	case o.IsNotPreRelease():
		return -1
	}

	// Precedence for two pre-release versions with the same major, minor, and patch version MUST be determined by
	// comparing each dot separated identifier from left to right until a difference is found as follows: identifiers
	// consisting of only digits are compared numerically and identifiers with letters or hyphens are compared
	// lexically in ASCII sort order.
	for i := range min(len(s.PreRelease), len(o.PreRelease)) {
		if res := comparePreReleaseIdentifier(s.PreRelease[i], o.PreRelease[i]); res != 0 {
			return res
		}
	}

	// A larger set of pre-release fields has a higher precedence than a smaller set, if all the preceding
	// identifiers are equal.
	if len(s.PreRelease) > len(o.PreRelease) {
		return 1
	} else if len(s.PreRelease) < len(o.PreRelease) {
		return -1
	}

	// => Core versions and pre-release identifiers are equal.
	return 0
}

func (s *SemVer) String() string {
	str := fmt.Sprintf("%d.%d.%d", s.Major, s.Minor, s.Patch)
	if s.IsPreRelease() {
		str = fmt.Sprintf("%s-%s", str, s.PreReleaseString())
	}
	if s.HasBuildMetadata() {
		str = fmt.Sprintf("%s+%s", str, s.BuildMetadataString())
	}
	return str
}

func (s *SemVer) PreReleaseString() string {
	return strings.Join(s.PreRelease, ".")
}

func (s *SemVer) BuildMetadataString() string {
	return strings.Join(s.BuildMetadata, ".")
}

func (s *SemVer) Equal(o SemVer) bool {
	return reflect.DeepEqual(*s, o)
}

func splitDotSeparatedString(s string) []string {
	if s == "" {
		return []string{}
	}
	return strings.Split(s, ".")
}

func comparePreReleaseIdentifier(a, b string) int {
	if a == b {
		return 0
	}
	if number.IsNumeric(a) && number.IsNumeric(b) {
		return number.CompareInt(number.MustParseUint(a), number.MustParseUint(b))
	}
	return strings.Compare(a, b)
}
