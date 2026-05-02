# String helpers

Pure functions for string validation. No validator instance required.

---

## API reference

### IsEmpty

```go
func IsEmpty(s string) bool
```

Reports whether `s` has zero length.

```go
ligovalidator.IsEmpty("")    // true
ligovalidator.IsEmpty("hi") // false
```

### IsNotEmpty

```go
func IsNotEmpty(s string) bool
```

Reports whether `s` has non-zero length.

```go
ligovalidator.IsNotEmpty("hi") // true
ligovalidator.IsNotEmpty("")   // false
```

### IsBlank

```go
func IsBlank(s string) bool
```

Reports whether `s` is empty or contains only whitespace characters.

```go
ligovalidator.IsBlank("")     // true
ligovalidator.IsBlank("   ")  // true
ligovalidator.IsBlank("  a ") // false
```

### IsNotBlank

```go
func IsNotBlank(s string) bool
```

Reports whether `s` contains at least one non-whitespace character.

```go
ligovalidator.IsNotBlank("hi")  // true
ligovalidator.IsNotBlank("   ") // false
```

### HasMinLength

```go
func HasMinLength(s string, n int) bool
```

Reports whether `s` has at least `n` runes. Rune-aware — multibyte characters count as one.

```go
ligovalidator.HasMinLength("hello", 5)     // true
ligovalidator.HasMinLength("こんにちは", 5) // true — 5 runes, 15 bytes
ligovalidator.HasMinLength("hi", 5)        // false
```

### HasMaxLength

```go
func HasMaxLength(s string, n int) bool
```

Reports whether `s` has at most `n` runes.

```go
ligovalidator.HasMaxLength("hi", 10)    // true
ligovalidator.HasMaxLength("toolong", 5) // false
```

### StrContains

```go
func StrContains(s, substr string) bool
```

Reports whether `s` contains `substr`.

```go
ligovalidator.StrContains("hello world", "world") // true
ligovalidator.StrContains("hello", "xyz")         // false
```

### StrNotContains

```go
func StrNotContains(s, substr string) bool
```

Reports whether `s` does not contain `substr`.

```go
ligovalidator.StrNotContains("hello", "xyz")         // true
ligovalidator.StrNotContains("hello world", "world") // false
```
