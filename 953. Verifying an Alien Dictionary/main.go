package main

import "log"

// In an alien language, surprisingly they also use english lowercase letters, but possibly in a
// different order. The order of the alphabet is some permutation of lowercase letters.
// Given a sequence of words written in the alien language, and the order of the alphabet,
// return true if and only if the given words are sorted lexicographicaly in this alien language.
func main() {
	log.Println(isAlienSorted([]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz"))    /// true
	log.Println(isAlienSorted([]string{"word", "world", "row"}, "worldabcefghijkmnpqstuvxyz")) /// false
	log.Println(isAlienSorted([]string{"apple", "app"}, "worldabcefghijkmnpqstuvxyz"))         /// false
	log.Println(isAlienSorted([]string{"kuvp", "q"}, "ngxlkthsjuoqcpavbfdermiywz"))            /// true
}

func isAlienSorted(words []string, order string) bool {
	ind := make(map[byte]int, len(order))
	for i := 0; i < len(order); i++ {
		ind[order[i]] = i
	}

	for i := 0; i < len(words)-1; i++ {
		if !checkWords(words[i], words[i+1], ind) {
			return false
		}
	}

	return true
}

func checkWords(w1, w2 string, order map[byte]int) bool {
	l1, l2 := len(w1), len(w2)
	for k := 0; k < min(l1, l2); k++ {
		if w1[k] != w2[k] {
			if order[w1[k]] > order[w2[k]] {
				return false
			}

			return true
		}
	}

	if l1 > l2 {
		return false
	}

	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
