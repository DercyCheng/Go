package slices

import (
	"fmt"
	"testing"
)

func TestContains(t *testing.T) {
	slice := []int{1, 2, 3, 4}
	if !Contains(slice, 3) {
		fmt.Println("NOT IN")
	}
	if !Contains(slice, 5) {
		fmt.Println("NOT IN")
	}
}
