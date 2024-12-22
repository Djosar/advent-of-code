package maths

import "golang.org/x/exp/constraints"

// Abs returns the absolute value of a numeric type.
func Abs[T constraints.Signed | constraints.Float](x T) T {
	if x < 0 {
		return -x
	}
	return x
}
