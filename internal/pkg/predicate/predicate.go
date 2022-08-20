// SPDX-License-Identifier: MIT

package predicate

func And[T any](p ...func(t T) bool) func(t T) bool {
	return func(t T) bool {
		result := true
		for _, f := range p {
			result = result && f(t)
		}
		return result
	}
}

func Negate[T any](p func(t T) bool) func(t T) bool {
	return func(t T) bool {
		return !p(t)
	}
}

func Or[T any](p ...func(t T) bool) func(t T) bool {
	return func(t T) bool {
		result := false
		for _, f := range p {
			result = result || f(t)
		}
		return result
	}
}
