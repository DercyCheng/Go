package _5_binarysearch

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	a := []int{1, 3, 5, 6, 8}
	fmt.Println("index :", BinarySearch(a, 5))
}

func TestBinarySearchRecursive(t *testing.T) {
	a := []int{1, 3, 5, 6, 8}
	if BinarySearchRecursive(a, 3) != -1 {
		fmt.Println(BinarySearch(a, 3))
	}
	if BinarySearchRecursive(a, 4) == -1 {
		fmt.Println("not in arr")
	}
}

func TestBinarySearchFirst(t *testing.T) {
	var a []int

	a = []int{1, 2, 2, 2, 3, 4}
	if BinarySearchFirst(a, 2) != 1 {
		fmt.Println(BinarySearchFirst(a, 2))
	}

	a = []int{1, 2, 2, 2, 3, 4}
	if BinarySearchFirst(a, 3) != 4 {
		fmt.Println(BinarySearchFirst(a, 3))
	}
}

func TestBinarySearchLast(t *testing.T) {
	var a []int

	a = []int{1, 2, 2, 2, 3, 4}
	if BinarySearchLast(a, 2) != 3 {
		fmt.Println(BinarySearchLast(a, 2))
	}

	a = []int{1, 2, 2, 2, 3, 4}
	if BinarySearchLast(a, 3) != 4 {
		fmt.Println(BinarySearchLast(a, 3))
	}
}

func TestBinarySearchFirstGT(t *testing.T) {
	var a []int

	a = []int{1, 2, 2, 2, 3, 4}
	if BinarySearchFirstGT(a, 2) != 4 {
		fmt.Println(BinarySearchFirstGT(a, 2))
	}

	a = []int{1, 2, 2, 2, 3, 4}
	if BinarySearchFirstGT(a, 1) != 1 {
		fmt.Println(BinarySearchFirstGT(a, 1))
	}

	a = []int{1, 2, 2, 2, 3, 4}
	if BinarySearchFirstGT(a, 3) != 5 {
		fmt.Println(BinarySearchFirstGT(a, 3))
	}

	a = []int{1, 2, 2, 2, 3, 4}
	if BinarySearchFirstGT(a, 4) != -1 {
		fmt.Println(BinarySearchFirstGT(a, 4))
	}
}

func TestBinarySearchLastLT(t *testing.T) {
	var a []int

	a = []int{1, 2, 2, 2, 3, 4}
	if BinarySearchLastLT(a, 2) != 0 {
		fmt.Println(BinarySearchLastLT(a, 2))
	}

	a = []int{1, 2, 2, 2, 3, 4}
	if BinarySearchLastLT(a, 1) != -1 {
		fmt.Println(BinarySearchLastLT(a, 1))
	}

	a = []int{1, 2, 2, 2, 3, 4}
	if BinarySearchLastLT(a, 3) != 3 {
		fmt.Println(BinarySearchLastLT(a, 3))
	}

	a = []int{1, 2, 2, 2, 3, 4}
	if BinarySearchLastLT(a, 4) != 4 {
		fmt.Println(BinarySearchLastLT(a, 4))
	}
}
