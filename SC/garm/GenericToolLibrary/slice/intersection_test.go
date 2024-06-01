package slice

import "testing"

func TestIntersection(t *testing.T) {
	slice1 := []int{1, 2, 3}
	slice2 := []int{3, 4, 5}
	res := Intersection(slice1, slice2)
	exp := []int{3}
	if len(res) != len(exp) {
		t.Fatalf("Expected length %d, got %d", len(exp), len(res))
	}
	for i, v := range res {
		if v != exp[i] {
			t.Errorf("Expected %d, got %d", exp[i], v)
		}
	}
}
