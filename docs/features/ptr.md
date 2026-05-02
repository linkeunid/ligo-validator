# Pointer helpers

Generic pure functions for nil-safety checks.

---

## API reference

### IsNil

```go
func IsNil[T any](v *T) bool
```

Reports whether `v` is nil.

```go
var p *int
ligovalidator.IsNil(p) // true

n := 42
ligovalidator.IsNil(&n) // false
```

### IsNotNil

```go
func IsNotNil[T any](v *T) bool
```

Reports whether `v` is not nil.

```go
s := "hello"
ligovalidator.IsNotNil(&s) // true

var p *string
ligovalidator.IsNotNil(p) // false
```
