package graph

import "github.com/ytanne/dstools/stack"

type Graph struct {
	nodes []int
	edges map[int][]int
	size  int
}

func NewGraph(size int) *Graph {
	return &Graph{
		nodes: make([]int, size),
		edges: make(map[int][]int, size),
		size:  size,
	}
}

func (g *Graph) AddNode(x int) {
	g.nodes = append(g.nodes, x)
	g.size++
}

func (g *Graph) AddEdge(x, y int) {
	g.edges[x] = append(g.edges[x], y)
}

func (g *Graph) TopologicalSort() []int {
	result := make([]int, 0, g.size)

	c := make([]int, g.size)
	for _, neighbors := range g.edges {
		for _, v := range neighbors {
			c[v]++
		}
	}

	s := stack.NewStack()

	for node, outcoming := range c {
		if outcoming == 0 {
			s.Push(node)
		}
	}

	visited := make([]int, g.size)
	for s.Size() != 0 {
		v, _ := s.Pop()
		idx := v.(int)
		if visited[idx] == 1 {
			return nil
		}
		if visited[idx] == 2 {
			continue
		}
		result = append(result, idx)
		visited[idx] = 1
		for _, node := range g.edges[idx] {
			c[node]--
			if c[node] == 0 {
				s.Push(node)
			}
		}
	}

	return result
}
