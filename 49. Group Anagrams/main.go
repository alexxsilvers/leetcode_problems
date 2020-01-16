package main

import (
	"log"
	"sort"
)

// Given an array of strings, group anagrams together.
func main() {
	log.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	/*
		[
			["eat", "tea", "ate"],
			["tan", "nat"],
			["bat"],
		]
	*/
}

func groupAnagrams(strs []string) [][]string {
	grouped := make(map[string][]string)

	for _, str := range strs {
		strBytes := []byte(str)
		sort.Slice(strBytes, func(i, j int) bool {
			return strBytes[i] < strBytes[j]
		})

		sortedStr := string(strBytes)
		if _, exist := grouped[sortedStr]; exist {
			grouped[sortedStr] = append(grouped[sortedStr], str)
		} else {
			grouped[sortedStr] = []string{str}
		}
	}

	ret := make([][]string, 0)
	for _, group := range grouped {
		ret = append(ret, group)
	}
	return ret
}
