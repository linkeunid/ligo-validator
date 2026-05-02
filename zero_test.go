package ligovalidator

import "testing"

func TestIsZero(t *testing.T) {
	if !IsZero(0) {
		t.Fatal("expected true for zero int")
	}
	if !IsZero("") {
		t.Fatal("expected true for zero string")
	}
	if !IsZero(false) {
		t.Fatal("expected true for zero bool")
	}
	if IsZero(1) {
		t.Fatal("expected false for non-zero int")
	}
	if IsZero("hello") {
		t.Fatal("expected false for non-zero string")
	}
	if IsZero(true) {
		t.Fatal("expected false for non-zero bool")
	}
}

func TestIsNotZero(t *testing.T) {
	if IsNotZero(0) {
		t.Fatal("expected false for zero int")
	}
	if IsNotZero("") {
		t.Fatal("expected false for zero string")
	}
	if !IsNotZero(42) {
		t.Fatal("expected true for non-zero int")
	}
	if !IsNotZero("hi") {
		t.Fatal("expected true for non-zero string")
	}
}
