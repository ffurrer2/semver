// SPDX-License-Identifier: MIT

package number

import (
	"fmt"
	"math/bits"
	"regexp"
	"strconv"

	"golang.org/x/exp/constraints"
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
	u, err := strconv.ParseUint(s, base, bits.UintSize)
	if err != nil {
		return 0, fmt.Errorf("failed to parse uint: %w", err)
	}
	return uint(u), nil
}

func MustParseUint(s string) uint {
	u, err := ParseUint(s)
	if err != nil {
		panic(err)
	}
	return u
}

func IsNumeric(s string) bool {
	return numericRegexp.MatchString(s)
}

func CompareInt[T constraints.Integer](a, b T) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return +1
}
