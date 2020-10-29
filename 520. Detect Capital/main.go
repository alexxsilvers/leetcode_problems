package main

import "log"

// Given a word, you need to judge whether the usage of capitals in it is right or not.
// We define the usage of capitals in a word to be right when one of the following cases holds:
// All letters in this word are capitals, like "USA".
// All letters in this word are not capitals, like "leetcode".
// Only the first letter in this word is capital, like "Google".
// Otherwise, we define that this word doesn't use capitals in a right way.
func main() {
	log.Println(detectCapitalUse("USA")) // true
	log.Println(detectCapitalUse("leetcode")) // true
	log.Println(detectCapitalUse("Google")) // true
	log.Println(detectCapitalUse("FlaG")) // false
}

// it is works only for ASCI
func detectCapitalUse(word string) bool {
	firstIsCapital := isCapital(word[0])
	capitalsCnt := 0

	for i := 1; i < len(word); i++ {
		if isCapital(word[i]) {
			capitalsCnt++
		}
	}

	if firstIsCapital && (capitalsCnt == len(word)-1 || capitalsCnt == 0) {
		return true
	} else if !firstIsCapital && capitalsCnt == 0 {
		return true
	}

	return false
}

func isCapital(char byte) bool {
	return char >= 65 && char  <= 90
}