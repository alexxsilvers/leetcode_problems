package main

import (
	"log"
	"strings"
)

// A sentence S is given, composed of words separated by spaces. Each word consists of lowercase and uppercase letters only.
//
// We would like to convert the sentence to "Goat Latin" (a made-up language similar to Pig Latin.)
//
// The rules of Goat Latin are as follows:
//
// If a word begins with a vowel (a, e, i, o, or u), append "ma" to the end of the word.
// For example, the word 'apple' becomes 'applema'.
//
// If a word begins with a consonant (i.e. not a vowel), remove the first letter and append it to the end, then add "ma".
// For example, the word "goat" becomes "oatgma".
//
// Add one letter 'a' to the end of each word per its word index in the sentence, starting with 1.
// For example, the first word gets "a" added to the end, the second word gets "aa" added to the end and so on.
// Return the final sentence representing the conversion from S to Goat Latin.
func main() {
	// Imaa peaksmaaa oatGmaaaa atinLmaaaaa
	log.Println(toGoatLatin("I speak Goat Latin"))

	// heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa
	log.Println(toGoatLatin("The quick brown fox jumped over the lazy dog"))

	// yDummmaa
	log.Println(toGoatLatin("yDumm"))
}

var vowelMap = map[rune]bool{
	'A': true,
	'a': true,
	'E': true,
	'e': true,
	'I': true,
	'i': true,
	'O': true,
	'o': true,
	'U': true,
	'u': true,
}

func toGoatLatin(S string) string {
	words := strings.Split(S, " ")
	maEnding := make([]byte, 0, 2+len(words))
	maEnding = append(maEnding, 'm', 'a', 'a')

	var builder strings.Builder
	for i, word := range words {
		if isVowel([]rune(word)[0]) {
			builder.WriteString(word)
			builder.Write(maEnding)
		} else {
			builder.WriteString(word[1:])
			builder.WriteString(word[:1])
			builder.Write(maEnding)
		}

		words[i] = builder.String()
		builder.Reset()
		maEnding = append(maEnding, 'a')
	}

	return strings.Join(words, " ")
}

func isVowel(char rune) bool {
	return vowelMap[char]
}
