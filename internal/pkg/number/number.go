// SPDX-License-Identifier: MIT

package number

import (
	"fmt"
	"math/bits"
	"regexp"
	"strconv"
)

const (
	pattern = `^(0|[1-9]\d*)$`
	base    = 10
)

var /* const */ numericRegexp *regexp.Regexp

func init() {
	numericRegexp = regexp.MustCompile(pattern)
}

func ParseUint(s string) (uint, error) {
	u64, err := strconv.ParseUint(s, base, bits.UintSize)
	if err != nil {
		return 0, fmt.Errorf("failed to parse uint: %w", err)
	}
	return uint(u64), nil
}

func MustParseUint(s string) uint {
	u64, err := ParseUint(s)
	if err != nil {
		panic(err)
	}
	return u64
}

func IsNumeric(s string) bool {
	return numericRegexp.MatchString(s)
}

func CompareUint(a, b uint) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return +1
}

func MinInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}
