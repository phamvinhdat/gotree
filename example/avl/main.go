package main

import (
	"fmt"

	"github.com/phamvinhdat/gotree/avl"
)

type IntComparable int

func (a IntComparable) Compare(b IntComparable) int {
	if a < b {
		return -1
	}

	if a > b {
		return 1
	}

	return 0
}

func main() {
	t := avl.New(IntComparable(10), IntComparable(6), IntComparable(15), IntComparable(4))
	t = t.Insert(IntComparable(5))
	t.NLR(func(val IntComparable) {
		fmt.Println(val)
	})
}
