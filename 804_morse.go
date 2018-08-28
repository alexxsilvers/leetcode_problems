package main

import "fmt"

//var morseToLetters = map[string]string{
//	"a":".-",
//	"b":"-...",
//	"c":"-.-.",
//	"d":"-..",
//	"e":".",
//	"f":"..-.",
//	"g":"--.",
//	"h":"....",
//	"i":"..",
//	"j":".---",
//	"k":"-.-",
//	"l":".-..",
//	"m":"--",
//	"n":"-.",
//	"o":"---",
//	"p":".--.",
//	"q":"--.-",
//	"r":".-.",
//	"s":"...",
//	"t":"-",
//	"u":"..-",
//	"v":"...-",
//	"w":".--",
//	"x":"-..-",
//	"y":"-.--",
//	"z":"--..",
//}

var lets = []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}

func uniqueMorseRepresentations(words []string) int {
	transformations := make(map[string]int)
	for _, word := range words {
		var morseStr string
		for _, l := range word {
			morseStr += lets[l - 'a']
		}
		transformations[morseStr]++
	}
	return len(transformations)
}

func main() {
	fmt.Println(uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"}))
}