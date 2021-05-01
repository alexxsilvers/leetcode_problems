package network_connection_problem

import "strings"

// Social network connectivity. Given a social network containing nn members and a log file containing mm
// timestamps at which times pairs of members formed friendships, design an algorithm to determine the earliest
// time at which all members are connected (i.e., every member is a friend of a friend of a friend ... of a friend).
// Assume that the log file is sorted by timestamp and that friendship is an equivalence relation. The running time of
// your algorithm should be m log nmlogn or better and use extra space proportional to nn.

func earliestConnectTime(members []string, log []string) string {
	uf := NewUF(members)

	for _, l := range log {
		logString := strings.Split(l, " ")
		ts := logString[0]
		first := logString[1]
		second := logString[2]

		rank := uf.union(first, second)
		if rank == len(members) {
			return ts
		}
	}

	return ""
}

type UF struct {
	buf    []int
	rank   []int
	maping map[string]int
}

func NewUF(members []string) UF {
	uf := UF{
		buf:    make([]int, len(members)),
		rank:   make([]int, len(members)),
		maping: make(map[string]int, len(members)),
	}

	for i := 0; i < len(members); i++ {
		uf.buf[i] = i
		uf.rank[i] = 1
		uf.maping[members[i]] = i
	}

	return uf
}

func (uf *UF) union(a, b string) int {
	p := uf.maping[a]
	q := uf.maping[b]

	pRoot := uf.find(p)
	qRoot := uf.find(q)
	pRank := uf.rank[pRoot]
	qRank := uf.rank[qRoot]

	if pRank < qRank {
		uf.buf[pRoot] = qRoot
		uf.rank[qRoot] += uf.rank[pRoot]
		return uf.rank[qRoot]
	} else {
		uf.buf[qRoot] = pRoot
		uf.rank[pRoot] += uf.rank[qRoot]
		return uf.rank[pRoot]
	}
}

func (uf *UF) find(val int) int {
	for val != uf.buf[val] {
		uf.buf[val] = uf.buf[uf.buf[val]]
		val = uf.buf[val]
	}
	return val
}
