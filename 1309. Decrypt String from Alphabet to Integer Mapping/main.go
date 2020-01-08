package main

import (
	"log"
	"strings"
)

// Given a string s formed by digits ('0' - '9') and '#' . We want to map s to
// English lowercase characters as follows:
//    Characters ('a' to 'i') are represented by ('1' to '9') respectively.
//    Characters ('j' to 'z') are represented by ('10#' to '26#') respectively.
// Return the string formed after mapping.
// It's guaranteed that a unique mapping will always exist.
func main() {
	log.Println(freqAlphabets("10#11#12"))                                                     // "jkab"
	log.Println(freqAlphabets("1326#"))                                                        // "acz"
	log.Println(freqAlphabets("25#"))                                                          // "y"
	log.Println(freqAlphabets("12345678910#11#12#13#14#15#16#17#18#19#20#21#22#23#24#25#26#")) // "abcdefghijklmnopqrstuvwxyz"
}

func freqAlphabets(s string) string {
	sb := &strings.Builder{}
	i := 0
	for i < len(s) {
		if i < len(s)-2 && s[i+2] == '#' {
			sb.WriteByte('j' + (s[i]-'1')*10 + s[i+1] - '0')
			i += 3
		} else {
			sb.WriteByte('a' + s[i] - '1')
			i++
		}
	}

	return sb.String()
}
