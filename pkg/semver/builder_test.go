// SPDX-License-Identifier: MIT
package semver_test

import (
	"github.com/ffurrer2/semver/pkg/semver"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("semver:", func() {
	Describe("Calling func NewBuilder() *Builder", func() {
		It("should return a new Builder", func() {
			builder := semver.NewBuilder()
			Expect(builder).ShouldNot(BeNil())
			actual, ok := builder.Build()
			Expect(actual).ShouldNot(BeNil())
			Expect(ok).Should(BeTrue())
			Expect(actual.Major).To(Equal(uint(0)))
			Expect(actual.Minor).To(Equal(uint(0)))
			Expect(actual.Patch).To(Equal(uint(0)))
			Expect(actual.PreRelease).To(Equal([]string{}))
			Expect(actual.BuildMetadata).To(Equal([]string{}))
		})
	})

	Describe("Calling func (b *Builder) Major(major uint) *Builder", func() {
		It("should set major version correctly", func() {
			builder := semver.NewBuilder().Major(42)
			Expect(builder).ShouldNot(BeNil())
			actual, ok := builder.Build()
			Expect(actual).ShouldNot(BeNil())
			Expect(ok).Should(BeTrue())
			Expect(actual.Major).To(Equal(uint(42)))
			Expect(actual.Minor).To(Equal(uint(0)))
			Expect(actual.Patch).To(Equal(uint(0)))
			Expect(actual.PreRelease).To(Equal([]string{}))
			Expect(actual.BuildMetadata).To(Equal([]string{}))
		})
	})

	Describe("Calling func (b *Builder) Minor(minor uint) *Builder", func() {
		It("should set minor version correctly", func() {
			builder := semver.NewBuilder().Minor(42)
			Expect(builder).ShouldNot(BeNil())
			actual, ok := builder.Build()
			Expect(actual).ShouldNot(BeNil())
			Expect(ok).Should(BeTrue())
			Expect(actual.Major).To(Equal(uint(0)))
			Expect(actual.Minor).To(Equal(uint(42)))
			Expect(actual.Patch).To(Equal(uint(0)))
			Expect(actual.PreRelease).To(Equal([]string{}))
			Expect(actual.BuildMetadata).To(Equal([]string{}))
		})
	})

	Describe("Calling func (b *Builder) Patch(patch uint) *Builder", func() {
		It("should set patch version correctly", func() {
			builder := semver.NewBuilder().Patch(42)
			Expect(builder).ShouldNot(BeNil())
			actual, ok := builder.Build()
			Expect(actual).ShouldNot(BeNil())
			Expect(ok).Should(BeTrue())
			Expect(actual.Major).To(Equal(uint(0)))
			Expect(actual.Minor).To(Equal(uint(0)))
			Expect(actual.Patch).To(Equal(uint(42)))
			Expect(actual.PreRelease).To(Equal([]string{}))
			Expect(actual.BuildMetadata).To(Equal([]string{}))
		})
	})

	Describe("Calling func (b Builder) PreRelease(preRelease []string) *Builder", func() {
		It("should set patch version correctly", func() {
			builder := semver.NewBuilder().PreRelease([]string{"alpha", "1"})
			Expect(builder).ShouldNot(BeNil())
			actual, ok := builder.Build()
			Expect(actual).ShouldNot(BeNil())
			Expect(ok).Should(BeTrue())
			Expect(actual.Major).To(Equal(uint(0)))
			Expect(actual.Minor).To(Equal(uint(0)))
			Expect(actual.Patch).To(Equal(uint(0)))
			Expect(actual.PreRelease).To(Equal([]string{"alpha", "1"}))
			Expect(actual.BuildMetadata).To(Equal([]string{}))
		})
	})

	Describe("Calling func (b Builder) PreReleaseField(preRelease string) *Builder", func() {
		It("should set patch version correctly", func() {
			builder := semver.NewBuilder().PreReleaseField("alpha")
			Expect(builder).ShouldNot(BeNil())
			actual, ok := builder.Build()
			Expect(actual).ShouldNot(BeNil())
			Expect(ok).Should(BeTrue())
			Expect(actual.Major).To(Equal(uint(0)))
			Expect(actual.Minor).To(Equal(uint(0)))
			Expect(actual.Patch).To(Equal(uint(0)))
			Expect(actual.PreRelease).To(Equal([]string{"alpha"}))
			Expect(actual.BuildMetadata).To(Equal([]string{}))
		})
	})

	Describe("Calling func (b Builder) BuildMetadata(buildMetadata []string) *Builder", func() {
		It("should set patch version correctly", func() {
			builder := semver.NewBuilder().BuildMetadata([]string{"2020.01.01", "1"})
			Expect(builder).ShouldNot(BeNil())
			actual, ok := builder.Build()
			Expect(actual).ShouldNot(BeNil())
			Expect(ok).Should(BeTrue())
			Expect(actual.Major).To(Equal(uint(0)))
			Expect(actual.Minor).To(Equal(uint(0)))
			Expect(actual.Patch).To(Equal(uint(0)))
			Expect(actual.PreRelease).To(Equal([]string{}))
			Expect(actual.BuildMetadata).To(Equal([]string{"2020.01.01", "1"}))
		})
	})

	Describe("Calling func (b Builder) BuildMetadataField(buildMetadata string) *Builder", func() {
		It("should set patch version correctly", func() {
			builder := semver.NewBuilder().BuildMetadataField("2020.01.01")
			Expect(builder).ShouldNot(BeNil())
			actual, ok := builder.Build()
			Expect(actual).ShouldNot(BeNil())
			Expect(ok).Should(BeTrue())
			Expect(actual.Major).To(Equal(uint(0)))
			Expect(actual.Minor).To(Equal(uint(0)))
			Expect(actual.Patch).To(Equal(uint(0)))
			Expect(actual.PreRelease).To(Equal([]string{}))
			Expect(actual.BuildMetadata).To(Equal([]string{"2020.01.01"}))
		})
	})

	Describe("Calling func (b Builder) Build() (*SemVer, bool)", func() {
		Describe("when building a valid semantic version", func() {
			It("should return the correct SemVer struct", func() {
				builder := semver.NewBuilder().
					Major(1).
					Minor(2).
					Patch(3).
					PreReleaseField("alpha").
					BuildMetadataField("2020.01.01")
				Expect(builder).ShouldNot(BeNil())
				actual, ok := builder.Build()
				Expect(ok).To(BeTrue())
				Expect(actual.Major).To(Equal(uint(1)))
				Expect(actual.Minor).To(Equal(uint(2)))
				Expect(actual.Patch).To(Equal(uint(3)))
				Expect(actual.PreRelease).To(Equal([]string{"alpha"}))
				Expect(actual.BuildMetadata).To(Equal([]string{"2020.01.01"}))
			})
		})
	})
})
