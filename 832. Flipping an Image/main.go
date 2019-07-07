package main

import "log"

// Given a binary matrix A, we want to flip the image horizontally, then invert it, and return the resulting image.
//
// To flip an image horizontally means that each row of the image is reversed.
// For example, flipping [1, 1, 0] horizontally results in [0, 1, 1].
//
// To invert an image means that each 0 is replaced by 1, and each 1 is replaced by 0.
// For example, inverting [0, 1, 1] results in [1, 0, 0].
func main() {
	e := make([][]int, 0, 3)
	e = append(e, []int{1, 0, 1}, []int{1, 1, 0, 0}, []int{1, 0, 0, 1, 0})
	ex1 := flipAndInvertImage(e)
	log.Printf("%#v", ex1)
}

func flipAndInvertImage(A [][]int) [][]int {
	for _, row := range A {
		// flip
		i := 0
		j := len(row) - 1
		for i < j {
			row[i], row[j] = row[j], row[i]
			i++
			j--
		}

		// invert
		for i := range row {
			row[i] = (row[i] - 1) * -1
		}
	}

	return A
}
