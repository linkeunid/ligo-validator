package ligovalidator

// IsZero reports whether v is the zero value of its type.
func IsZero[T comparable](v T) bool { var z T; return v == z }

// IsNotZero reports whether v is not the zero value of its type.
func IsNotZero[T comparable](v T) bool { var z T; return v != z }
