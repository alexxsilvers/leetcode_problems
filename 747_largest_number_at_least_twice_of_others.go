package main

import "log"

func main() {
	log.Println(dominantIndex([]int{3, 6, 1, 0})) // 1
	log.Println(dominantIndex([]int{1, 2, 3, 4})) // -1
	log.Println(dominantIndex([]int{1})) // 0
	log.Println(dominantIndex([]int{1, 0})) // 0
}

func dominantIndex(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	largest := 0
	secondLargest := 0
	key := 0

	for i := range nums {
		if nums[i] > largest {
			secondLargest = largest
			largest = nums[i]
			key = i
		} else if nums[i] > secondLargest {
			secondLargest = nums[i]
		}
	}

	if secondLargest * 2 <= largest {
		return key
	}

	return -1
}
