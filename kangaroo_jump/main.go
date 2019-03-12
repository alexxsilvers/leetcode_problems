package main

import "log"

func main() {
	log.Println(meetAtPoint(kangaroo{0, 2}, kangaroo{5, 3}))   //false
	log.Println(meetAtPoint(kangaroo{4, 2}, kangaroo{0, 3}))   //true
	log.Println(meetAtPoint(kangaroo{4, -3}, kangaroo{0, -2})) //true
	log.Println(meetAtPoint(kangaroo{0, 6}, kangaroo{2, 2}))   //false
	log.Println(meetAtPoint(kangaroo{0, 6}, kangaroo{8, 2}))   //false
	log.Println(meetAtPoint(kangaroo{0, 3}, kangaroo{24, 2}))  //true
	log.Println(meetAtPoint(kangaroo{0, 7}, kangaroo{2, 2}))   //false
	log.Println(meetAtPoint(kangaroo{0, 6}, kangaroo{4, 2}))   //true
}

type kangaroo struct {
	x int
	v int
}

func meetAtPoint(first kangaroo, second kangaroo) bool {

}
