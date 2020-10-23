package main

import (
	"bufio"
	"container/heap"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	t, _, _ := reader.ReadLine()
	dates, _, _ := reader.ReadLine()

	solution(string(t), strings.Split(string(dates), " "))
}

func solution(dateType string, dates []string) {
	
}

func tt(s string) int {
	cnt := 0
	memo := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		memo[s[i]]++
	}

	var priorityQueue PQ
	for _, val := range memo {
		priorityQueue = append(priorityQueue, val)
	}

	heap.Init(&priorityQueue)

	for priorityQueue.Len() > 0 {
		mostFrequent := priorityQueue.Top()
		priorityQueue.Pop()

		if priorityQueue.Len() == 0 {
			return cnt
		}

		if mostFrequent == priorityQueue.Top() {
			if mostFrequent > 1 {
				priorityQueue.Push(mostFrequent - 1)
			}
			cnt++
		}
	}

	return cnt
}

type PQ []int

func (pq PQ) Len() int {
	return len(pq)
}
func (pq PQ) Less(i, j int) bool {
	return pq[i] < pq[j]
}
func (pq PQ) Top() int {
	return pq[len(pq)-1]
}
func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
func (pq *PQ) Push(x interface{}) {
	item := x.(int)
	*pq = append(*pq, item)
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

//
//func specialArray(nums []int) int {
//	m := make(map[int]int)
//	for _, num := range nums {
//		m[num]++
//	}
//
//	x := len(nums)
//
//	for k, v := range m {
//		if k < x {
//			x -= v
//		}
//	}
//
//	return x
//}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//
//func minOperations(nums []int) int {
//	op := 0
//
//	sort.Ints(nums)
//
//	for _, n := range nums {
//		if n % 2 ==
//	}
//
//	return op
//}

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	memo := make(map[int]int, 6)
	for i := 0; i < n; i++ {
		memo[i] = 0
	}

	for _, edge := range edges {
		memo[edge[1]]++
	}

	res := make([]int, 0)
	for k, v := range memo {
		if v == 0 {
			res = append(res, k)
		}
	}

	return res
}

func thousandSeparator(n int) string {
	if n == 0 {
		return "0"
	}

	res := make([]byte, 0)
	tCounter := 0
	for n > 0 {
		res = append(res, byte(n%10)+'0')
		tCounter++

		if tCounter%3 == 0 {
			res = append(res, '.')
		}

		n = n / 10
	}

	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}

	if res[0] == '.' {
		res = res[1:]
	}

	return string(res)
}
