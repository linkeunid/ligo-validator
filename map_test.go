package ligovalidator

import "testing"

func TestMapNotEmpty(t *testing.T) {
	if MapNotEmpty(map[string]int{}) {
		t.Fatal("expected false for empty map")
	}
	if !MapNotEmpty(map[string]int{"a": 1}) {
		t.Fatal("expected true for non-empty map")
	}
}

func TestMapContainsKey(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	if !MapContainsKey(m, "a") {
		t.Fatal("expected true for present key")
	}
	if MapContainsKey(m, "z") {
		t.Fatal("expected false for absent key")
	}
	if MapContainsKey(map[string]int{}, "a") {
		t.Fatal("expected false for empty map")
	}
}

func TestMapNotContainsKey(t *testing.T) {
	m := map[string]int{"a": 1}
	if MapNotContainsKey(m, "a") {
		t.Fatal("expected false when key is present")
	}
	if !MapNotContainsKey(m, "z") {
		t.Fatal("expected true when key is absent")
	}
}
