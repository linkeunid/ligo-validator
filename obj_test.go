package ligovalidator

import "testing"

type emptyStruct struct{}

type partialStruct struct {
	Name string
	Age  int
}

func TestIsNotEmptyObject_Nil(t *testing.T) {
	if IsNotEmptyObject(nil) {
		t.Fatal("expected false for nil")
	}
}

func TestIsNotEmptyObject_EmptyStruct(t *testing.T) {
	if IsNotEmptyObject(emptyStruct{}) {
		t.Fatal("expected false for empty struct")
	}
}

func TestIsNotEmptyObject_AllZeroFields(t *testing.T) {
	if IsNotEmptyObject(partialStruct{}) {
		t.Fatal("expected false for struct with all zero fields")
	}
}

func TestIsNotEmptyObject_OneNonZeroField(t *testing.T) {
	if !IsNotEmptyObject(partialStruct{Name: "Alice"}) {
		t.Fatal("expected true for struct with non-zero field")
	}
}

func TestIsNotEmptyObject_PointerToStruct(t *testing.T) {
	p := &partialStruct{Age: 30}
	if !IsNotEmptyObject(p) {
		t.Fatal("expected true for pointer to struct with non-zero field")
	}
}

func TestIsNotEmptyObject_NilPointer(t *testing.T) {
	var p *partialStruct
	if IsNotEmptyObject(p) {
		t.Fatal("expected false for nil pointer")
	}
}

func TestIsNotEmptyObject_EmptyMap(t *testing.T) {
	if IsNotEmptyObject(map[string]any{}) {
		t.Fatal("expected false for empty map")
	}
}

func TestIsNotEmptyObject_NonEmptyMap(t *testing.T) {
	if !IsNotEmptyObject(map[string]any{"key": "value"}) {
		t.Fatal("expected true for non-empty map")
	}
}

func TestIsNotEmptyObject_NonStructType(t *testing.T) {
	if IsNotEmptyObject("string") {
		t.Fatal("expected false for non-struct/non-map type")
	}
	if IsNotEmptyObject(42) {
		t.Fatal("expected false for int")
	}
}
