package ligovalidator

// MapNotEmpty reports whether m has at least one entry.
func MapNotEmpty[K comparable, V any](m map[K]V) bool { return len(m) > 0 }

// MapContainsKey reports whether m contains key.
func MapContainsKey[K comparable, V any](m map[K]V, key K) bool { _, ok := m[key]; return ok }

// MapNotContainsKey reports whether m does not contain key.
func MapNotContainsKey[K comparable, V any](m map[K]V, key K) bool { _, ok := m[key]; return !ok }
