package binary_search

func BinarySearch(list []int, item int) int {
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := list[mid]
		if guess == item {
			return mid
		}

		if guess < item {
			low = mid
		} else {
			high = mid + 1
		}
	}

	return -1
}
