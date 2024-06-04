package slices

import "testing"

func TestMap(t *testing.T) {
	slice := []int{1, 2, 3}
	res := Map(slice, func(x int) int {
		return x * 2
	})
	exp := []int{2, 4, 6}
	if len(res) != len(exp) {
		t.Fatalf("Expected length %d,got %d", len(exp), len(res))
	}
	for i, v := range res {
		if v != exp[i] {
			t.Errorf("Expected %d,got %d", exp[i], v)
		}
	}
}
