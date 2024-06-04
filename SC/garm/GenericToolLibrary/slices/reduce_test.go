package slices

import "testing"

func TestReduce(t *testing.T) {
	slice := []int{1, 2, 3, 4}
	res := Reduce(slice, 0, func(acc, x int) int {
		return acc + x
	})
	exp := 10
	if res != exp {
		t.Errorf("Expected %d,got %d", exp, res)
	}
}
