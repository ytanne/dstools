package unionfind

type UnionFind struct {
	rank []int
	root []int
}

func NewUnionFind(size int) *UnionFind {
	root := make([]int, size)
	rank := make([]int, size)

	for i := 0; i < size; i++ {
		root[i] = i
		rank[i] = 1
	}

	return &UnionFind{
		root: root,
		rank: rank,
	}
}

func (u *UnionFind) Find(x int) int {
	if u.root[x] == x {
		return x
	}

	u.root[x] = u.Find(u.root[x])
	return u.root[x]
}

func (u *UnionFind) Union(x, y int) {
	rootX := u.Find(x)
	rootY := u.Find(y)
	if rootX != rootY {
		if u.rank[rootX] > u.rank[rootY] {
			u.root[rootY] = rootX
		} else if u.rank[rootY] > u.rank[rootX] {
			u.root[rootX] = rootY
		} else {
			u.root[rootY] = rootX
			u.rank[rootX]++
		}
	}
}

func (u *UnionFind) Connected(x, y int) bool {
	return u.Find(x) == u.Find(y)
}
