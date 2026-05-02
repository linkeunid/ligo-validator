package ligovalidator

import "cmp"

// IsPositive reports whether n > 0.
func IsPositive[T cmp.Ordered](n T) bool { var z T; return n > z }

// IsNegative reports whether n < 0.
func IsNegative[T cmp.Ordered](n T) bool { var z T; return n < z }

// IsNonNegative reports whether n >= 0.
func IsNonNegative[T cmp.Ordered](n T) bool { var z T; return n >= z }

// IsNonPositive reports whether n <= 0.
func IsNonPositive[T cmp.Ordered](n T) bool { var z T; return n <= z }

// InRange reports whether min <= n <= max.
func InRange[T cmp.Ordered](n, min, max T) bool { return n >= min && n <= max }
