// SPDX-License-Identifier: MIT

package predicate

import "github.com/samber/lo"

func And[T any](p ...func(t T) bool) func(t T) bool {
	return func(t T) bool {
		accumulator := func(agg bool, item func(t T) bool, _ int) bool {
			return agg && item(t)
		}
		return lo.Reduce(p, accumulator, true)
	}
}

func Negate[T any](p func(t T) bool) func(t T) bool {
	return func(t T) bool {
		return !p(t)
	}
}

func Or[T any](p ...func(t T) bool) func(t T) bool {
	return func(t T) bool {
		accumulator := func(agg bool, item func(t T) bool, _ int) bool {
			return agg || item(t)
		}
		return lo.Reduce(p, accumulator, false)
	}
}
