package slice

import "GTL/internal/slice"

func Delete[T any](ts []T, index int) ([]T, error) {
	res, _, err := slice.Delete[T](ts, index)
	return res, err
}

func FilterDelete[T any](ts []T, m func(idx int, src T) bool) []T {
	emptyPos := 0
	for idx := range ts {
		if m(idx, ts[idx]) {
			continue
		}
		ts[emptyPos] = ts[idx]
		emptyPos++
	}
	return ts[:emptyPos]
}
