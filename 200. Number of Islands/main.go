package main

import "log"

// Given a 2d grid map of '1's (land) and '0's (water), count the number of islands.
// An island is surrounded by water and is formed by connecting adjacent lands
// horizontally or vertically. You may assume all four edges of the grid are all
// surrounded by water.

// INTUITION
// Traverse by grid, if grid[r][c] = 1 - update counter and set grid[r][c]
// and vertical and horizontal around to 0 by dfs
func main() {
	ex1 := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	log.Println(numIslands(ex1)) // 1

	ex3 := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	log.Println(numIslands(ex3)) // 3
}

// Runtime O(m * n) m = len(grid) and n = len(grid[0])
// Space O(1), because we mark traversed cells in grid
func numIslands(grid [][]byte) int {
	m := len(grid)
	if m == 0 {
		return 0
	}

	n := len(grid[0])
	if n == 0 {
		return 0
	}

	cnt := 0
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if grid[r][c] == '0' {
				continue
			}
			cnt++
			dfs(grid, r, c, m, n)
		}
	}

	return cnt
}

func dfs(grid [][]byte, r, c, m, n int) {
	if !sanityCheck(r, c, m, n) {
		return
	}
	if grid[r][c] == '0' {
		return
	}
	grid[r][c] = '0'
	dfs(grid, r-1, c, m, n)
	dfs(grid, r+1, c, m, n)
	dfs(grid, r, c-1, m, n)
	dfs(grid, r, c+1, m, n)
}

func sanityCheck(r, c, m, n int) bool {
	if r < 0 || r > m-1 || c < 0 || c > n-1 {
		return false
	}
	return true
}
