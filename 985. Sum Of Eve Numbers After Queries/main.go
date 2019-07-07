package main

import "log"

// We have an array A of integers, and an array queries of queries.
//
// For the i-th query val = queries[i][0], index = queries[i][1], we add val to A[index].
// Then, the answer to the i-th query is the sum of the even values of A.
//
// (Here, the given index = queries[i][1] is a 0-based index, and each query permanently modifies the array A.)
//
// Return the answer to all queries.  Your answer array should have answer[i] as the answer to the i-th query.
func main() {
	log.Println(sumEvenAfterQueries([]int{1, 2, 3, 4, 5}, [][]int{{2, 0}, {-8, 0}, {6, 0}, {2, 3}})) // [6,6,6,2]
	log.Println(sumEvenAfterQueries([]int{0}, [][]int{{1, 0}, {-1, 0}}))                             // [0,0]
	log.Println(sumEvenAfterQueries([]int{1, 2, 3, 4}, [][]int{{1, 0}, {-3, 1}, {-4, 0}, {2, 3}}))   // [8,6,2,4]
}

func sumEvenAfterQueries(A []int, queries [][]int) []int {
	result := make([]int, len(queries))
	sum := 0
	for _, num := range A {
		if num&1 == 0 {
			sum += num
		}
	}

	for i, query := range queries {
		if A[query[1]]&1 == 0 {
			sum -= A[query[1]]
		}

		A[query[1]] += query[0]

		if A[query[1]]&1 == 0 {
			sum += A[query[1]]
		}
		result[i] = sum
	}
	return result
}
