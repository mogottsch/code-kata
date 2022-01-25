package search

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	if -1 != binarySearch(3, []int{}) {
		t.Fatalf("expected, got")
	}
	if -1 != binarySearch(3, []int{1}) {
		t.Fatalf("expected, got")
	}
	if 0 != binarySearch(1, []int{1}) {
		t.Fatalf("expected, got")
	}

	if 0 != binarySearch(1, []int{1, 3, 5}) {
		t.Fatalf("expected, got")
	}
	if 1 != binarySearch(3, []int{1, 3, 5}) {
		t.Fatalf("expected, got")
	}
	if 2 != binarySearch(5, []int{1, 3, 5}) {
		t.Fatalf("expected, got")
	}
	if -1 != binarySearch(0, []int{1, 3, 5}) {
		t.Fatalf("expected, got")
	}
	if -1 != binarySearch(2, []int{1, 3, 5}) {
		t.Fatalf("expected, got")
	}
	if -1 != binarySearch(4, []int{1, 3, 5}) {
		t.Fatalf("expected, got")
	}
	if -1 != binarySearch(6, []int{1, 3, 5}) {
		t.Fatalf("expected, got")
	}

	if 0 != binarySearch(1, []int{1, 3, 5, 7}) {
		t.Fatalf("expected, got")
	}
	if 1 != binarySearch(3, []int{1, 3, 5, 7}) {
		t.Fatalf("expected, got")
	}
	if 2 != binarySearch(5, []int{1, 3, 5, 7}) {
		t.Fatalf("expected, got")
	}
	if 3 != binarySearch(7, []int{1, 3, 5, 7}) {
		t.Fatalf("expected, got")
	}
	if -1 != binarySearch(0, []int{1, 3, 5, 7}) {
		t.Fatalf("expected, got")
	}
	if -1 != binarySearch(2, []int{1, 3, 5, 7}) {
		t.Fatalf("expected, got")
	}
	if -1 != binarySearch(4, []int{1, 3, 5, 7}) {
		t.Fatalf("expected, got")
	}
	if -1 != binarySearch(6, []int{1, 3, 5, 7}) {
		t.Fatalf("expected, got")
	}
	if -1 != binarySearch(8, []int{1, 3, 5, 7}) {
		t.Fatalf("expected, got")
	}
}
