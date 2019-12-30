package main

import (
	"container/heap"
	"log"
	"sort"
)

// Given a list of scores of different students, return the average score of each student's top
// five scores in the order of each student's id.
//
// Each entry items[i] has items[i][0] the student's id, and items[i][1] the student's score.
// The average score is calculated using integer division.
func main() {
	log.Println(highFive([][]int{{1, 91}, {1, 92}, {2, 93}, {2, 97}, {1, 60}, {2, 77}, {1, 65}, {1, 87}, {1, 100}, {2, 100}, {2, 76}}))   // [[1,87],[2,88]]
	log.Println(highFivePQ([][]int{{1, 91}, {1, 92}, {2, 93}, {2, 97}, {1, 60}, {2, 77}, {1, 65}, {1, 87}, {1, 100}, {2, 100}, {2, 76}})) // [[1,87],[2,88]]
}

func highFivePQ(items [][]int) [][]int {
	memo := make(map[int]PriorityQueue)
	studentsOrder := make([]int, 0)

	for _, item := range items {
		_, ok := memo[item[0]]
		if ok {
			memo[item[0]] = append(memo[item[0]], item[1])
		} else {
			priorityQueue := PriorityQueue{item[1]}
			memo[item[0]] = priorityQueue

			studentsOrder = append(studentsOrder, item[0])
		}
	}

	sort.Ints(studentsOrder)
	ret := make([][]int, 0, len(studentsOrder))
	for _, studentID := range studentsOrder {
		res := make([]int, 2)
		res[0] = studentID

		pq := memo[studentID]
		heap.Init(&pq)

		topScores := 5
		for topScores > 0 && pq.Len() > 0 {
			res[1] += heap.Pop(&pq).(int)
			topScores--
		}

		res[1] = res[1] / 5

		ret = append(ret, res)
	}

	return ret
}

func highFive(items [][]int) [][]int {
	memory := make(map[int][]int)
	order := make([]int, 0)

	for _, s := range items {
		_, ok := memory[s[0]]
		if ok {
			memory[s[0]] = append(memory[s[0]], s[1])
		} else {
			memory[s[0]] = []int{s[1]}
			order = append(order, s[0])
		}
	}

	sort.Ints(order)

	ret := make([][]int, 0, len(order))
	for _, sID := range order {
		scores := memory[sID]
		sort.Ints(scores)

		ret = append(ret, []int{sID, sum(scores[len(scores)-5:]) / 5})
	}

	return ret
}

func sum(a []int) int {
	s := 0
	for _, n := range a {
		s += n
	}

	return s
}
