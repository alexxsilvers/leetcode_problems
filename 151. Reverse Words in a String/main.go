package main

import (
	"log"
	"strings"
)

// Given an input string, reverse the string word by word.
// A word is defined as a sequence of non-space characters.
// Input string may contain leading or trailing spaces. However, your reversed string should not contain
// leading or trailing spaces.
// You need to reduce multiple spaces between two words to a single space in the reversed string.
func main() {
	log.Println(reverseWords(" this is  a   string")) // "string a is this"
	log.Println(reverseWords("the sky is blue"))      // "blue is sky the"
	log.Println(reverseWords("  hello world!  "))     // "world! hello"
	log.Println(reverseWords("a good   example"))     // "example good a"
}

func reverseWords(s string) string {
	words := make([]string, 0)

	start, end := 0, 0
	for end < len(s) {
		if s[start] == ' ' {
			start++
			end++
			continue
		}

		if s[end] == ' ' {
			words = append(words, s[start:end])
			start = end
			end++
			continue
		}

		if end+1 == len(s) {
			words = append(words, s[start:end+1])
		}

		end++
	}

	for i := 0; i < len(words)/2; i++ {
		words[i], words[len(words)-1-i] = words[len(words)-1-i], words[i]
	}

	return strings.Join(words, " ")
}
