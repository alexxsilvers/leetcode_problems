package main

import "log"

// Given a non-empty 2D array grid of 0's and 1's, an island is a group of 1's (representing land)
// connected 4-directionally (horizontal or vertical.) You may assume all four edges of the grid are surrounded by water.
//
// Find the maximum area of an island in the given 2D array. (If there is no island, the maximum area is 0.)
// SOLUTION
// Traverse for grid and if find grid[r][c]=1 - start dfs and change visited points to 0
func main() {
	ex1 := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}
	log.Println(maxAreaOfIsland(ex1)) // 6

	ex2 := [][]int{{0, 0, 0, 0, 0, 0, 0, 0}}
	log.Println(maxAreaOfIsland(ex2)) // 0

	ex3 := [][]int{{0, 0, 0, 0, 0, 0, 0, 1}}
	log.Println(maxAreaOfIsland(ex3)) // 1

	ex4 := [][]int{{1, 1, 0, 0, 0, 0, 1, 0}}
	log.Println(maxAreaOfIsland(ex4)) // 2
}

// Runtime O(m * n) m = len(grid) and n = len(grid[0])
// Space O(1)
func maxAreaOfIsland(grid [][]int) int {
	maxArea, m, n := 0, len(grid), len(grid[0])
	if m == 0 || n == 0 {
		return 0
	}

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if grid[r][c] == 0 {
				continue
			}

			localMA := dfs(grid, r, c, m, n)
			if localMA > maxArea {
				maxArea = localMA
			}
		}
	}

	return maxArea
}

func dfs(grid [][]int, r, c, m, n int) int {
	if !sanityCheck(r, c, m, n) {
		return 0
	}
	if grid[r][c] == 0 {
		return 0
	}

	maxArea := 1
	grid[r][c] = 0

	maxArea += dfs(grid, r-1, c, m, n)
	maxArea += dfs(grid, r+1, c, m, n)
	maxArea += dfs(grid, r, c-1, m, n)
	maxArea += dfs(grid, r, c+1, m, n)

	return maxArea
}

func sanityCheck(r, c, m, n int) bool {
	if r < 0 || r > m-1 || c < 0 || c > n-1 {
		return false
	}
	return true
}
