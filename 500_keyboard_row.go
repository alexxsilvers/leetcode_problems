package main

import (
	"fmt"
	"strings"
)

func findWords(words []string) []string {
	rows := map[string]int{
		"q":1,
		"w":1,
		"e":1,
		"r":1,
		"t":1,
		"y":1,
		"u":1,
		"i":1,
		"o":1,
		"p":1,
		"a":2,
		"s":2,
		"d":2,
		"f":2,
		"g":2,
		"h":2,
		"j":2,
		"k":2,
		"l":2,
		"z":3,
		"x":3,
		"c":3,
		"v":3,
		"b":3,
		"n":3,
		"m":3,
	}

	result := make([]string, 0)
	for _, word := range words {
		lowerWord := strings.ToLower(word)
		firstCharRow := rows[string(lowerWord[0])]
		allSame := true
		for i := 1; i < len(lowerWord); i++ {
			if firstCharRow != rows[string(lowerWord[i])] {
				allSame = false
				break
			}
		}

		if allSame {
			result = append(result, word)
		}
	}

	return result
}

func main() {
	fmt.Printf("%#v\n", findWords([]string{"Hello", "Alaska", "Dad", "Peace", "NcVZ"}))
}
