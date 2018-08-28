package main

import (
	"fmt"
	"strings"
)

func uncommonFromSentences(A string, B string) []string {
	uniqWords := make(map[string]int)
	for _, word := range strings.Split(A, " ") {
		_, exist := uniqWords[word]
		if exist {
			uniqWords[word]--
		} else {
			uniqWords[word]++
		}
	}

	for _, word := range strings.Split(B, " ") {
		_, exist := uniqWords[word]
		if exist {
			uniqWords[word]--
		} else {
			uniqWords[word]++
		}
	}

	result := make([]string, 0)
	for word, cnt := range uniqWords {
		if cnt > 0 {
			result = append(result, word)
		}
	}

	return result
}

func main() {
	fmt.Println(uncommonFromSentences("this apple is sweet", "this apple is sour"))
	fmt.Println(uncommonFromSentences("apple apple", "banana"))
}
