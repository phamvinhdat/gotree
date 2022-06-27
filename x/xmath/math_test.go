package xmath_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/phamvinhdat/gotree/x/xmath"
)

type IntComparable int

func (a IntComparable) Compare(b IntComparable) int {
	if a > b {
		return 1
	}

	if a < b {
		return -1
	}

	return 0
}

func TestMax(t *testing.T) {
	t.Parallel()

	tests := []struct {
		a, b IntComparable
		want IntComparable
	}{
		{
			a:    1,
			b:    2,
			want: 2,
		},
		{
			a:    1,
			b:    1,
			want: 1,
		},
		{
			a:    2,
			b:    1,
			want: 2,
		},
	}

	for _, test := range tests {
		got := xmath.Max(test.a, test.b)
		assert.Equal(t, test.want, got)
	}
}
