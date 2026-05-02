package ligovalidator

import "slices"

// SliceNotEmpty reports whether s has at least one element.
func SliceNotEmpty[T any](s []T) bool { return len(s) > 0 }

// SliceMinSize reports whether s has at least n elements.
func SliceMinSize[T any](s []T, n int) bool { return len(s) >= n }

// SliceMaxSize reports whether s has at most n elements.
func SliceMaxSize[T any](s []T, n int) bool { return len(s) <= n }

// SliceContains reports whether s contains elem.
func SliceContains[T comparable](s []T, elem T) bool { return slices.Contains(s, elem) }

// SliceNotContains reports whether s does not contain elem.
func SliceNotContains[T comparable](s []T, elem T) bool { return !slices.Contains(s, elem) }

// SliceUnique reports whether all elements in s are distinct.
func SliceUnique[T comparable](s []T) bool {
	if len(s) == 0 {
		return true
	}
	seen := make(map[T]struct{}, len(s))
	for _, v := range s {
		if _, exists := seen[v]; exists {
			return false
		}
		seen[v] = struct{}{}
	}
	return true
}
