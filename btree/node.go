package btree

import (
	"fmt"
	"sort"
	"strings"

	"github.com/phamvinhdat/gotree/x/array"
)

type Node struct {
	keys     []int
	children []*Node
}

func (n *Node) removeAt(t, index int) {
	if n.isLeaf() {
		n.keys = array.RemoveAt(n.keys, index)
		return
	}

	if len(n.children[index].keys) >= t {
		next := n.children[index]
		predecessor := next.keys[len(next.keys)-1]
		next.keys[len(next.keys)-1] = n.keys[index]
		n.keys[index] = predecessor
		next.removeAt(t, len(next.keys)-1)
		return
	}

	// one child has at least t keys
	if siblingIndex := n.findSiblingHasAtLestTKeys(t, index); siblingIndex >= 0 {
		// right rotate
		if siblingIndex == index-1 {
			n.rightRotate(index)
			n.children[index+1].removeAt(t, 0)
			return
		}

		// left rotate
		n.leftRotate(index)
		n.children[index].removeAt(t, len(n.children[index].keys)-1)
		return
	}

	// no one have t key
	// merge child[index] +  [key] + child[index + 1] and remove
	leftSibling := n.children[index]
	rightSibling := n.children[index+1]
	k := n.keys[index]
	n.keys = array.RemoveAt(n.keys, index)
	n.children = array.RemoveAt(n.children, index+1)
	leftSibling.keys = append(leftSibling.keys, k)
	leftSibling.keys = append(leftSibling.keys, rightSibling.keys...)
	if !leftSibling.isLeaf() {
		leftSibling.children = append(leftSibling.children, rightSibling.children...)
	}

	leftSibling.removeAt(t, t)
}

func (n *Node) leftRotate(index int) {
	leftSibling := n.children[index]
	rightSibling := n.children[index+1] // has at least t keys
	movingSiblingKey := rightSibling.keys[0]
	rightSibling.keys = array.RemoveAt(rightSibling.keys, 0)
	leftSibling.keys = append(leftSibling.keys, n.keys[index])

	n.keys[index] = movingSiblingKey

	if !rightSibling.isLeaf() {
		leftSibling.children = array.InsertAt(
			leftSibling.children,
			len(leftSibling.children),
			rightSibling.children[0],
		)

		rightSibling.children = array.RemoveAt(rightSibling.children, 0)
	}
}

func (n *Node) rightRotate(index int) {
	leftSibling := n.children[index-1] // has at least t keys
	rightSibling := n.children[index]

	movingSiblingKey := leftSibling.keys[len(leftSibling.keys)-1]
	leftSibling.keys = array.RemoveAt(leftSibling.keys, len(leftSibling.keys)-1)
	rightSibling.keys = array.InsertAt(
		rightSibling.keys,
		0,
		n.keys[index-1],
	)

	n.keys[index-1] = movingSiblingKey

	if !leftSibling.isLeaf() {
		rightSibling.children = array.InsertAt(
			rightSibling.children,
			0,
			leftSibling.children[len(leftSibling.children)-1],
		)

		leftSibling.children = array.RemoveAt(
			leftSibling.children,
			len(leftSibling.children)-1,
		)
	}
}

func (n *Node) findSiblingHasAtLestTKeys(t, index int) int {
	right := index < len(n.keys)
	left := index > 0

	// left
	if left {
		if len(n.children[index-1].keys) >= t {
			return index - 1
		}
	}

	// right
	if right {
		if len(n.children[index+1].keys) >= t {
			return index + 1
		}
	}

	return -1
}

func splitRoot(n *Node, t int) *Node {
	root := &Node{
		children: []*Node{n},
	}

	splitChild(root, t, 0)
	return root
}

func (n *Node) insertNonFull(t, value int) {
	i := sort.SearchInts(n.keys, value)
	if i < len(n.keys) && n.keys[i] == value { // already exists
		return
	}

	if n.isLeaf() {
		n.keys = array.InsertAt(n.keys, i, value)
		return
	}

	if len(n.children[i].keys) == 2*t-1 { // full
		splitChild(n, t, i)
		if value > n.keys[i] { // ? tree change
			i++
		}
	}

	n.children[i].insertNonFull(t, value)
}

func splitChild(n *Node, t, index int) {
	splitNode := n.children[index] // full node to split
	halfNode := &Node{
		keys: splitNode.keys[t:], // halfNode gets splitNode's greatest keys
	}

	if !splitNode.isLeaf() {
		halfNode.children = splitNode.children[t:]
		splitNode.children = splitNode.children[:t]
	}

	n.keys = array.InsertAt(n.keys, index, splitNode.keys[t-1])
	n.children = array.InsertAt(n.children, index+1, halfNode)
	splitNode.keys = splitNode.keys[:t-1]
}

func (n *Node) traverse() {
	for i, child := range n.children {
		if len(child.children) != 0 {
			child.traverse()
		}

		fmt.Println(n.keys[i])
	}
}

func (n *Node) isLeaf() bool {
	return len(n.children) == 0
}

func (n *Node) String() string {
	if n == nil {
		return ""
	}

	var str strings.Builder
	for _, nodes := range n.Array() {
		for i, node := range nodes {
			str.WriteString(strKeys(node))
			if i != len(nodes)-1 {
				str.WriteString(" | ")
			}
		}

		str.WriteString("\n")
	}

	return str.String()
}

// Array convert note to an array. format: [level][node][keys]
func (n *Node) Array() [][][]int {
	var arr [][][]int
	n.separateByLevel(0, &arr)
	return arr
}

func (n *Node) separateByLevel(level int, result *[][][]int) {
	if len(*result) < level+1 {
		*result = append(*result, [][]int{})
	}

	temp := *result
	(temp[level]) = append(temp[level], n.keys)
	*result = temp
	for _, child := range n.children {
		child.separateByLevel(level+1, result)
	}
}

// merge and return a merged *Node
func (n *Node) merge(index int) *Node {
	leftIndex := index
	rightIndex := index + 1
	if leftIndex == len(n.keys) { // edge
		leftIndex = index - 1
		rightIndex = index
	}

	leftChild := n.children[leftIndex]
	rightChild := n.children[rightIndex]
	median := n.keys[leftIndex]
	n.keys = array.RemoveAt(n.keys, leftIndex)
	leftChild.keys = append(leftChild.keys, median)
	leftChild.keys = append(leftChild.keys, rightChild.keys...)
	if !leftChild.isLeaf() {
		leftChild.children = append(leftChild.children, rightChild.children...)
	}

	if len(n.keys) == 0 {
		*n = *leftChild
		return n
	}

	n.children = array.RemoveAt(n.children, rightIndex)
	return leftChild
}

func strKeys(keys []int) string {
	var str strings.Builder
	for _, key := range keys {
		str.WriteString(fmt.Sprintf("[%d]", key))
	}

	return str.String()
}

func (n *Node) len() int {
	if n == nil {
		return 0
	}

	l := 0
	for _, child := range n.children {
		l += child.len()
	}

	return l + len(n.keys)
}
