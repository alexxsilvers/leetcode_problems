package main

import (
	"log"
	"strings"
)

// Given a string, you need to reverse the order of characters in each word
// within a sentence while still preserving whitespace and initial word order.

// In the string, each word is separated by single space and there will not be any extra space in the string.
func main() {
	log.Println(reverseWords("Let's take LeetCode contest")) // "s'teL ekat edoCteeL tsetnoc"
	log.Println(reverseWords("I am cool prOgrámmer"))       // "I ma looc remmárgOrp"
}

func reverseWords(s string) string {
	var builder strings.Builder
	words := strings.Split(s, " ")
	result := make([]string, 0, len(words))

	for _, word := range words {
		wordRunes := []rune(word)
		for i := len(wordRunes) - 1; i >= 0; i-- {
			builder.WriteRune(wordRunes[i])
		}
		result = append(result, builder.String())
		builder.Reset()
	}

	return strings.Join(result, " ")
}