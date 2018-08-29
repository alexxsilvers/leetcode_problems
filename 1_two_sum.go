package main

import "fmt"

func twoSum(nums []int, target int) []int {
	n := make(map[int]int)
	for index, num := range nums {
		n[num] = index
	}
	for index, num := range nums {
		complement := target - num
		secondIndex, exist := n[complement]
		if exist && index != secondIndex {
			return []int{index, secondIndex}
		}
	}
	return nil
}

func main() {
	fmt.Print(twoSum([]int{2, 7, 11, 15}, 18))
}