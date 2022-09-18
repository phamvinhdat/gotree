package btree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNode_splitChild(t *testing.T) {
	// t = 3
	//        [3]
	//     /   |
	// [1][2] [4][5][6][7][8]
	//
	// ->    [3][6]
	//     /    |   \
	// [1][2] [4][5] [7][8]

	give := &Node{
		keys: []int{3},
		children: []*Node{
			{
				keys: []int{1, 2},
			},
			{
				keys: []int{4, 5, 6, 7, 8},
			},
		},
	}

	want := &Node{
		keys: []int{3, 6},
		children: []*Node{
			{
				keys: []int{1, 2},
			},
			{
				keys: []int{4, 5},
			},
			{
				keys: []int{7, 8},
			},
		},
	}

	splitChild(give, 3, 1)
	assert.Equal(t, want, give)
}
