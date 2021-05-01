package main

import (
	"log"
	"sort"
	"strconv"
	"strings"
)

// На вход дается последовательность из n чисел
// Вернуть строку которая выводит в порядке возрастания
// интервалы из чисел из последовательности
func main() {
	log.Println(getIntervalString([]int{1, 4, 5, 2, 3, 9, 8, 11})) // "1-5,8-9,11"
	log.Println(getIntervalString([]int{1, 4, 3, 2}))              // "1-4"
	log.Println(getIntervalString([]int{1, 4}))                    // "1,4"
	log.Println(getIntervalString([]int{10}))                      // "10"
	log.Println(getIntervalString([]int{}))                        // ""
}

func getIntervalString(n []int) string {
	sort.Ints(n)

	if len(n) == 0 {
		return ""
	}

	if len(n) == 1 {
		return strconv.Itoa(n[0])
	}

	result := make([]string, 0)
	l, r := 0, 1
	// [1,2,3,4,5,8,9,11]
	// [1,2,3,4]
	for r < len(n) {
		if n[r]-n[r-1] == 1 {
			r++
			continue
		}

		latest := r - 1
		if latest-l == 0 { // sequence contains 1 number
			result = append(result, strconv.Itoa(n[latest]))
		} else {
			result = append(result, strconv.Itoa(n[l])+"-"+strconv.Itoa(n[latest]))
		}

		l = r
		r++
	}

	return strings.Join(result, ",")
}
