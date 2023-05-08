package unionfind

type UnionFind struct {
	root []int
}

func NewUnionFind(size int) *UnionFind {
	root := make([]int, size)
	for i := 0; i < size; i++ {
		root[i] = i
	}

	return &UnionFind{
		root: root,
	}
}

func (uf *UnionFind) Find(x int) int {
	if x == uf.root[x] {
		return x
	}

	uf.root[x] = uf.Find(uf.root[x])
	return uf.root[x]
}

func (uf *UnionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	if rootX != rootY {
		uf.root[rootY] = rootX
	}
}

func (uf *UnionFind) Connected(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}
