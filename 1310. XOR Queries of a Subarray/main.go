package main

import "log"

// Given the array arr of positive integers and the array queries where
// queries[i] = [Li, Ri], for each query i compute the XOR of elements
// from Li to Ri (that is, arr[Li] xor arr[Li+1] xor ... xor arr[Ri] ).
// Return an array containing the result for the given queries.
func main() {
	log.Println(xorQueries([]int{1, 3, 4, 8}, [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}}))  // [2, 7, 14, 8]
	log.Println(xorQueries([]int{4, 8, 2, 10}, [][]int{{2, 3}, {1, 3}, {0, 0}, {0, 3}})) // [8, 0, 4, 4]
}

func xorQueries(arr []int, queries [][]int) []int {
	ret := make([]int, 0, len(queries))

	for _, q := range queries {
		if q[0] == q[1] {
			ret = append(ret, arr[q[0]])
			continue
		}

		res := arr[q[0]]
		for i := q[0] + 1; i <= q[1]; i++ {
			res = res ^ arr[i]
		}

		ret = append(ret, res)
	}

	return ret
}
