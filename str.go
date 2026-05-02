package ligovalidator

import (
	"slices"
	"strings"
)

// IsEmpty reports whether s has zero length.
func IsEmpty(s string) bool { return s == "" }

// IsNotEmpty reports whether s has non-zero length.
func IsNotEmpty(s string) bool { return s != "" }

// IsBlank reports whether s is empty or contains only whitespace characters.
func IsBlank(s string) bool { return strings.TrimSpace(s) == "" }

// IsNotBlank reports whether s contains at least one non-whitespace character.
func IsNotBlank(s string) bool { return strings.TrimSpace(s) != "" }

// HasMinLength reports whether s has at least n runes.
func HasMinLength(s string, n int) bool { return len([]rune(s)) >= n }

// HasMaxLength reports whether s has at most n runes.
func HasMaxLength(s string, n int) bool { return len([]rune(s)) <= n }

// StrContains reports whether s contains substr.
func StrContains(s, substr string) bool { return strings.Contains(s, substr) }

// StrNotContains reports whether s does not contain substr.
func StrNotContains(s, substr string) bool { return !strings.Contains(s, substr) }

// StrOneOf reports whether s is equal to one of the given values.
// Returns false if no values are provided.
func StrOneOf(s string, values ...string) bool { return slices.Contains(values, s) }
