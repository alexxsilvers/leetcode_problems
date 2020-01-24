package main

func trapDP(height []int) int {
	ans := 0
	left := make([]int, len(height))
	right := make([]int, len(height))
	left[0], right[len(height)-1] = height[0], height[len(height)-1]

	// dp left
	for i := 1; i < len(height); i++ {
		left[i] = max(left[i-1], height[i])
	}

	// dp right
	for i := len(height) - 2; i >= 0; i-- {
		right[i] = max(right[i+1], height[i])
	}

	for i := 0; i < len(height); i++ {
		ans += min(left[i], right[i]) - height[i]
	}

	return ans
}
