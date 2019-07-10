package main

import "log"

// Given an array nums of n integers where n > 1,
// return an array output such that output[i] is equal to
// the product of all the elements of nums except nums[i].
func main() {
	log.Println(productExceptSelf([]int{1, 2, 3, 4})) // [24,12,8,6]
}

func productExceptSelf(nums []int) []int {
	l := len(nums)
	if l == 0 {
		return nums
	}

	out := make([]int, l)
	out[0] = 1
	for i := 1; i < l; i++ {
		out[i] = out[i-1] * nums[i-1]
	}

	tail := 1
	for i := l - 2; i >= 0; i-- {
		tail *= nums[i+1]
		out[i] *= tail
	}
	return out
}
