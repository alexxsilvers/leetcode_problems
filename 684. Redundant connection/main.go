package main

import "log"

// In this problem, a tree is an undirected graph that is connected and has no cycles.
//
// The given input is a graph that started as a tree with N nodes (with distinct values 1, 2, ..., N),
// with one additional edge added. The added edge has two different vertices chosen from 1 to N,
// and was not an edge that already existed.
//
// The resulting graph is given as a 2D-array of edges. Each element of edges is a pair [u, v]
// with u < v, that represents an undirected edge connecting nodes u and v.
//
// Return an edge that can be removed so that the resulting graph is a tree of N nodes.
// If there are multiple answers, return the answer that occurs last in the given 2D-array.
// The answer edge [u, v] should be in the same format, with u < v.
func main() {
	log.Println(findRedundantConnection([][]int{{1, 2}, {1, 3}, {2, 3}}))                 // [2,3]
	log.Println(findRedundantConnection([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}})) // [1, 4]
}

func findRedundantConnection(edges [][]int) []int {
	dsu := newDSU()
	for _, edge := range edges {
		if !dsu.union(edge[0], edge[1]) {
			return edge
		}
	}
	return nil
}

type DSU struct {
	root []int
	rank []int
}

func (dsu *DSU) find(a int) int {
	if dsu.root[a] == a {
		return a
	} else {
		dsu.root[a] = dsu.find(dsu.root[a])
		return dsu.root[a]
	}
}
func (dsu *DSU) union(a, b int) bool {
	aR := dsu.find(a)
	bR := dsu.find(b)
	if aR == bR {
		return false
	} else if dsu.rank[aR] < dsu.rank[bR] {
		dsu.root[aR] = dsu.root[bR]
		dsu.rank[bR] += dsu.rank[aR]
	} else {
		dsu.root[bR] = dsu.root[aR]
		dsu.rank[aR] += dsu.rank[bR]
	}
	return true
}
func newDSU() *DSU {
	dsu := &DSU{
		root: make([]int, 1001),
		rank: make([]int, 1001),
	}
	for i := 0; i < 1001; i++ {
		dsu.root[i] = i
		dsu.rank[i] = 1
	}
	return dsu
}
