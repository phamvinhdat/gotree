package avl

import (
	"github.com/phamvinhdat/gotree/x/xconstraints"
	"github.com/phamvinhdat/gotree/x/xmath"
)

type Tree[T xconstraints.Comparable[T]] struct {
	value  T
	height int
	left   *Tree[T]
	right  *Tree[T]
}

func New[T xconstraints.Comparable[T]](values ...T) *Tree[T] {
	if len(values) == 0 {
		return nil
	}

	var root *Tree[T]
	for _, value := range values {
		root = root.Insert(value)
	}

	return root
}

func (t *Tree[T]) Insert(value T) *Tree[T] {
	// 1. perform a normal BST insertion
	if t == nil {
		t = &Tree[T]{
			height: 1,
			value:  value,
		}

		return t
	}

	switch value.Compare(t.value) {
	case 1:
		t.right = t.right.Insert(value)
	case -1:
		t.left = t.left.Insert(value)
	default:
		return t
	}

	// 2. update height of ancestor tree
	t.height = t.Height()

	// 3. get the balance factor of this ancestor
	balance := t.getBalance()

	// if this tree became unbalanced, then there are four cases
	// 3.1. left-left case
	if balance > 1 && value.Compare(t.left.value) == -1 {
		return t.rightRotation()
	}

	// 3.2. right-right case
	if balance < -1 && value.Compare(t.left.value) == 1 {
		return t.leftRotation()
	}

	// 3.3. left-right case
	if balance > 1 && value.Compare(t.left.value) == 1 {
		t.left = t.left.leftRotation()
		return t.rightRotation()
	}

	// 3.4. right-left case
	if balance < -1 && value.Compare(t.left.value) == 1 {
		t.right = t.right.rightRotation()
		return t.leftRotation()
	}

	return t
}

func (t *Tree[T]) LRN(fn func(val T)) {
	if t.left != nil {
		t.left.LRN(fn)
	}

	fn(t.value)

	if t.right != nil {
		t.right.LRN(fn)
	}
}

func (t *Tree[T]) RNL(fn func(val T)) {
	if t.right != nil {
		t.right.RNL(fn)
	}

	fn(t.value)

	if t.left != nil {
		t.left.RNL(fn)
	}
}

func (t *Tree[T]) NLR(fn func(val T)) {
	fn(t.value)

	if t.left != nil {
		t.left.NLR(fn)
	}

	if t.right != nil {
		t.right.NLR(fn)
	}
}

func (t *Tree[T]) Equal(tree *Tree[T]) bool {
	if tree == nil {
		return t == nil
	}

	if t.value.Compare(tree.value) != 0 {
		return false
	}

	if t.left != nil {
		t.left.Equal(tree.left)
	}

	if t.right != nil {
		t.right.Equal(tree.right)
	}

	return true
}

func (t *Tree[T]) Height() int {
	if t == nil {
		return 0
	}

	return 1 + xmath.Max(t.left.Height(), t.right.Height())
}

func (t *Tree[T]) getBalance() int {
	if t == nil {
		return 0
	}

	return t.left.Height() - t.right.Height()
}

func (t *Tree[T]) leftRotation() *Tree[T] {
	if t.right == nil {
		return t
	}

	root := *t
	newRoot := root.right
	root.right = newRoot.left
	newRoot.left = &root

	*t = *newRoot
	return t
}

func (t *Tree[T]) rightRotation() *Tree[T] {
	if t.left == nil {
		return t
	}

	root := *t
	newRoot := root.left
	root.left = newRoot.right
	newRoot.right = &root

	*t = *newRoot
	return t
}
