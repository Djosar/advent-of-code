package maths

import "golang.org/x/exp/constraints"

func Min[T constraints.Float | constraints.Signed](a, b T) T {
	if a > b {
		return b
	}
	return a
}
