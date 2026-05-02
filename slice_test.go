package ligovalidator

import "testing"

func TestSliceNotEmpty(t *testing.T) {
	if SliceNotEmpty([]int{}) {
		t.Fatal("expected false for empty slice")
	}
	if !SliceNotEmpty([]int{1}) {
		t.Fatal("expected true for non-empty slice")
	}
}

func TestSliceMinSize(t *testing.T) {
	if !SliceMinSize([]int{1, 2, 3}, 3) {
		t.Fatal("expected true for len==min")
	}
	if !SliceMinSize([]int{1, 2, 3}, 2) {
		t.Fatal("expected true for len>min")
	}
	if SliceMinSize([]int{1}, 3) {
		t.Fatal("expected false for len<min")
	}
}

func TestSliceMaxSize(t *testing.T) {
	if !SliceMaxSize([]int{1, 2}, 5) {
		t.Fatal("expected true for len<max")
	}
	if !SliceMaxSize([]int{1, 2, 3}, 3) {
		t.Fatal("expected true for len==max")
	}
	if SliceMaxSize([]int{1, 2, 3, 4}, 3) {
		t.Fatal("expected false for len>max")
	}
}

func TestSliceContains(t *testing.T) {
	if !SliceContains([]string{"a", "b", "c"}, "b") {
		t.Fatal("expected true for present element")
	}
	if SliceContains([]string{"a", "b"}, "z") {
		t.Fatal("expected false for absent element")
	}
	if SliceContains([]int{}, 1) {
		t.Fatal("expected false for empty slice")
	}
}

func TestSliceNotContains(t *testing.T) {
	if SliceNotContains([]int{1, 2, 3}, 2) {
		t.Fatal("expected false when element is present")
	}
	if !SliceNotContains([]int{1, 2, 3}, 9) {
		t.Fatal("expected true when element is absent")
	}
}

func TestSliceUnique(t *testing.T) {
	if !SliceUnique([]int{1, 2, 3}) {
		t.Fatal("expected true for unique slice")
	}
	if SliceUnique([]int{1, 2, 2}) {
		t.Fatal("expected false for slice with duplicate")
	}
	if !SliceUnique([]string{}) {
		t.Fatal("expected true for empty slice")
	}
	if !SliceUnique([]string{"a"}) {
		t.Fatal("expected true for single-element slice")
	}
}
