package selectsort_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	ss "github.com/ytanne/dstools/sorting/selection"
)

func TestSort(t *testing.T) {
	testCases := []struct {
		input  []int
		output []int
	}{
		{
			[]int{1, 2, 3},
			[]int{1, 2, 3},
		},
		{
			[]int{3, 2, 1},
			[]int{1, 2, 3},
		},
		{
			[]int{1, 5, 3, 3, 6},
			[]int{1, 3, 3, 5, 6},
		},
	}

	for _, tc := range testCases {
		ss.Sort(tc.input)
		require.Equal(t, tc.output, tc.input)
	}
}
