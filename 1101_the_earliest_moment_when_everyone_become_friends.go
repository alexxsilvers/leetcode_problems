package main

import (
	"log"
	"sort"
)

// In a social group, there are N people, with unique integer ids from 0 to N-1.
//
// We have a list of logs, where each logs[i] = [timestamp, id_A, id_B] contains a
// non-negative integer timestamp, and the ids of two different people.
//
// Each log represents the time in which two different people became friends.
// Friendship is symmetric: if A is friends with B, then B is friends with A.
//
// Let's say that person A is acquainted with person B if A is friends with B,
// or A is a friend of someone acquainted with B.
//
// Return the earliest time for which every person became acquainted with every
// other person. Return -1 if there is no such earliest time.
func main() {
	log.Println(earliestAcq([][]int{
		{20190101, 0, 1},
		{20190104, 3, 4},
		{20190107, 2, 3},
		{20190211, 1, 5},
		{20190224, 2, 4},
		{20190301, 0, 3},
		{20190312, 1, 2},
		{20190322, 4, 5},
	}, 6)) // 20190301

	log.Println(earliestAcq([][]int{
		{0, 2, 0},
		{1, 0, 1},
		{2, 0, 2},
		{3, 0, 3},
		{4, 1, 2},
		{5, 1, 2},
		{6, 1, 0},
		{7, 3, 1},
		{8, 0, 3},
		{9, 3, 0},
	}, 4)) // 3

	log.Println(earliestAcq([][]int{
		{0, 0, 1},
		{1, 1, 2},
		{2, 2, 3},
		{3, 2, 0},
	}, 5)) // -1

	log.Println(earliestAcq([][]int{
		{0, 0, 2},
		{1, 1, 2},
		{2, 0, 4},
		{3, 1, 4},
		{4, 1, 2},
		{5, 4, 3},
		{6, 1, 0},
		{7, 4, 0},
		{8, 2, 4},
		{9, 1, 3},
		{10, 0, 2},
	}, 5)) // 5
}

func earliestAcq(logs [][]int, N int) int {
	dsu := NewDisjointSetUnion(N)
	sort.Slice(logs, func(i, j int) bool {
		return logs[i][0] < logs[j][0]
	})

	for _, l := range logs {
		t := l[0]
		a := l[1]
		b := l[2]
		dsu.union(a, b)
		if dsu.maxRank == N {
			return t
		}
	}

	return -1
}

type DisjointSetUnion struct {
	root    []int
	rank    []int
	maxRank int
}

func NewDisjointSetUnion(N int) *DisjointSetUnion {
	set := make([]int, N)
	rank := make([]int, N)
	for i := 0; i < N; i++ {
		set[i] = i
		rank[i] = 1
	}
	return &DisjointSetUnion{root: set, rank: rank}
}

func (dsu *DisjointSetUnion) find(a int) int {
	if dsu.root[a] == a {
		return a
	} else {
		dsu.root[a] = dsu.find(dsu.root[a])
		return dsu.root[a]
	}
}

func (dsu *DisjointSetUnion) union(a int, b int) bool {
	aRep := dsu.find(a)
	bRep := dsu.find(b)
	if aRep == bRep {
		return false
	} else if dsu.rank[aRep] < dsu.rank[bRep] {
		dsu.root[aRep] = dsu.root[bRep]
		dsu.rank[bRep] += dsu.rank[aRep]
		if dsu.rank[bRep] > dsu.maxRank {
			dsu.maxRank = dsu.rank[bRep]
		}
	} else {
		dsu.root[bRep] = dsu.root[aRep]
		dsu.rank[aRep] += dsu.rank[bRep]
		if dsu.rank[aRep] > dsu.maxRank {
			dsu.maxRank = dsu.rank[aRep]
		}
	}
	return true
}
