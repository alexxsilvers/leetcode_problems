package main

import (
	"log"
	"math"
)

// Given an array A of strings made only from lowercase letters,
// return a list of all characters that show up in all strings within the list (including duplicates).
// For example, if a character occurs 3 times in all strings but not 4 times, you need to include that
// character three times in the final answer.
//
// You may return the answer in any order.
func main() {
	log.Println(commonChars([]string{"bella", "label", "roller"}))                                                                     // ["e","l","l"]
	log.Println(commonChars([]string{"cool", "lock", "cook"}))                                                                         // ["c","o"]
	log.Println(commonChars([]string{"acabcddd", "bcbdbcbd", "baddbadb", "cbdddcac", "aacbcccd", "ccccddda", "cababaab", "addcaccd"})) // []
}

func commonChars(A []string) []string {
	cnt := [26]int{}
	for i := range cnt {
		cnt[i] = math.MaxUint16
	}

	cntInWord := [26]int{}
	for _, word := range A {
		for _, char := range []byte(word) { // compiler trick - here we will not allocate new memory
			cntInWord[char-'a']++
		}

		for i := 0; i < 26; i++ {
			if cntInWord[i] < cnt[i] {
				cnt[i] = cntInWord[i]
			}
		}

		for i := range cntInWord {
			cntInWord[i] = 0
		}
	}

	result := make([]string, 0)
	for i := 0; i < 26; i++ {
		for j := 0; j < cnt[i]; j++ {
			result = append(result, string(i+'a'))
		}
	}

	return result
}
