package maps

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	m := map[int]string{1: "a", 2: "b"}
	key := 3
	value := "c"
	result := Add(m, key, value)
	expected := map[int]string{1: "a", 2: "b", 3: "c"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Add(%v, %d, %s) = %v; want %v", m, key, value, result, expected)
	}
}
