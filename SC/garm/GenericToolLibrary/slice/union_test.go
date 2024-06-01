package slice

import "testing"

func TestUnion(t *testing.T) {
	slice1 := []int{1, 2, 3}
	slice2 := []int{3, 4, 5}
	res := Union(slice1, slice2)
	exp := []int{1, 2, 3, 4, 5}
	if len(res) != len((exp)) {
		t.Fatalf("Expected length %d, got %d", len(exp), len(res))
	}
	for i, v := range res {
		if v != exp[i] {
			t.Errorf("Expected %d, got %d", exp[i], v)
		}
	}
}
