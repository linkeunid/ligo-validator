# Object helpers

Pure functions for object/map validation using reflection.

---

## API reference

### IsNotEmptyObject

```go
func IsNotEmptyObject(v any) bool
```

Reports whether `v` (a struct or map) has at least one non-zero field or entry. Handles pointer dereferencing automatically.

Returns `false` for:
- `nil`
- nil pointers
- non-struct/non-map types
- empty maps
- structs where every exported field is the zero value

```go
// maps
ligovalidator.IsNotEmptyObject(map[string]any{"k": "v"}) // true
ligovalidator.IsNotEmptyObject(map[string]any{})          // false

// structs
type Profile struct {
    Name string
    Age  int
}

ligovalidator.IsNotEmptyObject(Profile{Name: "Alice"}) // true
ligovalidator.IsNotEmptyObject(Profile{})              // false

// pointers are dereferenced
ligovalidator.IsNotEmptyObject(&Profile{Name: "Alice"}) // true
ligovalidator.IsNotEmptyObject(&Profile{})              // false

// nil
ligovalidator.IsNotEmptyObject(nil)             // false
var p *Profile
ligovalidator.IsNotEmptyObject(p)               // false
```

---

## Notes

- Only **exported** struct fields are considered. Unexported fields are ignored even if non-zero.
- For maps, any non-empty map (regardless of value types) returns `true`.
- The `not_empty_obj` struct tag uses this function internally.
