package slice

import "GTL"

func Max[T GTL.RealNumber](ts []T) T {
	res := ts[0]
	for i := 1; i < len(ts); i++ {
		if ts[i] > res {
			res = ts[i]
		}
	}
	return res
}

func Min[T GTL.RealNumber](ts []T) T {
	res := ts[0]
	for i := 1; i < len(ts); i++ {
		if ts[i] < res {
			res = ts[i]
		}
	}
	return res
}

func Sum[T GTL.RealNumber](ts []T) T {
	var res T
	for _, n := range ts {
		res += n
	}
	return res
}
