package btree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTree_Insert(t *testing.T) {
	give := &Tree{
		t: 100,
	}

	for i := 0; i < 1000; i++ {
		give.Insert(i)
	}

	assert.Equal(t, 1000, give.Len())
}

func TestTree_Remove(t *testing.T) {
	tests := []struct {
		giveTree *Tree
		giveKey  int
		want     *Tree
	}{
		{
			giveTree: &Tree{
				t: 3,
				root: &Node{
					keys: []int{1},
				},
			},
			giveKey: 1,
			want: &Tree{
				root: nil,
				t:    3,
			},
		},
		{
			giveTree: &Tree{
				root: &Node{
					keys: []int{
						1, 2, 3,
					},
				},
				t: 3,
			},
			giveKey: 2,
			want: &Tree{
				root: &Node{
					keys: []int{
						1, 3,
					},
				},
				t: 3,
			},
		},
		{
			giveTree: &Tree{
				root: &Node{
					keys: []int{
						8,
					},
					children: []*Node{
						{
							keys: []int{2, 5},
							children: []*Node{
								{
									keys: []int{0, 1},
								},
								{
									keys: []int{3, 4},
								},
								{
									keys: []int{6, 7},
								},
							},
						},
						{
							keys: []int{11, 14, 17, 20, 26},
							children: []*Node{
								{
									keys: []int{9, 10},
								},
								{
									keys: []int{12, 13},
								},
								{
									keys: []int{15, 16},
								},
								{
									keys: []int{18, 19},
								},
								{
									keys: []int{22, 24, 25},
								},
								{
									keys: []int{27, 28, 29},
								},
							},
						},
					},
				},
				t: 3,
			},
			giveKey: 4,
			want: &Tree{
				root: &Node{
					keys: []int{
						11,
					},
					children: []*Node{
						{
							keys: []int{2, 8},
							children: []*Node{
								{
									keys: []int{0, 1},
								},
								{
									keys: []int{3, 5, 6, 7},
								},
								{
									keys: []int{9, 10},
								},
							},
						},
						{
							keys: []int{14, 17, 20, 26},
							children: []*Node{
								{
									keys: []int{12, 13},
								},
								{
									keys: []int{15, 16},
								},
								{
									keys: []int{18, 19},
								},
								{
									keys: []int{22, 24, 25},
								},
								{
									keys: []int{27, 28, 29},
								},
							},
						},
					},
				},
				t: 3,
			},
		},
		{
			giveTree: &Tree{
				root: &Node{
					keys: []int{
						11,
					},
					children: []*Node{
						{
							keys: []int{2, 8},
							children: []*Node{
								{
									keys: []int{0, 1},
								},
								{
									keys: []int{3, 5, 6, 7},
								},
								{
									keys: []int{9, 10},
								},
							},
						},
						{
							keys: []int{17, 20, 26},
							children: []*Node{
								{
									keys: []int{13, 14, 15, 16},
								},
								{
									keys: []int{18, 19},
								},
								{
									keys: []int{22, 24, 25},
								},
								{
									keys: []int{27, 28, 29},
								},
							},
						},
					},
				},
				t: 3,
			},
			giveKey: 8,
			want: &Tree{
				root: &Node{
					keys: []int{
						17,
					},
					children: []*Node{
						{
							keys: []int{2, 7, 11},
							children: []*Node{
								{
									keys: []int{0, 1},
								},
								{
									keys: []int{3, 5, 6},
								},
								{
									keys: []int{9, 10},
								},
								{
									keys: []int{13, 14, 15, 16},
								},
							},
						},
						{
							keys: []int{20, 26},
							children: []*Node{
								{
									keys: []int{18, 19},
								},
								{
									keys: []int{22, 24, 25},
								},
								{
									keys: []int{27, 28, 29},
								},
							},
						},
					},
				},
				t: 3,
			},
		},
	}

	for _, test := range tests {
		got := test.giveTree.Remove(test.giveKey)
		assert.Equal(t, test.want, got)
	}
}
