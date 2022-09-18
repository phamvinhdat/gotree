package btree

import (
	"sort"
)

type Tree struct {
	root *Node
	t    int
}

func New(t int) *Tree {
	return &Tree{t: t}
}

func (t *Tree) Traverse() {
	t.root.traverse()
}

func (t *Tree) Insert(k int) {
	if t.root == nil {
		t.root = new(Node)
	}

	if len(t.root.keys) >= 2*t.t-1 {
		t.root = splitRoot(t.root, t.t)
	}

	t.root.insertNonFull(t.t, k)
}

func (t *Tree) Remove(k int) *Tree {
	if t.root == nil {
		return t
	}

	index := -1
	isEdgeCase := len(t.root.keys) == 1 && t.root.keys[0] == k
	if isEdgeCase {
		if t.root.isLeaf() {
			t.root = nil
			return t
		}

		leftChild := t.root.children[0]
		rightChild := t.root.children[1]
		if len(leftChild.keys) < t.t && len(rightChild.keys) < t.t {
			index = len(leftChild.keys)
			leftChild.keys = append(leftChild.keys, k)
			leftChild.keys = append(leftChild.keys, rightChild.keys...)
			if !leftChild.isLeaf() {
				leftChild.children = append(leftChild.children, rightChild.children...)
			}

			t.root = leftChild
		}
	}

	t.remove(k, index)
	return t
}

func (t *Tree) remove(k int, index int) {
	if index < 0 {
		index = sort.SearchInts(t.root.keys, k)
	}

	if index < len(t.root.keys) && t.root.keys[index] == k {
		t.root.removeAt(t.t, index)
		return
	}

	if t.root.isLeaf() {
		return
	}

	next := t.root.children[index]
	nextTree := t.softClone(func(root *Node) *Node {
		return next
	})
	if len(next.keys) < t.t {
		// one child has at least t keys
		siblingIndex := t.root.findSiblingHasAtLestTKeys(t.t, index)
		if siblingIndex >= 0 {
			// right rotate
			if siblingIndex == index-1 {
				t.root.rightRotate(index)
				nextTree.remove(k, -1)
				return
			}

			// left rotate
			t.root.leftRotate(index)
			nextTree.remove(k, -1)
			return
		}

		// two children have t - 1 keys
		mergedNode := t.root.merge(index)
		nextTree = t.softClone(func(root *Node) *Node {
			return mergedNode
		})
	}

	nextTree.remove(k, -1)
}

func (t *Tree) softClone(hook func(root *Node) *Node) *Tree {
	root := *t.root
	return &Tree{
		root: hook(&root),
		t:    t.t,
	}
}

func (t *Tree) String() string {
	return t.root.String()
}

func (t *Tree) Len() int {
	return t.root.len()
}
