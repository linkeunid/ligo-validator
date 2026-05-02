# Slice helpers

Generic pure functions for slice validation. No validator instance required.

---

## API reference

### SliceNotEmpty

```go
func SliceNotEmpty[T any](s []T) bool
```

Reports whether `s` has at least one element.

```go
ligovalidator.SliceNotEmpty([]string{"a"}) // true
ligovalidator.SliceNotEmpty([]string{})    // false
```

### SliceMinSize

```go
func SliceMinSize[T any](s []T, n int) bool
```

Reports whether `s` has at least `n` elements.

```go
ligovalidator.SliceMinSize([]int{1, 2, 3}, 2) // true
ligovalidator.SliceMinSize([]int{1}, 3)        // false
```

### SliceMaxSize

```go
func SliceMaxSize[T any](s []T, n int) bool
```

Reports whether `s` has at most `n` elements.

```go
ligovalidator.SliceMaxSize([]int{1, 2}, 5)       // true
ligovalidator.SliceMaxSize([]int{1, 2, 3, 4}, 3) // false
```

### SliceContains

```go
func SliceContains[T comparable](s []T, elem T) bool
```

Reports whether `s` contains `elem`.

```go
ligovalidator.SliceContains([]string{"a", "b", "c"}, "b") // true
ligovalidator.SliceContains([]string{"a", "b"}, "z")      // false
```

### SliceNotContains

```go
func SliceNotContains[T comparable](s []T, elem T) bool
```

Reports whether `s` does not contain `elem`.

```go
ligovalidator.SliceNotContains([]int{1, 2, 3}, 9) // true
ligovalidator.SliceNotContains([]int{1, 2, 3}, 2) // false
```

### SliceAll

```go
func SliceAll[T any](s []T, fn func(T) bool) bool
```

Reports whether `fn` returns true for every element in `s`. Returns true for an empty slice.

```go
ligovalidator.SliceAll([]int{1, 2, 3}, func(n int) bool { return n > 0 }) // true
ligovalidator.SliceAll([]int{1, -1, 3}, func(n int) bool { return n > 0 }) // false
ligovalidator.SliceAll([]int{}, func(n int) bool { return n > 0 })          // true
```

### SliceAny

```go
func SliceAny[T any](s []T, fn func(T) bool) bool
```

Reports whether `fn` returns true for at least one element in `s`. Returns false for an empty slice.

```go
ligovalidator.SliceAny([]int{-1, 0, 1}, func(n int) bool { return n > 0 }) // true
ligovalidator.SliceAny([]int{-1, -2}, func(n int) bool { return n > 0 })    // false
ligovalidator.SliceAny([]int{}, func(n int) bool { return n > 0 })           // false
```

### SliceUnique

```go
func SliceUnique[T comparable](s []T) bool
```

Reports whether all elements in `s` are distinct.

```go
ligovalidator.SliceUnique([]string{"a", "b", "c"}) // true
ligovalidator.SliceUnique([]string{"a", "b", "b"}) // false
ligovalidator.SliceUnique([]string{})               // true
```
