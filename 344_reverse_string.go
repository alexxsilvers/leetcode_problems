package main

import "log"

// Write a function that reverses a string. The input string is given as an array of characters char[].
//
// Do not allocate extra space for another array, you must do this by modifying the input array
// in-place with O(1) extra memory.
//
// You may assume all the characters consist of printable ascii characters.
func main() {
	ex1 := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(ex1)
	log.Printf("%v", ex1) // o l l e h
}

func reverseString(s []byte)  {

}