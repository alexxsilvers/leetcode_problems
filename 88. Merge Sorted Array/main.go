package main

import "log"

// Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.
// The number of elements initialized in nums1 and nums2 are m and n respectively.
// You may assume that nums1 has enough space (size that is greater or equal to m + n) to hold additional elements from nums2.
func main() {
	arr1 := []int{1, 2, 3, 0, 0, 0}
	arr2 := []int{2, 5, 6}
	merge1(arr1, 3, arr2, 3)
	log.Println(arr1) // [1,2,2,3,5,6]

	arr1 = []int{2, 0}
	arr2 = []int{1}
	merge1(arr1, 1, arr2, 1)
	log.Println(arr1) // [1,2]
}

// from end
func merge1(nums1 []int, m int, nums2 []int, n int) {
	for m > 0 && n > 0 {
		if nums1[m-1] >= nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}

	for n > 0 {
		nums1[n-1] = nums2[n-1]
		n--
	}
}

// from start
func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	save := make([]int, m)
	for i := 0; i < m; i++ {
		save[i] = nums1[i]
	}

	i, j, k := 0, 0, 0
	for j < m && k < n {
		if save[j] < nums2[k] {
			nums1[i] = save[j]
			j++
		} else {
			nums1[i] = nums2[k]
			k++
		}

		i++
	}

	for j < m {
		nums1[i] = save[j]
		j++
		i++
	}

	for k < n {
		nums1[i] = nums2[k]
		k++
		i++
	}
}
