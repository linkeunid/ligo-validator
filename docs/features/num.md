# Number helpers

Generic pure functions for numeric validation using `cmp.Ordered`. Works with any ordered type: `int`, `float64`, `string`, and their aliases.

---

## API reference

### IsPositive

```go
func IsPositive[T cmp.Ordered](n T) bool
```

Reports whether `n > 0`.

```go
ligovalidator.IsPositive(1)    // true
ligovalidator.IsPositive(0.1)  // true
ligovalidator.IsPositive(0)    // false
ligovalidator.IsPositive(-1)   // false
```

### IsNegative

```go
func IsNegative[T cmp.Ordered](n T) bool
```

Reports whether `n < 0`.

```go
ligovalidator.IsNegative(-1) // true
ligovalidator.IsNegative(0)  // false
ligovalidator.IsNegative(1)  // false
```

### IsNonNegative

```go
func IsNonNegative[T cmp.Ordered](n T) bool
```

Reports whether `n >= 0`.

```go
ligovalidator.IsNonNegative(0)  // true
ligovalidator.IsNonNegative(5)  // true
ligovalidator.IsNonNegative(-1) // false
```

### IsNonPositive

```go
func IsNonPositive[T cmp.Ordered](n T) bool
```

Reports whether `n <= 0`.

```go
ligovalidator.IsNonPositive(0)  // true
ligovalidator.IsNonPositive(-3) // true
ligovalidator.IsNonPositive(1)  // false
```

### InRange

```go
func InRange[T cmp.Ordered](n, min, max T) bool
```

Reports whether `min <= n <= max`. Both boundaries are inclusive.

```go
ligovalidator.InRange(5, 1, 10)      // true
ligovalidator.InRange(1, 1, 10)      // true  — min boundary
ligovalidator.InRange(10, 1, 10)     // true  — max boundary
ligovalidator.InRange(0, 1, 10)      // false
ligovalidator.InRange(3.5, 1.0, 5.0) // true
```
