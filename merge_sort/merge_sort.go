package merge_sort

func MergeSortStupid(nums []int) {
	if len(nums) <= 1 {
		return
	}

	mergeSortStupid(nums, 0, len(nums)-1)
}

func mergeSortStupid(nums []int, low int, high int) {
	if high <= low {
		return
	}

	mid := low + (high-low)/2

	mergeSortStupid(nums, low, mid)
	mergeSortStupid(nums, mid+1, high)

	mergeStupid(nums, low, mid, high)
}

func mergeStupid(nums []int, low, mid, high int) {
	left := make([]int, 0, mid-low+1)
	right := make([]int, 0, high-mid)

	for i := low; i <= mid; i++ {
		left = append(left, nums[i])
	}
	for k := mid + 1; k <= high; k++ {
		right = append(right, nums[k])
	}

	i, j, k := 0, 0, low
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			nums[k] = left[i]
			i++
			k++
		} else {
			nums[k] = right[j]
			j++
			k++
		}
	}

	for i < len(left) {
		nums[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		nums[k] = right[j]
		j++
		k++
	}
}

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
