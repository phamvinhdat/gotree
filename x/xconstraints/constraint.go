package xconstraints

type Comparable[T interface{}] interface {
	Compare(other T) int
}
