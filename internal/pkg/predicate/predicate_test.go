// SPDX-License-Identifier: MIT

package predicate_test

import (
	"github.com/ffurrer2/semver/internal/pkg/predicate"
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = Describe("predicate: ", func() {
	Describe("func And[T any](p ...func(t T) bool) func(t T) bool", func() {
		Context(`when len(p) is 1 and p0(t) is true`, func() {
			It("should return true", func() {
				p0 := func(any) bool {
					return true
				}
				actual := predicate.And(p0)(nil)
				gomega.Expect(actual).To(gomega.BeTrue())
			})
		})
		Context(`when len(p) is 1 and p0(t) is false`, func() {
			It("should return false", func() {
				p0 := func(any) bool {
					return false
				}
				actual := predicate.And(p0)(nil)
				gomega.Expect(actual).To(gomega.BeFalse())
			})
		})
		Context(`when len(p) > 1, p0(t) is true and p1(t) is false`, func() {
			It("should return false", func() {
				p0 := func(any) bool {
					return true
				}
				p1 := func(any) bool {
					return false
				}
				actual := predicate.And(p0, p1)(nil)
				gomega.Expect(actual).To(gomega.BeFalse())
			})
		})
		Context(`when len(p) > 1, p0(t) is false and p1(t) is true`, func() {
			It("should return true", func() {
				p0 := func(any) bool {
					return false
				}
				p1 := func(any) bool {
					return true
				}
				actual := predicate.And(p0, p1)(nil)
				gomega.Expect(actual).To(gomega.BeFalse())
			})
		})
		Context(`when len(p) > 1, p0(t) is false and p1(t) is false`, func() {
			It("should return false", func() {
				p0 := func(any) bool {
					return false
				}
				p1 := func(any) bool {
					return false
				}
				actual := predicate.And(p0, p1)(nil)
				gomega.Expect(actual).To(gomega.BeFalse())
			})
		})
		Context(`when len(p) > 1, p0(t) is true and p1(t) is true`, func() {
			It("should return true", func() {
				p0 := func(any) bool {
					return true
				}
				p1 := func(any) bool {
					return true
				}
				actual := predicate.And(p0, p1)(nil)
				gomega.Expect(actual).To(gomega.BeTrue())
			})
		})
	})

	Describe("func Negate[T any](p func(t T) bool) func(t T) bool", func() {
		Context(`when p(t) is true`, func() {
			It("should return false", func() {
				p := func(any) bool {
					return true
				}
				actual := predicate.Negate(p)(nil)
				gomega.Expect(actual).To(gomega.BeFalse())
			})
		})
		Context(`when p(t) is false`, func() {
			It("should return true", func() {
				p := func(any) bool {
					return false
				}
				actual := predicate.Negate(p)(nil)
				gomega.Expect(actual).To(gomega.BeTrue())
			})
		})
	})

	Describe("Or[T any](p ...func(t T) bool) func(t T) bool", func() {
		Context(`when len(p) is 1 and p0(t) is true`, func() {
			It("should return true", func() {
				p0 := func(any) bool {
					return true
				}
				actual := predicate.Or(p0)(nil)
				gomega.Expect(actual).To(gomega.BeTrue())
			})
		})
		Context(`when len(p) is 1 and p0(t) is false`, func() {
			It("should return false", func() {
				p0 := func(any) bool {
					return false
				}
				actual := predicate.Or(p0)(nil)
				gomega.Expect(actual).To(gomega.BeFalse())
			})
		})
		Context(`when len(p) > 1, p0(t) is true and p1(t) is false`, func() {
			It("should return true", func() {
				p0 := func(any) bool {
					return true
				}
				p1 := func(any) bool {
					return false
				}
				actual := predicate.Or(p0, p1)(nil)
				gomega.Expect(actual).To(gomega.BeTrue())
			})
		})
		Context(`when len(p) > 1, p0(t) is false and p1(t) is true`, func() {
			It("should return true", func() {
				p0 := func(any) bool {
					return false
				}
				p1 := func(any) bool {
					return true
				}
				actual := predicate.Or(p0, p1)(nil)
				gomega.Expect(actual).To(gomega.BeTrue())
			})
		})
		Context(`when len(p) > 1, p0(t) is false and p1(t) is false`, func() {
			It("should return false", func() {
				p0 := func(any) bool {
					return false
				}
				p1 := func(any) bool {
					return false
				}
				actual := predicate.Or(p0, p1)(nil)
				gomega.Expect(actual).To(gomega.BeFalse())
			})
		})
		Context(`when len(p) > 1, p0(t) is true and p1(t) is true`, func() {
			It("should return true", func() {
				p0 := func(any) bool {
					return true
				}
				p1 := func(any) bool {
					return true
				}
				actual := predicate.Or(p0, p1)(nil)
				gomega.Expect(actual).To(gomega.BeTrue())
			})
		})
	})
})
