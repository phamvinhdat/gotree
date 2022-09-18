package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveAt(t *testing.T) {
	t.Parallel()

	tests := []struct {
		giveArr   []int
		giveIndex int
		want      []int
	}{
		{
			giveArr:   []int{1, 2, 3},
			giveIndex: 0,
			want:      []int{2, 3},
		},
		{
			giveArr:   []int{1, 2, 3},
			giveIndex: 1,
			want:      []int{1, 3},
		},
		{
			giveArr:   []int{1, 2, 3},
			giveIndex: 2,
			want:      []int{1, 2},
		},
	}

	for _, test := range tests {
		got := RemoveAt(test.giveArr, test.giveIndex)
		assert.Equal(t, test.want, got)
	}
}
