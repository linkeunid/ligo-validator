package ligovalidator

import "testing"

func TestIsPositive(t *testing.T) {
	if !IsPositive(1) {
		t.Fatal("expected true for 1")
	}
	if !IsPositive(0.1) {
		t.Fatal("expected true for 0.1")
	}
	if IsPositive(0) {
		t.Fatal("expected false for 0")
	}
	if IsPositive(-1) {
		t.Fatal("expected false for -1")
	}
}

func TestIsNegative(t *testing.T) {
	if !IsNegative(-1) {
		t.Fatal("expected true for -1")
	}
	if IsNegative(0) {
		t.Fatal("expected false for 0")
	}
	if IsNegative(1) {
		t.Fatal("expected false for 1")
	}
}

func TestIsNonNegative(t *testing.T) {
	if !IsNonNegative(0) {
		t.Fatal("expected true for 0")
	}
	if !IsNonNegative(5) {
		t.Fatal("expected true for 5")
	}
	if IsNonNegative(-1) {
		t.Fatal("expected false for -1")
	}
}

func TestIsNonPositive(t *testing.T) {
	if !IsNonPositive(0) {
		t.Fatal("expected true for 0")
	}
	if !IsNonPositive(-3) {
		t.Fatal("expected true for -3")
	}
	if IsNonPositive(1) {
		t.Fatal("expected false for 1")
	}
}

func TestInRange(t *testing.T) {
	if !InRange(5, 1, 10) {
		t.Fatal("expected true for 5 in [1, 10]")
	}
	if !InRange(1, 1, 10) {
		t.Fatal("expected true for min boundary")
	}
	if !InRange(10, 1, 10) {
		t.Fatal("expected true for max boundary")
	}
	if InRange(0, 1, 10) {
		t.Fatal("expected false for value below min")
	}
	if InRange(11, 1, 10) {
		t.Fatal("expected false for value above max")
	}
	if !InRange(3.5, 1.0, 5.0) {
		t.Fatal("expected true for float in range")
	}
}
