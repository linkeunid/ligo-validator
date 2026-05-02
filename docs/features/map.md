# Map helpers

Generic pure functions for map validation. Mirrors the slice helper vocabulary for maps.

---

## API reference

### MapNotEmpty

```go
func MapNotEmpty[K comparable, V any](m map[K]V) bool
```

Reports whether `m` has at least one entry.

```go
ligovalidator.MapNotEmpty(map[string]int{"a": 1}) // true
ligovalidator.MapNotEmpty(map[string]int{})        // false
```

### MapContainsKey

```go
func MapContainsKey[K comparable, V any](m map[K]V, key K) bool
```

Reports whether `m` contains `key`.

```go
m := map[string]int{"a": 1, "b": 2}
ligovalidator.MapContainsKey(m, "a") // true
ligovalidator.MapContainsKey(m, "z") // false
```

### MapNotContainsKey

```go
func MapNotContainsKey[K comparable, V any](m map[K]V, key K) bool
```

Reports whether `m` does not contain `key`.

```go
m := map[string]int{"a": 1}
ligovalidator.MapNotContainsKey(m, "z") // true
ligovalidator.MapNotContainsKey(m, "a") // false
```
