package main

import "log"

func main() {
	log.Println(increment(101))
}

func increment(y uint) uint {
	if y == 0 {
		return 1
	} else if y % 2 == 1 {
		return 2 * increment(y / 2)
	} else {
		return y + 1
	}
}