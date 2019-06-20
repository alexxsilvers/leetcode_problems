package main

import "log"

// Given an array A of strings made only from lowercase letters,
// return a list of all characters that show up in all strings within the list (including duplicates).
// For example, if a character occurs 3 times in all strings but not 4 times, you need to include that
// character three times in the final answer.
//
// You may return the answer in any order.
func main() {
	log.Println(commonChars([]string{"bella","label","roller"})) // ["e","l","l"]
	log.Println(commonChars([]string{"cool","lock","cook"})) // ["c","o"]
	log.Println(commonChars([]string{"acabcddd","bcbdbcbd","baddbadb","cbdddcac","aacbcccd","ccccddda","cababaab","addcaccd"})) // []
}

type CharInWords struct {
	words int
	cnt int
}
func commonChars(A []string) []string {
	charMap := make(map[rune]*CharInWords)
	for _, word := range A {
		charInWord := make(map[rune]int)
		for _, char := range word {
			if _, ok := charInWord[char]; !ok {
				charInWord[char] = 0
			}
			charInWord[char]++
		}

		for char, cnt := range charInWord {
			charInWords, ok := charMap[char]
			if ok {
				charInWords.words++
				if cnt < charInWords.cnt {
					charInWords.cnt = cnt
				}
			} else {
				charMap[char] = &CharInWords{
					words: 1,
					cnt: cnt,
				}
			}
		}
	}

	result := make([]string, 0)
	for char, charInWords := range charMap {
		if charInWords.words == len(A) {
			for i := 0; i < charInWords.cnt; i++ {
				result = append(result, string(char))
			}
		}
	}

	return result
}
