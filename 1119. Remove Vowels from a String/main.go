package main

import (
	"log"
	"strings"
)

// Given a string S, remove the vowels 'a', 'e', 'i', 'o', and 'u' from it, and return the new string.
func main() {
	log.Println(removeVowels("leetcodeisacommunityforcoders")) // ltcdscmmntyfrcdrs
	log.Println(removeVowels("aeiou"))                         // []
}

func removeVowels(S string) string {
	removeMap := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}

	sb := strings.Builder{}
	sb.Grow(len(S))
	for i := 0; i < len(S); i++ {
		if !removeMap[S[i]] {
			sb.WriteByte(S[i])
		}
	}

	return sb.String()
}
