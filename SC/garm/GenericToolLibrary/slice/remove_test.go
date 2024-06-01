package slice

import (
	"testing"
)

func TestRemove(t *testing.T) {
	slice := []int{1, 2, 3, 4}
	newSlice := Remove(slice, 3)
	expected := []int{1, 2, 4}

	if len(newSlice) != len(expected) {
		t.Fatalf("Expected length %d, got %d", len(expected), len(newSlice))
	}

	for i, v := range newSlice {
		if v != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], v)
		}
	}
}
