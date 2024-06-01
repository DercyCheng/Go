package slice

import (
	"testing"
)

func TestAdd(t *testing.T) {
	slice := []int{1, 2, 3}
	element := 4
	result := Add(slice, element)
	expected := []int{1, 2, 3, 4}

	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Add(%v, %d) = %v; want %v", slice, element, result, expected)
		}
	}
}
