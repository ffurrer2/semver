// SPDX-License-Identifier: MIT
package semver

import (
	"fmt"
	"reflect"
	"regexp"
	"sort"
	"strings"

	"github.com/ffurrer2/semver/internal/pkg/numeric"
)

const MaxMajor = ^uint(0)
const MaxMinor = ^uint(0)
const MaxPatch = ^uint(0)

const NamedGroupsPattern = `^(?P<major>0|[1-9]\d*)\.(?P<minor>0|[1-9]\d*)\.(?P<patch>0|[1-9]\d*)` +
	`(?:-(?P<prerelease>(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?` +
	`(?:\+(?P<buildmetadata>[0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`

var /* const */ semverRegexp *regexp.Regexp

func init() {
	semverRegexp = regexp.MustCompile(NamedGroupsPattern)
}

type SemVer struct {
	Major         uint     `json:"major"`
	Minor         uint     `json:"minor"`
	Patch         uint     `json:"patch"`
	PreRelease    []string `json:"pre_release"`
	BuildMetadata []string `json:"build_metadata"`
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
		return nil, fmt.Errorf("invalid semantic version: %s", s)
	}
	groupNames := semverRegexp.SubexpNames()
	semver := &SemVer{}
	for groupIdx, group := range matches[0] {
		name := groupNames[groupIdx]
		switch name {
		case "major":
			semver.Major = numeric.MustParseUint(group)
		case "minor":
			semver.Minor = numeric.MustParseUint(group)
		case "patch":
			semver.Patch = numeric.MustParseUint(group)
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

func Sort(data []SemVer) {
	sort.Sort(BySemVer(data))
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

func (s SemVer) NextMajor() *SemVer {
	if len(s.PreRelease) > 0 {
		return &SemVer{
			Major:         s.Major,
			Minor:         s.Minor,
			Patch:         s.Patch,
			PreRelease:    []string{},
			BuildMetadata: []string{},
		}
	}
	if s.Major == MaxMajor {
		panic(fmt.Errorf("next major version overflows uint"))
	}
	return &SemVer{
		Major:         s.Major + 1,
		Minor:         0,
		Patch:         0,
		PreRelease:    []string{},
		BuildMetadata: []string{},
	}
}

func (s SemVer) NextMinor() *SemVer {
	if len(s.PreRelease) > 0 {
		return &SemVer{
			Major:         s.Major,
			Minor:         s.Minor,
			Patch:         s.Patch,
			PreRelease:    []string{},
			BuildMetadata: []string{},
		}
	}
	if s.Minor == MaxMinor {
		panic(fmt.Errorf("next minor version overflows uint"))
	}
	return &SemVer{
		Major:         s.Major,
		Minor:         s.Minor + 1,
		Patch:         0,
		PreRelease:    []string{},
		BuildMetadata: []string{},
	}
}

func (s SemVer) NextPatch() *SemVer {
	if len(s.PreRelease) > 0 {
		return &SemVer{
			Major:         s.Major,
			Minor:         s.Minor,
			Patch:         s.Patch,
			PreRelease:    []string{},
			BuildMetadata: []string{},
		}
	}
	if s.Patch == MaxPatch {
		panic(fmt.Errorf("next patch version overflows uint"))
	}
	return &SemVer{
		Major:         s.Major,
		Minor:         s.Minor,
		Patch:         s.Patch + 1,
		PreRelease:    []string{},
		BuildMetadata: []string{},
	}
}

func (s SemVer) IsValid() bool {
	return semverRegexp.MatchString(s.String())
}

func (s SemVer) CompareTo(o SemVer) int {
	// Major, minor, and patch versions are always compared numerically.
	if res := numeric.CompareUint(s.Major, o.Major); res != 0 {
		return res
	} else if res := numeric.CompareUint(s.Minor, o.Minor); res != 0 {
		return res
	} else if res := numeric.CompareUint(s.Patch, o.Patch); res != 0 {
		return res
	}
	// => Major, minor, and patch are equal and pre-release identifiers are not equal

	// When major, minor, and patch are equal, a pre-release version has lower precedence than a normal version.
	if len(s.PreRelease) == 0 && len(o.PreRelease) == 0 {
		return 0
	} else if len(s.PreRelease) == 0 {
		return 1
	} else if len(o.PreRelease) == 0 {
		return -1
	}

	// Precedence for two pre-release versions with the same major, minor, and patch version MUST be determined by
	// comparing each dot separated identifier from left to right until a difference is found as follows: identifiers
	// consisting of only digits are compared numerically and identifiers with letters or hyphens are compared
	// lexically in ASCII sort order.
	for i := 0; i < numeric.MinInt(len(s.PreRelease), len(o.PreRelease)); i++ {
		if res := comparePreReleaseIdentifier(s.PreRelease[i], o.PreRelease[i]); res != 0 {
			return res
		}
	}

	// A larger set of pre-release fields has a higher precedence than a smaller set, if all of the preceding
	// identifiers are equal.
	if len(s.PreRelease) > len(o.PreRelease) {
		return 1
	} else if len(s.PreRelease) < len(o.PreRelease) {
		return -1
	}

	// => Core versions and pre-release identifiers are equal.
	return 0
}

func (s SemVer) String() string {
	str := fmt.Sprintf("%d.%d.%d", s.Major, s.Minor, s.Patch)
	if len(s.PreRelease) > 0 {
		str = fmt.Sprintf("%s-%s", str, s.PreReleaseString())
	}
	if len(s.BuildMetadata) > 0 {
		str = fmt.Sprintf("%s+%s", str, s.BuildMetadataString())
	}
	return str
}

func (s SemVer) PreReleaseString() string {
	return strings.Join(s.PreRelease, ".")
}

func (s SemVer) BuildMetadataString() string {
	return strings.Join(s.BuildMetadata, ".")
}

func (s SemVer) Equal(o SemVer) bool {
	return reflect.DeepEqual(s, o)
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
	if numeric.IsNumeric(a) && numeric.IsNumeric(b) {
		return numeric.CompareUint(numeric.MustParseUint(a), numeric.MustParseUint(b))
	}
	return strings.Compare(a, b)
}
