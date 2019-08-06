package main

import "log"

// In a string S of lowercase letters, these letters form consecutive groups of the same character.
//
// For example, a string like S = "abbxxxxzyy" has the groups "a", "bb", "xxxx", "z" and "yy".
//
// Call a group large if it has 3 or more characters.  We would like the starting and ending positions of every large group.
//
// The final answer should be in lexicographic order.
func main() {
	log.Println(largeGroupPositions("abbxxxxzyy"))        // [[3,6]]
	log.Println(largeGroupPositions("abc"))               // nil
	log.Println(largeGroupPositions("abcdddeeeeaabbbcd")) // [[3,5],[6,9],[12,14]]
	log.Println(largeGroupPositions("aaa"))               // [0, 2]
}

func largeGroupPositions(S string) [][]int {
	var res [][]int
	for i, j := 0, 0; i < len(S); i = j {
		for j < len(S) && S[i] == S[j] {
			j++
		}

		if j-i >= 3 {
			res = append(res, []int{i, j - 1})
		}
	}

	return res
}
