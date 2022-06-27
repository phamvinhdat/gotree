package xmath

import (
	"golang.org/x/exp/constraints"

	"github.com/phamvinhdat/gotree/x/xconstraints"
)

func MaxXComparable[T xconstraints.Comparable[T]](a, b T) T {
	if a.Compare(b) == 1 {
		return a
	}

	return b
}

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}

	return b
}
