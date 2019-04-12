package main

import "log"

func main() {
	log.Print(arrLen([]int{1, 2, 3, 4, 5, 6, 10}))
}

func arrLen(items []int) int {
	if len(items) == 1 {
		return 1
	}

	return 1 + arrLen(items[1:])
}
