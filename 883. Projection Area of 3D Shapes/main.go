package main

import "log"

// On a N * N grid, we place some 1 * 1 * 1 cubes that are axis-aligned with the x, y, and z axes.
//
// Each value v = grid[i][j] represents a tower of v cubes placed on top of grid cell (i, j).
//
// Now we view the projection of these cubes onto the xy, yz, and zx planes.
//
// A projection is like a shadow, that maps our 3 dimensional figure to a 2 dimensional plane.
//
// Here, we are viewing the "shadow" when looking at the cubes from the top, the front, and the side.
//
// Return the total area of all three projections.
func main() {
	ex1 := [][]int{
		{5},
	}
	log.Println(projectionArea(ex1)) // 1(ground)+5(sum of top in the each row)+5(sum of top in the each col)=11

	ex2 := [][]int{
		{1, 2},
		{3, 4},
	}
	log.Println(projectionArea(ex2)) // 4(ground)+6(sum of top in the each row)+7(sum of top in the each col)=17

	ex3 := [][]int{
		{1, 1, 1},
		{1, 0, 0},
		{1, 0, 1},
	}
	log.Println(projectionArea(ex3)) // 6(ground)+3(sum of top in the each row)+3(sum of the top in the each col)=12

	ex4 := [][]int{
		{1, 0},
		{0, 2},
	}
	log.Println(projectionArea(ex4)) // 2(ground)+3(sum of top in the each row)+3(sum of the top in the each col)=8
}

func projectionArea(grid [][]int) int {
	total := 0
	for i := 0; i < len(grid); i++ {
		bestRow, bestCol := 0, 0
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] > 0 {
				total++
			}
			if grid[i][j] > bestRow {
				bestRow = grid[i][j]
			}
			if grid[j][i] > bestCol {
				bestCol = grid[j][i]
			}
		}
		total += bestRow + bestCol
	}

	return total
}
