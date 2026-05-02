package ligovalidator

import "testing"

func TestIsNil(t *testing.T) {
	var p *int
	if !IsNil(p) {
		t.Fatal("expected true for nil pointer")
	}
	v := 42
	if IsNil(&v) {
		t.Fatal("expected false for non-nil pointer")
	}
}

func TestIsNotNil(t *testing.T) {
	var p *string
	if IsNotNil(p) {
		t.Fatal("expected false for nil pointer")
	}
	s := "hello"
	if !IsNotNil(&s) {
		t.Fatal("expected true for non-nil pointer")
	}
}
