package ligovalidator

// IsNil reports whether v is nil.
func IsNil[T any](v *T) bool { return v == nil }

// IsNotNil reports whether v is not nil.
func IsNotNil[T any](v *T) bool { return v != nil }
