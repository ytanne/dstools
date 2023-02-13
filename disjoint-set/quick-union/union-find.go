package unionfind

type UnionFind interface {
	Find(x int) int
	Union(x, y int)
	Connected(x, y int) bool
}

type unionFind struct {
	root []int
}

func Init(size int) *unionFind {
	root := make([]int, size)
	for i := 0; i < size; i++ {
		root[i] = i
	}
	return &unionFind{
		root: root,
	}
}

func (uf *unionFind) Find(x int) int {
	for x != uf.root[x] {
		x = uf.root[x]
	}
	return uf.root[x]
}

func (uf *unionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	if rootX != rootY {
		uf.root[rootY] = rootX
	}
}

func (uf *unionFind) Connected(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}
