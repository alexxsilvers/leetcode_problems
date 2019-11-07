package main

import "log"

// Write a function to find the longest common prefix string amongst an array of strings.
//
// If there is no common prefix, return an empty string "".
func main() {
	log.Println(longestCommonPrefix([]string{"flower", "flow"}))           // flow
	log.Println(longestCommonPrefix([]string{"flower", "flow", "flight"})) // fl
	log.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))    // ""
	log.Println(longestCommonPrefix([]string{"l", "la", "ld"}))            // "l"
	log.Println(longestCommonPrefix([]string{"a"}))                        // "a"
	log.Println(longestCommonPrefix([]string{"", "b"}))                    // ""
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	minWord := strs[0]
	for _, str := range strs {
		if len(str) < len(minWord) {
			minWord = str
		}
	}

	if len(minWord) == 0 {
		return ""
	}

	for i := 0; i < len(minWord); i++ {
		for _, str := range strs {
			if str[i] != minWord[i] {
				return minWord[:i]
			}
		}
	}

	return minWord
}
