package main

import "log"

func main() {
	log.Println(judgeCircle("UD"))   // true
	log.Println(judgeCircle("LL"))   // false
	log.Println(judgeCircle("URDL")) // true
}

func judgeCircle(moves string) bool {
	x, y := 0, 0
	for _, command := range []byte(moves) {
		if command == 'U' {
			y++
		}
		if command == 'D' {
			y--
		}
		if command == 'R' {
			x--
		}
		if command == 'L' {
			x++
		}
	}

	return x == 0 && y == 0
}
