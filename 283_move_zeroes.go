package main

import "log"

// Given an array nums, write a function to move all 0's to the end of it
// while maintaining the relative order of the non-zero elements.
func main() {
	ex1 := []int{0,1,0,3,12}
 	moveZeroes(ex1)
	log.Println(ex1) // [1,3,12,0,0]

	ex2 := []int{0,0,0,1,2,1,0,10}
	moveZeroes(ex2)
	log.Println(ex2)// [1,2,1,10,0,0,0,0]

	ex3 := []int{0,0,5,0,0,22,11,0,0,0,0}
	moveZeroes(ex3)
	log.Println(ex3) // [5,22,11,0,0,0,0,0,0,0,0]

	ex4 := []int{0,0,0,0,1}
	moveZeroes(ex4)
	log.Println(ex4) // [1,0,0,0,0]
}

func moveZeroes(nums []int)  {
	for lastNonZero, cur := 0, 0; cur < len(nums); cur++ {
		if nums[cur] != 0 {
			nums[cur], nums[lastNonZero] = nums[lastNonZero], nums[cur]
			lastNonZero++
		}
	}
}