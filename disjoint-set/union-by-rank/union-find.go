package unionfind

type UnionFind interface {
	Find(x int) int
	Union(x, y int)
	Connected(x, y int) bool
}

type unionFind struct {
	root []int
	rank []int
}

func Init(size int) *unionFind {
	root := make([]int, size)
	for i := 0; i < size; i++ {
		root[i] = i
	}
	rank := make([]int, size)
	for i := 0; i < size; i++ {
		rank[i] = 1
	}
	return &unionFind{
		root: root,
		rank: rank,
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
		if uf.rank[rootX] > uf.rank[rootY] {
			uf.root[rootY] = rootX
		} else if uf.rank[rootX] < uf.rank[rootY] {
			uf.root[rootX] = rootY
		} else {
			uf.root[rootY] = rootX
			uf.rank[rootX]++
		}
	}
}

func (uf *unionFind) Connected(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}
