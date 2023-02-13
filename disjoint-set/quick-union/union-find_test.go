package unionfind_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	unionfind "github.com/ytanne/dstools/disjoint-set/quick-union"
)

func TestUnionFind(t *testing.T) {
	uf := unionfind.Init(10)
	uf.Union(1, 2)
	uf.Union(2, 5)
	uf.Union(5, 6)
	uf.Union(6, 7)
	uf.Union(3, 8)
	uf.Union(8, 9)
	require.True(t, uf.Connected(1, 5))
	require.True(t, uf.Connected(5, 7))
	require.False(t, uf.Connected(4, 9))

	uf.Union(9, 4)
	require.True(t, uf.Connected(4, 9))
}
