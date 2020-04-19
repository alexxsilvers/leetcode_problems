package merge_sort

func MergeSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	memo := make([]int, len(nums))
	mergeSort(nums, memo, 0, len(nums)-1)
}

func mergeSort(nums []int, memo []int, low int, high int) {
	if high <= low {
		return
	}

	mid := low + (high-low)/2

	mergeSort(nums, memo, low, mid)
	mergeSort(nums, memo, mid+1, high)

	merge(nums, memo, low, mid, high)
}

func merge(nums []int, memo []int, low, mid, high int) {
	i := low
	j := mid + 1

	for k := low; k <= high; k++ {
		memo[k] = nums[k]
	}

	for k := low; k <= high; k++ {
		if i > mid {
			nums[k] = memo[j]
			j++
		} else if j > high {
			nums[k] = memo[i]
			i++
		} else if memo[j] < memo[i] {
			nums[k] = memo[j]
			j++
		} else {
			nums[k] = memo[i]
			i++
		}
	}
}
