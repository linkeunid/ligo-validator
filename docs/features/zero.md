# Zero helpers

Generic pure functions for zero-value checks. Works with any comparable type — `int`, `string`, `bool`, structs, UUIDs, and their aliases.

---

## API reference

### IsZero

```go
func IsZero[T comparable](v T) bool
```

Reports whether `v` is the zero value of its type.

```go
ligovalidator.IsZero(0)     // true
ligovalidator.IsZero("")    // true
ligovalidator.IsZero(false) // true
ligovalidator.IsZero(1)     // false
ligovalidator.IsZero("hi")  // false
```

### IsNotZero

```go
func IsNotZero[T comparable](v T) bool
```

Reports whether `v` is not the zero value of its type.

```go
ligovalidator.IsNotZero(42)    // true
ligovalidator.IsNotZero("hi")  // true
ligovalidator.IsNotZero(0)     // false
ligovalidator.IsNotZero("")    // false
```

---

## Notes

- Complements `IsNil`/`IsNotNil` (pointer-only) with a generic zero check for any comparable value.
- For structs, the zero check uses `==` — the struct must be comparable. Use `IsNotEmptyObject` for reflect-based field inspection.
