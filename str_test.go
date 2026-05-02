package ligovalidator

import "testing"

func TestIsEmpty(t *testing.T) {
	if !IsEmpty("") {
		t.Fatal("expected true for empty string")
	}
	if IsEmpty("a") {
		t.Fatal("expected false for non-empty string")
	}
}

func TestIsNotEmpty(t *testing.T) {
	if IsNotEmpty("") {
		t.Fatal("expected false for empty string")
	}
	if !IsNotEmpty("a") {
		t.Fatal("expected true for non-empty string")
	}
}

func TestIsBlank(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"", true},
		{" ", true},
		{"\t\n", true},
		{"  a  ", false},
		{"hello", false},
	}
	for _, c := range cases {
		if got := IsBlank(c.in); got != c.want {
			t.Fatalf("IsBlank(%q) = %v, want %v", c.in, got, c.want)
		}
	}
}

func TestIsNotBlank(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"", false},
		{"   ", false},
		{"a", true},
		{"  hi  ", true},
	}
	for _, c := range cases {
		if got := IsNotBlank(c.in); got != c.want {
			t.Fatalf("IsNotBlank(%q) = %v, want %v", c.in, got, c.want)
		}
	}
}

func TestHasMinLength(t *testing.T) {
	if !HasMinLength("hello", 5) {
		t.Fatal("expected true for len==min")
	}
	if !HasMinLength("hello", 3) {
		t.Fatal("expected true for len>min")
	}
	if HasMinLength("hi", 5) {
		t.Fatal("expected false for len<min")
	}
	// rune-aware: "こんにちは" is 5 runes, 15 bytes
	if !HasMinLength("こんにちは", 5) {
		t.Fatal("expected true for multibyte string with len==min")
	}
}

func TestHasMaxLength(t *testing.T) {
	if !HasMaxLength("hi", 5) {
		t.Fatal("expected true for len<max")
	}
	if !HasMaxLength("hello", 5) {
		t.Fatal("expected true for len==max")
	}
	if HasMaxLength("toolong", 5) {
		t.Fatal("expected false for len>max")
	}
}

func TestStrContains(t *testing.T) {
	if !StrContains("hello world", "world") {
		t.Fatal("expected true for present substr")
	}
	if StrContains("hello", "xyz") {
		t.Fatal("expected false for absent substr")
	}
}

func TestStrNotContains(t *testing.T) {
	if StrNotContains("hello world", "world") {
		t.Fatal("expected false when substr is present")
	}
	if !StrNotContains("hello", "xyz") {
		t.Fatal("expected true when substr is absent")
	}
}

func TestStrOneOf(t *testing.T) {
	if !StrOneOf("admin", "admin", "user", "viewer") {
		t.Fatal("expected true for matching value")
	}
	if !StrOneOf("viewer", "admin", "user", "viewer") {
		t.Fatal("expected true for last value")
	}
	if StrOneOf("root", "admin", "user", "viewer") {
		t.Fatal("expected false for non-matching value")
	}
	if StrOneOf("admin") {
		t.Fatal("expected false when no values provided")
	}
}
