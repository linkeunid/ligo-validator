package ligovalidator

import "reflect"

// IsNotEmptyObject reports whether v (a struct or map) has at least one non-zero
// field or entry. Returns false for nil, non-struct/non-map types, empty maps,
// and structs where every exported field is the zero value.
func IsNotEmptyObject(v any) bool {
	if v == nil {
		return false
	}
	rv := reflect.ValueOf(v)
	for rv.Kind() == reflect.Pointer {
		if rv.IsNil() {
			return false
		}
		rv = rv.Elem()
	}
	switch rv.Kind() {
	case reflect.Map:
		return rv.Len() > 0
	case reflect.Struct:
		rt := rv.Type()
		for i := range rv.NumField() {
			if rt.Field(i).IsExported() && !rv.Field(i).IsZero() {
				return true
			}
		}
		return false
	default:
		return false
	}
}
