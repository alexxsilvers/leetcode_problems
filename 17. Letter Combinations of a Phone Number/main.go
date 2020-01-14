package main

import "log"

// Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.
// A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.
func main() {
	log.Println(letterCombinations("23")) // ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"]
	log.Println(letterCombinations("24")) // ["ag", "ah", "ai", "bg", "bh", "bi", "cg", "ch", "ci"]
	log.Println(letterCombinations(""))   // []
	log.Println(letterCombinations("2"))  // ["a", "b", "c"]
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	letters := map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}

	combinations := make([]string, 0)
	backtracking(letters, &combinations, []byte(digits), []byte{})

	return combinations
}

func backtracking(letters map[byte][]byte, combinations *[]string, digits []byte, str []byte) {
	if len(digits) == 0 {
		*combinations = append(*combinations, string(str))
		return
	}

	for _, char := range letters[digits[0]] {
		backtracking(letters, combinations, digits[1:], append(str, char))
	}
}
