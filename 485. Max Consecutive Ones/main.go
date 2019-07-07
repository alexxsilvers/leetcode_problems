package main

import "log"

// Given a binary array, find the maximum number of consecutive 1s in this array.
func main() {
	log.Println(findMaxConsecutiveOnes([]int{1, 1, 1, 1, 1, 1, 1})) // 7
	log.Println(findMaxConsecutiveOnes([]int{1}))                   // 1
	log.Println(findMaxConsecutiveOnes([]int{1, 1, 1, 0, 1}))       // 3
	log.Println(findMaxConsecutiveOnes([]int{0, 1, 1, 1, 0, 1}))    // 3
	log.Println(findMaxConsecutiveOnes([]int{1, 0, 1, 1}))          // 2
}

func findMaxConsecutiveOnes(nums []int) int {
	max, tmp := 0, 0
	for _, num := range nums {
		if num == 0 {
			if tmp > max {
				max = tmp
			}
			tmp = 0
			if max >= len(nums)/2 {
				return max
			}
		} else {
			tmp++
		}
	}
	if tmp > max {
		return tmp
	}
	return max
}
