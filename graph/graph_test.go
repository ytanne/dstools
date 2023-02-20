package graph_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/ytanne/dstools/graph"
)

func TestGraph(t *testing.T) {
	g := graph.NewGraph(5)
	g.AddEdge(4, 1)
	g.AddEdge(1, 2)
	g.AddEdge(4, 3)
	g.AddEdge(2, 0)

	result := []int{
		4, 3, 1, 2, 0,
	}
	require.Equal(t, result, g.TopologicalSort())
}
