package main

func trapBruteForce(height []int) int {
	if len(height) <= 1 {
		return 0
	}
	ans := 0
	for i := 0; i < len(height); i++ {
		maxLeft := 0
		maxRight := 0
		for j := i; j < len(height); j++ {
			maxRight = max(maxRight, height[j])
		}

		for j := i; j >= 0; j-- {
			maxLeft = max(maxLeft, height[j])
		}

		ans += min(maxLeft, maxRight) - height[i]
	}

	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
