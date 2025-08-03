// SPDX-License-Identifier: MIT

package semver_test

import (
	"sort"
	"strconv"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ffurrer2/semver/v2/internal/pkg/number"
	"github.com/ffurrer2/semver/v2/pkg/semver"
)

var _ = Describe("semver:", func() {
	Describe("type InvalidSemVerError string", func() {
		Context("when error message is set", func() {
			It("should return the correct error message", func() {
				err := semver.InvalidSemVerError("0")
				Expect(err.Error()).To(Equal("invalid semantic version: 0"))
			})
		})
	})

	Describe("type SemVer struct", func() {
		It("should implement the sort.Interface interface", func() {
			actual := []semver.SemVer{
				{Major: 1, Minor: 0, Patch: 0, PreRelease: []string{}, BuildMetadata: []string{}},
				{Major: 1, Minor: 1, Patch: 1, PreRelease: []string{}, BuildMetadata: []string{}},
				{Major: 1, Minor: 0, Patch: 1, PreRelease: []string{}, BuildMetadata: []string{}},
				{Major: 1, Minor: 0, Patch: 0, PreRelease: []string{}, BuildMetadata: []string{}},
			}
			sort.Sort(semver.BySemVer(actual))
			expected := []semver.SemVer{
				{Major: 1, Minor: 0, Patch: 0, PreRelease: []string{}, BuildMetadata: []string{}},
				{Major: 1, Minor: 0, Patch: 0, PreRelease: []string{}, BuildMetadata: []string{}},
				{Major: 1, Minor: 0, Patch: 1, PreRelease: []string{}, BuildMetadata: []string{}},
				{Major: 1, Minor: 1, Patch: 1, PreRelease: []string{}, BuildMetadata: []string{}},
			}
			Expect(actual).To(Equal(expected))
		})
	})

	Describe("func Parse(s string) (*SemVer, error)", func() {
		Describe("when input is a valid semantic version", func() {
			DescribeTable("it should not error and return the correct SemVer struct",
				func(input string, major, minor, patch int, preRelease, buildMetadata string) {
					actual, err := semver.Parse(input)
					expected := semver.SemVer{
						Major:         uint(major),
						Minor:         uint(minor),
						Patch:         uint(patch),
						PreRelease:    splitDotSeparatedString(preRelease),
						BuildMetadata: splitDotSeparatedString(buildMetadata),
					}
					Expect(err).ShouldNot(HaveOccurred())
					Expect(actual).To(Equal(&expected))
				},
				Entry("0.0.0", "0.0.0", 0, 0, 0, "", ""),
				Entry("0.0.1", "0.0.1", 0, 0, 1, "", ""),
				Entry("0.1.0", "0.1.0", 0, 1, 0, "", ""),
				Entry("1.0.0", "1.0.0", 1, 0, 0, "", ""),
				Entry("1.0.1", "1.0.1", 1, 0, 1, "", ""),
				Entry("1.1.0", "1.1.0", 1, 1, 0, "", ""),
				Entry("1.1.1", "1.1.1", 1, 1, 1, "", ""),
				Entry("1.10.0", "1.10.0", 1, 10, 0, "", ""),
				Entry("1.0.10", "1.0.10", 1, 0, 10, "", ""),
				Entry("10.0.0", "10.0.0", 10, 0, 0, "", ""),
				Entry("10.10.0", "10.10.0", 10, 10, 0, "", ""),
				Entry("10.10.10", "10.10.10", 10, 10, 10, "", ""),
				Entry("1.0.0-0.3.7", "1.0.0-0.3.7", 1, 0, 0, "0.3.7", ""),
				Entry("1.0.0-0alpha1", "1.0.0-0alpha1", 1, 0, 0, "0alpha1", ""),
				Entry("1.0.0-alpha.1", "1.0.0-alpha.1", 1, 0, 0, "alpha.1", ""),
				Entry("1.0.0-alpha.12", "1.0.0-alpha.12", 1, 0, 0, "alpha.12", ""),
				Entry("1.0.0-alpha.beta", "1.0.0-alpha.beta", 1, 0, 0, "alpha.beta", ""),
				Entry("1.0.0-alpha", "1.0.0-alpha", 1, 0, 0, "alpha", ""),
				Entry("1.0.0-alpha+001", "1.0.0-alpha+001", 1, 0, 0, "alpha", "001"),
				Entry("1.0.0-beta.11", "1.0.0-beta.11", 1, 0, 0, "beta.11", ""),
				Entry("1.0.0-beta.2", "1.0.0-beta.2", 1, 0, 0, "beta.2", ""),
				Entry("1.0.0-beta.511485", "1.0.0-beta.511485", 1, 0, 0, "beta.511485", ""),
				Entry("1.0.0-beta.5114fa", "1.0.0-beta.5114fa", 1, 0, 0, "beta.5114fa", ""),
				Entry("1.0.0-beta", "1.0.0-beta", 1, 0, 0, "beta", ""),
				Entry("1.0.0-beta+exp.sha.5114f85", "1.0.0-beta+exp.sha.5114f85", 1, 0, 0, "beta", "exp.sha.5114f85"),
				Entry("1.0.0-rc.1", "1.0.0-rc.1", 1, 0, 0, "rc.1", ""),
				Entry("1.0.0-x.7.z.92", "1.0.0-x.7.z.92", 1, 0, 0, "x.7.z.92", ""),
				Entry("1.0.0+20130313144700", "1.0.0+20130313144700", 1, 0, 0, "", "20130313144700"),
			)
		})

		Describe("when input is not a valid semantic version", func() {
			DescribeTable("it should error",
				func(input string) {
					actual, err := semver.Parse(input)
					Expect(err).To(HaveOccurred())
					Expect(actual).To(BeNil())
				},
				Entry("empty word", ""),
				Entry("abc", "abc"),
				Entry("0", "0"),
				Entry("0.0", "0.0"),
				Entry("0.0.", "0.0."),
				Entry(".0.0", ".0.0"),
				Entry("0.0.a", "0.0.a"),
				Entry("00.0.0", "00.0.0"),
				Entry("0.00.0", "0.00.0"),
				Entry("0.0.00", "0.0.00"),
				Entry("0.0.0-", "0.0.0-"),
				Entry("0.0.0-alpha+", "0.0.0-alpha+"),
			)
		})
	})

	Describe("func MustParse(s string) *SemVer", func() {
		Describe("when input is a valid semantic version", func() {
			DescribeTable("it should not panic and return the correct SemVer struct",
				func(input string, major, minor, patch int, preRelease, buildMetadata string) {
					actual := semver.MustParse(input)
					expected := semver.SemVer{
						Major:         uint(major),
						Minor:         uint(minor),
						Patch:         uint(patch),
						PreRelease:    splitDotSeparatedString(preRelease),
						BuildMetadata: splitDotSeparatedString(buildMetadata),
					}
					Expect(actual).To(Equal(&expected))
				},
				Entry("1.0.0", "1.0.0", 1, 0, 0, "", ""),
				Entry("1.0.0-alpha.1", "1.0.0-alpha.1", 1, 0, 0, "alpha.1", ""),
				Entry("1.0.0-rc.1", "1.0.0-rc.1", 1, 0, 0, "rc.1", ""),
			)
		})

		Describe("when input is not a valid semantic version", func() {
			DescribeTable("it should panic",
				func(input string) {
					Expect(func() { semver.MustParse(input) }).Should(Panic())
				},
				Entry("empty word", ""),
				Entry("abc", "abc"),
				Entry("0", "0"),
			)
		})
	})

	Describe("func IsValid(s string) bool", func() {
		Describe("when input is a valid semantic version", func() {
			DescribeTable("it should return true",
				func(input string) {
					actual := semver.IsValid(input)
					Expect(actual).To(BeTrue())
				},
				Entry("0.0.0", "0.0.0"),
				Entry("0.0.1", "0.0.1"),
				Entry("0.1.0", "0.1.0"),
				Entry("1.0.0", "1.0.0"),
				Entry("1.0.1", "1.0.1"),
				Entry("1.1.0", "1.1.0"),
				Entry("1.1.1", "1.1.1"),
				Entry("1.10.0", "1.10.0"),
				Entry("1.0.10", "1.0.10"),
				Entry("10.0.0", "10.0.0"),
				Entry("10.10.0", "10.10.0"),
				Entry("10.10.10", "10.10.10"),
				Entry("1.0.0-0.3.7", "1.0.0-0.3.7"),
				Entry("1.0.0-0alpha1", "1.0.0-0alpha1"),
				Entry("1.0.0-alpha.1", "1.0.0-alpha.1"),
				Entry("1.0.0-alpha.12", "1.0.0-alpha.12"),
				Entry("1.0.0-alpha.beta", "1.0.0-alpha.beta"),
				Entry("1.0.0-alpha", "1.0.0-alpha"),
				Entry("1.0.0-alpha+001", "1.0.0-alpha+001"),
				Entry("1.0.0-beta.11", "1.0.0-beta.11"),
				Entry("1.0.0-beta.2", "1.0.0-beta.2"),
				Entry("1.0.0-beta.511485", "1.0.0-beta.511485"),
				Entry("1.0.0-beta.5114fa", "1.0.0-beta.5114fa"),
				Entry("1.0.0-beta", "1.0.0-beta"),
				Entry("1.0.0-beta+exp.sha.5114f85", "1.0.0-beta+exp.sha.5114f85"),
				Entry("1.0.0-rc.1", "1.0.0-rc.1"),
				Entry("1.0.0-x.7.z.92", "1.0.0-x.7.z.92"),
				Entry("1.0.0+20130313144700", "1.0.0+20130313144700"),
			)
		})

		Describe("when input is not a valid semantic version", func() {
			DescribeTable("it should return false",
				func(input string) {
					actual := semver.IsValid(input)
					Expect(actual).To(BeFalse())
				},
				Entry("empty word", ""),
				Entry("abc", "abc"),
				Entry("0", "0"),
				Entry("0.0", "0.0"),
				Entry("0.0.", "0.0."),
				Entry(".0.0", ".0.0"),
				Entry("0.0.a", "0.0.a"),
				Entry("00.0.0", "00.0.0"),
				Entry("0.00.0", "0.00.0"),
				Entry("0.0.00", "0.0.00"),
				Entry("0.0.0-", "0.0.0-"),
				Entry("0.0.0-alpha+", "0.0.0-alpha+"),
			)
		})
	})

	Describe("func (s *SemVer) SetMajor(major uint)", func() {
		It("should set the major version correctly", func() {
			sut := semver.SemVer{
				Major:         0,
				Minor:         0,
				Patch:         0,
				PreRelease:    []string{},
				BuildMetadata: []string{},
			}
			sut.SetMajor(1)
			Expect(sut.Major).To(Equal(uint(1)))
			Expect(sut.Minor).To(Equal(uint(0)))
			Expect(sut.Patch).To(Equal(uint(0)))
			Expect(sut.PreRelease).To(Equal([]string{}))
			Expect(sut.BuildMetadata).To(Equal([]string{}))
		})
	})

	Describe("func (s *SemVer) SetMinor(minor uint)", func() {
		It("should set the minor version correctly", func() {
			sut := semver.SemVer{
				Major:         0,
				Minor:         0,
				Patch:         0,
				PreRelease:    []string{},
				BuildMetadata: []string{},
			}
			sut.SetMinor(1)
			Expect(sut.Major).To(Equal(uint(0)))
			Expect(sut.Minor).To(Equal(uint(1)))
			Expect(sut.Patch).To(Equal(uint(0)))
			Expect(sut.PreRelease).To(Equal([]string{}))
			Expect(sut.BuildMetadata).To(Equal([]string{}))
		})
	})

	Describe("func (s *SemVer) SetPatch(patch uint)", func() {
		It("should set the patch version correctly", func() {
			sut := semver.SemVer{
				Major:         0,
				Minor:         0,
				Patch:         0,
				PreRelease:    []string{},
				BuildMetadata: []string{},
			}
			sut.SetPatch(1)
			Expect(sut.Major).To(Equal(uint(0)))
			Expect(sut.Minor).To(Equal(uint(0)))
			Expect(sut.Patch).To(Equal(uint(1)))
			Expect(sut.PreRelease).To(Equal([]string{}))
			Expect(sut.BuildMetadata).To(Equal([]string{}))
		})
	})

	Describe("func (s *SemVer) SetPreRelease(preRelease []string)", func() {
		It("should set the preRelease identifiers correctly", func() {
			sut := semver.SemVer{
				Major:         0,
				Minor:         0,
				Patch:         0,
				PreRelease:    []string{},
				BuildMetadata: []string{},
			}
			sut.SetPreRelease([]string{"alpha", "1"})
			Expect(sut.Major).To(Equal(uint(0)))
			Expect(sut.Minor).To(Equal(uint(0)))
			Expect(sut.Patch).To(Equal(uint(0)))
			Expect(sut.PreRelease).To(Equal([]string{"alpha", "1"}))
			Expect(sut.BuildMetadata).To(Equal([]string{}))
		})
	})

	Describe("func (s *SemVer) SetBuildMetadata(buildMetadata []string)", func() {
		It("should set the buildMetadata identifiers correctly", func() {
			sut := semver.SemVer{
				Major:         0,
				Minor:         0,
				Patch:         0,
				PreRelease:    []string{},
				BuildMetadata: []string{},
			}
			sut.SetBuildMetadata([]string{"2020.01.01", "001"})
			Expect(sut.Major).To(Equal(uint(0)))
			Expect(sut.Minor).To(Equal(uint(0)))
			Expect(sut.Patch).To(Equal(uint(0)))
			Expect(sut.PreRelease).To(Equal([]string{}))
			Expect(sut.BuildMetadata).To(Equal([]string{"2020.01.01", "001"}))
		})
	})

	Describe("func (s SemVer) NextMajor() *SemVer", func() {
		DescribeTable("it should return the next major semantic version",
			func(input, expected string) {
				sut, err := semver.Parse(input)
				Expect(err).ShouldNot(HaveOccurred())
				actual := sut.NextMajor().String()
				Expect(actual).To(Equal(expected))
			},
			Entry("0.0.0", "0.0.0", "1.0.0"),
			Entry("0.0.1", "0.0.1", "1.0.0"),
			Entry("0.1.0", "0.1.0", "1.0.0"),
			Entry("1.0.0", "1.0.0", "2.0.0"),
			Entry("1.0.1", "1.0.1", "2.0.0"),
			Entry("1.1.0", "1.1.0", "2.0.0"),
			Entry("1.1.1", "1.1.1", "2.0.0"),
			Entry("1.10.0", "1.10.0", "2.0.0"),
			Entry("1.0.10", "1.0.10", "2.0.0"),
			Entry("10.0.0", "10.0.0", "11.0.0"),
			Entry("10.10.0", "10.10.0", "11.0.0"),
			Entry("10.10.10", "10.10.10", "11.0.0"),
			Entry("1.0.0-0.3.7", "1.0.0-0.3.7", "1.0.0"),
			Entry("1.0.0-0alpha1", "1.0.0-0alpha1", "1.0.0"),
			Entry("1.0.0-alpha.1", "1.0.0-alpha.1", "1.0.0"),
			Entry("1.0.0-alpha.12", "1.0.0-alpha.12", "1.0.0"),
			Entry("1.0.0-alpha.beta", "1.0.0-alpha.beta", "1.0.0"),
			Entry("1.0.0-alpha", "1.0.0-alpha", "1.0.0"),
			Entry("1.0.0-alpha+001", "1.0.0-alpha+001", "1.0.0"),
			Entry("1.0.0-beta.11", "1.0.0-beta.11", "1.0.0"),
			Entry("1.0.0-beta.2", "1.0.0-beta.2", "1.0.0"),
			Entry("1.0.0-beta.511485", "1.0.0-beta.511485", "1.0.0"),
			Entry("1.0.0-beta.5114fa", "1.0.0-beta.5114fa", "1.0.0"),
			Entry("1.0.0-beta", "1.0.0-beta", "1.0.0"),
			Entry("1.0.0-beta+exp.sha.5114f85", "1.0.0-beta+exp.sha.5114f85", "1.0.0"),
			Entry("1.0.0-rc.1", "1.0.0-rc.1", "1.0.0"),
			Entry("1.0.0-x.7.z.92", "1.0.0-x.7.z.92", "1.0.0"),
			Entry("1.0.0+20130313144700", "1.0.0+20130313144700", "2.0.0"),
		)

		Context("when s.Major is equal to MaxMajor", func() {
			It("should panic", func() {
				semver := semver.SemVer{
					Major:         semver.MaxMajor,
					Minor:         0,
					Patch:         0,
					PreRelease:    []string{},
					BuildMetadata: []string{},
				}
				sut := func() {
					semver.NextMajor()
				}
				Expect(sut).To(Panic())
			})
		})
	})

	Describe("func (s SemVer) NextMinor() *SemVer", func() {
		DescribeTable("it should return the next minor semantic version",
			func(input, expected string) {
				sut, err := semver.Parse(input)
				Expect(err).ShouldNot(HaveOccurred())
				actual := sut.NextMinor().String()
				Expect(actual).To(Equal(expected))
			},
			Entry("0.0.0", "0.0.0", "0.1.0"),
			Entry("0.0.1", "0.0.1", "0.1.0"),
			Entry("0.1.0", "0.1.0", "0.2.0"),
			Entry("1.0.0", "1.0.0", "1.1.0"),
			Entry("1.0.1", "1.0.1", "1.1.0"),
			Entry("1.1.0", "1.1.0", "1.2.0"),
			Entry("1.1.1", "1.1.1", "1.2.0"),
			Entry("1.10.0", "1.10.0", "1.11.0"),
			Entry("1.0.10", "1.0.10", "1.1.0"),
			Entry("10.0.0", "10.0.0", "10.1.0"),
			Entry("10.10.0", "10.10.0", "10.11.0"),
			Entry("10.10.10", "10.10.10", "10.11.0"),
			Entry("1.0.0-0.3.7", "1.0.0-0.3.7", "1.0.0"),
			Entry("1.0.0-0alpha1", "1.0.0-0alpha1", "1.0.0"),
			Entry("1.0.0-alpha.1", "1.0.0-alpha.1", "1.0.0"),
			Entry("1.0.0-alpha.12", "1.0.0-alpha.12", "1.0.0"),
			Entry("1.0.0-alpha.beta", "1.0.0-alpha.beta", "1.0.0"),
			Entry("1.0.0-alpha", "1.0.0-alpha", "1.0.0"),
			Entry("1.0.0-alpha+001", "1.0.0-alpha+001", "1.0.0"),
			Entry("1.0.0-beta.11", "1.0.0-beta.11", "1.0.0"),
			Entry("1.0.0-beta.2", "1.0.0-beta.2", "1.0.0"),
			Entry("1.0.0-beta.511485", "1.0.0-beta.511485", "1.0.0"),
			Entry("1.0.0-beta.5114fa", "1.0.0-beta.5114fa", "1.0.0"),
			Entry("1.0.0-beta", "1.0.0-beta", "1.0.0"),
			Entry("1.0.0-beta+exp.sha.5114f85", "1.0.0-beta+exp.sha.5114f85", "1.0.0"),
			Entry("1.0.0-rc.1", "1.0.0-rc.1", "1.0.0"),
			Entry("1.0.0-x.7.z.92", "1.0.0-x.7.z.92", "1.0.0"),
			Entry("1.0.0+20130313144700", "1.0.0+20130313144700", "1.1.0"),
		)

		Context("when s.Minor is equal to MaxMinor", func() {
			It("should panic", func() {
				semver := semver.SemVer{
					Major:         0,
					Minor:         semver.MaxMinor,
					Patch:         0,
					PreRelease:    []string{},
					BuildMetadata: []string{},
				}
				sut := func() {
					semver.NextMinor()
				}
				Expect(sut).To(Panic())
			})
		})
	})

	Describe("func (s SemVer) NextPatch() *SemVer", func() {
		DescribeTable("it should return the next patch semantic version",
			func(input, expected string) {
				sut, err := semver.Parse(input)
				Expect(err).ShouldNot(HaveOccurred())
				actual := sut.NextPatch().String()
				Expect(actual).To(Equal(expected))
			},
			Entry("0.0.0", "0.0.0", "0.0.1"),
			Entry("0.0.1", "0.0.1", "0.0.2"),
			Entry("0.1.0", "0.1.0", "0.1.1"),
			Entry("1.0.0", "1.0.0", "1.0.1"),
			Entry("1.0.1", "1.0.1", "1.0.2"),
			Entry("1.1.0", "1.1.0", "1.1.1"),
			Entry("1.1.1", "1.1.1", "1.1.2"),
			Entry("1.10.0", "1.10.0", "1.10.1"),
			Entry("1.0.10", "1.0.10", "1.0.11"),
			Entry("10.0.0", "10.0.0", "10.0.1"),
			Entry("10.10.0", "10.10.0", "10.10.1"),
			Entry("10.10.10", "10.10.10", "10.10.11"),
			Entry("1.0.0-0.3.7", "1.0.0-0.3.7", "1.0.0"),
			Entry("1.0.0-0alpha1", "1.0.0-0alpha1", "1.0.0"),
			Entry("1.0.0-alpha.1", "1.0.0-alpha.1", "1.0.0"),
			Entry("1.0.0-alpha.12", "1.0.0-alpha.12", "1.0.0"),
			Entry("1.0.0-alpha.beta", "1.0.0-alpha.beta", "1.0.0"),
			Entry("1.0.0-alpha", "1.0.0-alpha", "1.0.0"),
			Entry("1.0.0-alpha+001", "1.0.0-alpha+001", "1.0.0"),
			Entry("1.0.0-beta.11", "1.0.0-beta.11", "1.0.0"),
			Entry("1.0.0-beta.2", "1.0.0-beta.2", "1.0.0"),
			Entry("1.0.0-beta.511485", "1.0.0-beta.511485", "1.0.0"),
			Entry("1.0.0-beta.5114fa", "1.0.0-beta.5114fa", "1.0.0"),
			Entry("1.0.0-beta", "1.0.0-beta", "1.0.0"),
			Entry("1.0.0-beta+exp.sha.5114f85", "1.0.0-beta+exp.sha.5114f85", "1.0.0"),
			Entry("1.0.0-rc.1", "1.0.0-rc.1", "1.0.0"),
			Entry("1.0.0-x.7.z.92", "1.0.0-x.7.z.92", "1.0.0"),
			Entry("1.0.0+20130313144700", "1.0.0+20130313144700", "1.0.1"),
		)

		Context("when s.Patch is equal to MaxPatch", func() {
			It("should panic", func() {
				semver := semver.SemVer{
					Major:         0,
					Minor:         0,
					Patch:         semver.MaxPatch,
					PreRelease:    []string{},
					BuildMetadata: []string{},
				}
				sut := func() {
					semver.NextPatch()
				}
				Expect(sut).To(Panic())
			})
		})
	})

	Describe("func (s SemVer) IsValid() bool", func() {
		Describe("when s is a valid semantic version", func() {
			DescribeTable("it should return true",
				func(input string) {
					sut, err := semver.Parse(input)
					Expect(err).ShouldNot(HaveOccurred())
					actual := sut.IsValid()
					Expect(actual).To(BeTrue())
				},
				Entry("0.0.0", "0.0.0"),
				Entry("0.0.1", "0.0.1"),
				Entry("0.1.0", "0.1.0"),
				Entry("1.0.0", "1.0.0"),
				Entry("1.0.1", "1.0.1"),
				Entry("1.1.0", "1.1.0"),
				Entry("1.1.1", "1.1.1"),
				Entry("1.10.0", "1.10.0"),
				Entry("1.0.10", "1.0.10"),
				Entry("10.0.0", "10.0.0"),
				Entry("10.10.0", "10.10.0"),
				Entry("10.10.10", "10.10.10"),
				Entry("1.0.0-0.3.7", "1.0.0-0.3.7"),
				Entry("1.0.0-0alpha1", "1.0.0-0alpha1"),
				Entry("1.0.0-alpha.1", "1.0.0-alpha.1"),
				Entry("1.0.0-alpha.12", "1.0.0-alpha.12"),
				Entry("1.0.0-alpha.beta", "1.0.0-alpha.beta"),
				Entry("1.0.0-alpha", "1.0.0-alpha"),
				Entry("1.0.0-alpha+001", "1.0.0-alpha+001"),
				Entry("1.0.0-beta.11", "1.0.0-beta.11"),
				Entry("1.0.0-beta.2", "1.0.0-beta.2"),
				Entry("1.0.0-beta.511485", "1.0.0-beta.511485"),
				Entry("1.0.0-beta.5114fa", "1.0.0-beta.5114fa"),
				Entry("1.0.0-beta", "1.0.0-beta"),
				Entry("1.0.0-beta+exp.sha.5114f85", "1.0.0-beta+exp.sha.5114f85"),
				Entry("1.0.0-rc.1", "1.0.0-rc.1"),
				Entry("1.0.0-x.7.z.92", "1.0.0-x.7.z.92"),
				Entry("1.0.0+20130313144700", "1.0.0+20130313144700"),
			)
		})

		Describe("when s is not a valid semantic version", func() {
			It("should return false", func() {
				sut := semver.SemVer{
					Major:         0,
					Minor:         0,
					Patch:         0,
					PreRelease:    []string{"00"},
					BuildMetadata: []string{},
				}
				actual := sut.IsValid()
				Expect(actual).To(BeFalse())
			})
		})
	})

	Describe("func (s *SemVer) IsRelease() bool", func() {
		Describe("when s is valid a release version", func() {
			DescribeTable("should return true",
				func(input string) {
					sut, err := semver.Parse(input)
					Expect(err).ShouldNot(HaveOccurred())
					actual := sut.IsRelease()
					Expect(actual).To(BeTrue())
				},
				Entry("0.0.0", "0.0.0"),
				Entry("0.0.1", "0.0.1"),
				Entry("0.1.0", "0.1.0"),
				Entry("1.0.0", "1.0.0"),
				Entry("1.0.1", "1.0.1"),
				Entry("1.1.0", "1.1.0"),
				Entry("1.1.1", "1.1.1"),
				Entry("1.10.0", "1.10.0"),
				Entry("1.0.10", "1.0.10"),
				Entry("10.0.0", "10.0.0"),
				Entry("10.10.0", "10.10.0"),
				Entry("10.10.10", "10.10.10"),
			)
		})
		Describe("when s is valid not a release version", func() {
			DescribeTable("should return true",
				func(input string) {
					sut, err := semver.Parse(input)
					Expect(err).ShouldNot(HaveOccurred())
					actual := sut.IsRelease()
					Expect(actual).To(BeFalse())
				},
				Entry("1.0.0-0.3.7", "1.0.0-0.3.7"),
				Entry("1.0.0-0alpha1", "1.0.0-0alpha1"),
				Entry("1.0.0-alpha.1", "1.0.0-alpha.1"),
				Entry("1.0.0-alpha.12", "1.0.0-alpha.12"),
				Entry("1.0.0-alpha.beta", "1.0.0-alpha.beta"),
				Entry("1.0.0-alpha", "1.0.0-alpha"),
				Entry("1.0.0-alpha+001", "1.0.0-alpha+001"),
				Entry("1.0.0-beta.11", "1.0.0-beta.11"),
				Entry("1.0.0-beta.2", "1.0.0-beta.2"),
				Entry("1.0.0-beta.511485", "1.0.0-beta.511485"),
				Entry("1.0.0-beta.5114fa", "1.0.0-beta.5114fa"),
				Entry("1.0.0-beta", "1.0.0-beta"),
				Entry("1.0.0-beta+exp.sha.5114f85", "1.0.0-beta+exp.sha.5114f85"),
				Entry("1.0.0-rc.1", "1.0.0-rc.1"),
				Entry("1.0.0-x.7.z.92", "1.0.0-x.7.z.92"),
				Entry("1.0.0+20130313144700", "1.0.0+20130313144700"),
			)
		})
	})

	Describe("func (s SemVer) CompareTo(o SemVer) int", func() {
		Describe("when s != o", func() {
			DescribeTable("it should return -1",
				func(a, b string) {
					semverA, err := semver.Parse(a)
					Expect(err).ShouldNot(HaveOccurred())
					semverB, err := semver.Parse(b)
					Expect(err).ToNot(HaveOccurred())
					actual := semverA.CompareTo(*semverB)
					Expect(actual).To(Equal(-1))
				},
				Entry("1.0.0 != 2.0.0", "1.0.0", "2.0.0"),
				Entry("2.0.0 != 2.1.0", "2.0.0", "2.1.0"),
				Entry("2.1.0 != 2.1.1", "2.1.0", "2.1.1"),
				Entry("1.0.0-alpha != 1.0.0", "1.0.0-alpha", "1.0.0"),
				Entry("1.0.0-alpha != 1.0.0-alpha.1", "1.0.0-alpha", "1.0.0-alpha.1"),
				Entry("1.0.0-alpha.1 != 1.0.0-alpha.beta", "1.0.0-alpha.1", "1.0.0-alpha.beta"),
				Entry("1.0.0-alpha.beta != 1.0.0-beta", "1.0.0-alpha.beta", "1.0.0-beta"),
				Entry("1.0.0-beta != 1.0.0-beta.2", "1.0.0-beta", "1.0.0-beta.2"),
				Entry("1.0.0-beta.2 != 1.0.0-beta.11", "1.0.0-beta.2", "1.0.0-beta.11"),
				Entry("1.0.0-beta.11 != 1.0.0-rc.1", "1.0.0-beta.11", "1.0.0-rc.1"),
				Entry("1.0.0-rc.1 != 1.0.0", "1.0.0-rc.1", "1.0.0"),
			)
		})

		Describe("when s > o", func() {
			DescribeTable("it should return 1",
				func(a, b string) {
					semverA, err := semver.Parse(a)
					Expect(err).ShouldNot(HaveOccurred())
					semverB, err := semver.Parse(b)
					Expect(err).ToNot(HaveOccurred())
					actual := semverA.CompareTo(*semverB)
					Expect(actual).To(Equal(1))
				},
				Entry("2.0.0 > 1.0.0", "2.0.0", "1.0.0"),
				Entry("2.1.0 > 2.0.0", "2.1.0", "2.0.0"),
				Entry("2.1.1 > 2.1.0", "2.1.1", "2.1.0"),
				Entry("1.0.0 > 1.0.0-alpha", "1.0.0", "1.0.0-alpha"),
				Entry("1.0.0-alpha.1 > 1.0.0-alpha", "1.0.0-alpha.1", "1.0.0-alpha"),
				Entry("1.0.0-alpha.beta > 1.0.0-alpha.1", "1.0.0-alpha.beta", "1.0.0-alpha.1"),
				Entry("1.0.0-beta > 1.0.0-alpha.beta", "1.0.0-beta", "1.0.0-alpha.beta"),
				Entry("1.0.0-beta.2 > 1.0.0-beta", "1.0.0-beta.2", "1.0.0-beta"),
				Entry("1.0.0-beta.11 > 1.0.0-beta.2", "1.0.0-beta.11", "1.0.0-beta.2"),
				Entry("1.0.0-rc.1 > 1.0.0-beta.11", "1.0.0-rc.1", "1.0.0-beta.11"),
				Entry("1.0.0 > 1.0.0-rc.1", "1.0.0", "1.0.0-rc.1"),
			)
		})

		Describe("when s == o", func() {
			DescribeTable("it should return 0",
				func(a, b string) {
					semverA, err := semver.Parse(a)
					Expect(err).ShouldNot(HaveOccurred())
					semverB, err := semver.Parse(b)
					Expect(err).ShouldNot(HaveOccurred())
					actual := semverA.CompareTo(*semverB)
					Expect(actual).To(Equal(0))
				},
				Entry("0.0.0 == 0.0.0", "0.0.0", "0.0.0"),
				Entry("1.0.0 == 1.0.0", "1.0.0", "1.0.0"),
				Entry("1.1.0 == 1.1.0", "1.1.0", "1.1.0"),
				Entry("1.1.1 == 1.1.1", "1.1.1", "1.1.1"),
				Entry("1.0.0-alpha == 1.0.0-alpha", "1.0.0-alpha", "1.0.0-alpha"),
				Entry("1.0.0-rc.1 == 1.0.0-rc.1", "1.0.0-rc.1", "1.0.0-rc.1"),
				Entry("1.0.0-rc.1+123 == 1.0.0-rc.1+123", "1.0.0-rc.1+123", "1.0.0-rc.1+123"),
				Entry("1.0.0-rc.1+123 == 1.0.0-rc.1+42", "1.0.0-rc.1+123", "1.0.0-rc.1+42"),
				Entry("1.0.0+123 == 1.0.0+123", "1.0.0+123", "1.0.0+123"),
				Entry("1.0.0+123 == 1.0.0+42", "1.0.0+123", "1.0.0+42"),
			)
		})
	})

	Describe("func (s SemVer) Equal(o SemVer) bool", func() {
		Describe("when s is equal to o", func() {
			DescribeTable("it should return true",
				func(a, b string) {
					semverA, err := semver.Parse(a)
					Expect(err).ShouldNot(HaveOccurred())
					semverB, err := semver.Parse(b)
					Expect(err).ShouldNot(HaveOccurred())
					actual := semverA.Equal(*semverB)
					Expect(actual).To(BeTrue())
				},
				Entry("0.0.0 == 0.0.0", "0.0.0", "0.0.0"),
				Entry("1.0.0 == 1.0.0", "1.0.0", "1.0.0"),
				Entry("1.1.0 == 1.1.0", "1.1.0", "1.1.0"),
				Entry("1.1.1 == 1.1.1", "1.1.1", "1.1.1"),
				Entry("1.0.0-alpha == 1.0.0-alpha", "1.0.0-alpha", "1.0.0-alpha"),
				Entry("1.0.0-rc.1 == 1.0.0-rc.1", "1.0.0-rc.1", "1.0.0-rc.1"),
				Entry("1.0.0-rc.1+123 == 1.0.0-rc.1+123", "1.0.0-rc.1+123", "1.0.0-rc.1+123"),
				Entry("1.0.0+123 == 1.0.0+123", "1.0.0+123", "1.0.0+123"),
				Entry("1.0.0+1.2.3 == 1.0.0+1.2.3", "1.0.0+1.2.3", "1.0.0+1.2.3"),
			)
		})

		Describe("when s is not equal to o", func() {
			DescribeTable("it should return false",
				func(a, b string) {
					semverA, err := semver.Parse(a)
					Expect(err).ShouldNot(HaveOccurred())
					semverB, err := semver.Parse(b)
					Expect(err).ShouldNot(HaveOccurred())
					actual := semverA.Equal(*semverB)
					Expect(actual).To(BeFalse())
				},
				Entry("1.0.0 != 2.0.0", "1.0.0", "2.0.0"),
				Entry("2.0.0 != 2.1.0", "2.0.0", "2.1.0"),
				Entry("2.1.0 != 2.1.1", "2.1.0", "2.1.1"),
				Entry("1.0.0-alpha != 1.0.0", "1.0.0-alpha", "1.0.0"),
				Entry("1.0.0-alpha != 1.0.0-alpha.1", "1.0.0-alpha", "1.0.0-alpha.1"),
				Entry("1.0.0-alpha.1 != 1.0.0-alpha.beta", "1.0.0-alpha.1", "1.0.0-alpha.beta"),
				Entry("1.0.0-alpha.beta != 1.0.0-beta", "1.0.0-alpha.beta", "1.0.0-beta"),
				Entry("1.0.0-beta != 1.0.0-beta.2", "1.0.0-beta", "1.0.0-beta.2"),
				Entry("1.0.0-beta.2 != 1.0.0-beta.11", "1.0.0-beta.2", "1.0.0-beta.11"),
				Entry("1.0.0-beta.11 != 1.0.0-rc.1", "1.0.0-beta.11", "1.0.0-rc.1"),
				Entry("1.0.0-rc.1 != 1.0.0", "1.0.0-rc.1", "1.0.0"),
			)
		})
	})

	Describe("func (s SemVer) String() string", func() {
		Describe("when s is a valid SemVer struct", func() {
			DescribeTable("it should return the correct semver string",
				func(major, minor, patch int, preRelease, buildMetadata, expected string) {
					s := semver.SemVer{
						Major:         uint(major),
						Minor:         uint(minor),
						Patch:         uint(patch),
						PreRelease:    splitDotSeparatedString(preRelease),
						BuildMetadata: splitDotSeparatedString(buildMetadata),
					}
					actual := s.String()
					Expect(actual).To(Equal(expected))
				},
				Entry("0.0.0", 0, 0, 0, "", "", "0.0.0"),
				Entry("0.0.1", 0, 0, 1, "", "", "0.0.1"),
				Entry("0.1.0", 0, 1, 0, "", "", "0.1.0"),
				Entry("1.0.0", 1, 0, 0, "", "", "1.0.0"),
				Entry("1.0.1", 1, 0, 1, "", "", "1.0.1"),
				Entry("1.1.0", 1, 1, 0, "", "", "1.1.0"),
				Entry("1.1.1", 1, 1, 1, "", "", "1.1.1"),
				Entry("1.10.0", 1, 10, 0, "", "", "1.10.0"),
				Entry("1.0.10", 1, 0, 10, "", "", "1.0.10"),
				Entry("10.0.0", 10, 0, 0, "", "", "10.0.0"),
				Entry("10.10.0", 10, 10, 0, "", "", "10.10.0"),
				Entry("10.10.10", 10, 10, 10, "", "", "10.10.10"),
				Entry("1.0.0-0.3.7", 1, 0, 0, "0.3.7", "", "1.0.0-0.3.7"),
				Entry("1.0.0-0alpha1", 1, 0, 0, "0alpha1", "", "1.0.0-0alpha1"),
				Entry("1.0.0-alpha.1", 1, 0, 0, "alpha.1", "", "1.0.0-alpha.1"),
				Entry("1.0.0-alpha.12", 1, 0, 0, "alpha.12", "", "1.0.0-alpha.12"),
				Entry("1.0.0-alpha.beta", 1, 0, 0, "alpha.beta", "", "1.0.0-alpha.beta"),
				Entry("1.0.0-alpha", 1, 0, 0, "alpha", "", "1.0.0-alpha"),
				Entry("1.0.0-alpha+001", 1, 0, 0, "alpha", "001", "1.0.0-alpha+001"),
				Entry("1.0.0-beta.11", 1, 0, 0, "beta.11", "", "1.0.0-beta.11"),
				Entry("1.0.0-beta.2", 1, 0, 0, "beta.2", "", "1.0.0-beta.2"),
				Entry("1.0.0-beta.511485", 1, 0, 0, "beta.511485", "", "1.0.0-beta.511485"),
				Entry("1.0.0-beta.5114fa", 1, 0, 0, "beta.5114fa", "", "1.0.0-beta.5114fa"),
				Entry("1.0.0-beta", 1, 0, 0, "beta", "", "1.0.0-beta"),
				Entry("1.0.0-beta+exp.sha.5114f85", 1, 0, 0, "beta", "exp.sha.5114f85", "1.0.0-beta+exp.sha.5114f85"),
				Entry("1.0.0-rc.1", 1, 0, 0, "rc.1", "", "1.0.0-rc.1"),
				Entry("1.0.0-x.7.z.92", 1, 0, 0, "x.7.z.92", "", "1.0.0-x.7.z.92"),
				Entry("1.0.0+20130313144700", 1, 0, 0, "", "20130313144700", "1.0.0+20130313144700"),
			)
		})
	})

	Describe("func mustParseUint(s string) uint", func() {
		Context("when s is a valid unsigned integer", func() {
			DescribeTable("it should return the correct uint value", func(input string, expected uint) {
				actual := number.MustParseUint(input)
				Expect(actual).To(Equal(expected))
			},
				Entry("0", "0", uint(0)),
				Entry("1", "1", uint(1)),
				Entry("42", "42", uint(42)),
				Entry(strconv.FormatUint(uint64(semver.MaxMajor), 10), strconv.FormatUint(uint64(semver.MaxMinor), 10), semver.MaxPatch),
			)
		})

		Context("when s is not a valid unsigned integer", func() {
			It("should panic", func() {
				sut := func() {
					number.MustParseUint("0.0.0-00")
				}
				Expect(sut).To(Panic())
			})
		})
	})
})

func splitDotSeparatedString(s string) []string {
	if s == "" {
		return []string{}
	}

	return strings.Split(s, ".")
}
