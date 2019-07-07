package main

import "fmt"

var lets = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

func uniqueMorseRepresentations(words []string) int {
	transformations := make(map[string]int)
	for _, word := range words {
		var morseStr string
		for _, l := range word {
			morseStr += lets[l-'a']
		}
		transformations[morseStr]++
	}
	return len(transformations)
}

func main() {
	fmt.Println(uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"}))
}
