// SPDX-License-Identifier: MIT

package number_test

import (
	"fmt"
	"math/bits"

	"github.com/ffurrer2/semver/v2/internal/pkg/number"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("number: ", func() {
	var (
		maxUint         = uint(1<<bits.UintSize - 1)
		maxUintAsString = fmt.Sprintf("%d", maxUint)
	)

	Describe("func ParseUint(s string) (uint, error)", func() {
		Context(`when s is "0"`, func() {
			It("should return the corresponding uint", func() {
				number, _ := number.ParseUint("0")
				Expect(number).To(Equal(uint(0)))
			})
			It("should not error", func() {
				_, err := number.ParseUint("0")
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
		Context("when s corresponds to a negative number", func() {
			It("should return 0", func() {
				number, _ := number.ParseUint("-1")
				Expect(number).To(Equal(uint(0)))
			})
			It("should error", func() {
				_, err := number.ParseUint("-1")
				Expect(err).Should(HaveOccurred())
			})
		})
		Context("when s corresponds to the maximum unsigned integer", func() {
			It("should return the corresponding uint", func() {
				number, _ := number.ParseUint(maxUintAsString)
				Expect(number).To(Equal(maxUint))
			})
			It("should not error", func() {
				_, err := number.ParseUint(maxUintAsString)
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
		Context("when s corresponds to the maximum unsigned integer + 1", func() {
			maxUintPlusOneAsString := fmt.Sprintf("%.0f", float64(maxUint)+1.0)
			It("should return 0", func() {
				number, _ := number.ParseUint(maxUintPlusOneAsString)
				Expect(number).To(Equal(uint(0)))
			})
			It("should error", func() {
				_, err := number.ParseUint(maxUintPlusOneAsString)
				Expect(err).Should(HaveOccurred())
			})
		})
	})

	Describe("func MustParseUint(s string) uint", func() {
		Context(`when s corresponds to a valid unsigned integer`, func() {
			It("should return the corresponding uint", func() {
				number := number.MustParseUint("0")
				Expect(number).To(Equal(uint(0)))
			})
			It("should not panic", func() {
				actual := func() {
					_ = number.MustParseUint("0")
				}
				Expect(actual).ShouldNot(Panic())
			})
		})
		Context(`when s corresponds to an invalid unsigned integer`, func() {
			It("should panic", func() {
				actual := func() {
					_ = number.MustParseUint("-1")
				}
				Expect(actual).Should(Panic())
			})
		})
	})

	Describe("func IsNumeric(s string) bool", func() {
		Describe("when s corresponds to a valid unsigned integer", func() {
			DescribeTable("it ",
				func(input string) {
					actual := number.IsNumeric(input)
					Expect(actual).To(BeTrue())
				},
				Entry("should return true", "0"),
				Entry("should return true", "1"),
				Entry("should return true", "4"),
				Entry(maxUintAsString, maxUintAsString),
			)
		})
		DescribeTable("when s corresponds to an invalid unsigned integer",
			func(input string) {
				actual := number.IsNumeric(input)
				Expect(actual).To(BeFalse())
			},
			Entry("should return false", ""),
			Entry("should return false", "-1"),
			Entry("should return false", "123d"),
			Entry("should return false", "abc"),
		)
	})

	Describe("CompareUint[T constraints.Ordered](a, b T) int", func() {
		DescribeTable("when a == b",
			func(a uint, b uint) {
				result := number.CompareUint(a, b)
				Expect(result).To(Equal(0))
			},
			Entry("should return 0", uint(0), uint(0)),
			Entry("should return 0", uint(1), uint(1)),
			Entry("should return 0", maxUint, maxUint),
		)
		DescribeTable("when a < b ",
			func(a uint, b uint) {
				result := number.CompareUint(a, b)
				Expect(result).To(Equal(-1))
			},
			Entry("should return -1", uint(0), uint(1)),
			Entry("should return -1", uint(1), uint(2)),
			Entry("should return -1", maxUint-1, maxUint),
		)
		DescribeTable("when a > b ",
			func(a uint, b uint) {
				result := number.CompareUint(a, b)
				Expect(result).To(Equal(1))
			},
			Entry("should return 1", uint(1), uint(0)),
			Entry("should return 1", uint(2), uint(1)),
			Entry("should return 1", maxUint, maxUint-1),
		)
	})

	Describe("MinInt[T constraints.Ordered](x, y T) T", func() {
		DescribeTable("when a == b",
			func(a uint, b uint) {
				result := number.MinInt(a, b)
				Expect(result).To(Equal(a))
				Expect(result).To(Equal(b))
			},
			Entry("should return 0", uint(0), uint(0)),
			Entry("should return 0", uint(1), uint(1)),
			Entry("should return 0", maxUint, maxUint),
		)
		DescribeTable("when a < b ",
			func(a uint, b uint) {
				result := number.MinInt(a, b)
				Expect(result).To(Equal(a))
			},
			Entry("should return -1", uint(0), uint(1)),
			Entry("should return -1", uint(1), uint(2)),
			Entry("should return -1", maxUint-1, maxUint),
		)
		DescribeTable("when a > b ",
			func(a uint, b uint) {
				result := number.MinInt(a, b)
				Expect(result).To(Equal(b))
			},
			Entry("should return 1", uint(1), uint(0)),
			Entry("should return 1", uint(2), uint(1)),
			Entry("should return 1", maxUint, maxUint-1),
		)
	})
})
