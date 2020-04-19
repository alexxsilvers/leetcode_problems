package main

import (
	"log"
	"sort"
)

func main() {
	//log.Println(arrayRankTransform([]int{40, 10, 20, 30}))                    // [4,1,2,3]
	//log.Println(arrayRankTransform([]int{100, 100, 100}))                     // [1,1,1]
	//log.Println(arrayRankTransform([]int{37, 12, 28, 9, 100, 56, 80, 5, 12})) // [5,3,4,2,8,6,7,1,3]


	log.Println(mergeSort([]int{10,9,8,7,6,5,4,3,2,1}))
}

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	for j := 1; j < len(nums); j++ {
		key := nums[j]
		i := j-1

		for i >= 0 && nums[i] > key {
			nums[i+1] = nums[i]
			i--
		}

		nums[i+1] = key
	}

	return nums
}

func arrayRankTransform(arr []int) []int {
	rank := make([]int, len(arr))
	copy(rank, arr)

	sort.Ints(rank)
	minus := 0
	memo := make(map[int]int, len(arr))
	for i, n := range rank {
		if _, ok := memo[n]; !ok {
			memo[n] = i + 1 - minus
		} else {
			minus++
		}
	}

	for i := 0; i < len(arr); i++ {
		arr[i] = memo[arr[i]]
	}

	return arr
}

func getNoZeroIntegers(n int) []int {
	first, second := 1, n-1

	for containsZero(first) || containsZero(second) {
		first++
		second--
	}

	return []int{first, second}
}

func containsZero(n int) bool {
	if n == 0 {
		return true
	}

	for n > 0 {
		if n%10 == 0 {
			return true
		}

		n /= 10
	}

	return false
}
