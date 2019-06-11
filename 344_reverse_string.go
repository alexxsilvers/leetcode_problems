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
	log.Printf("%s", ex1) // olleh

	ex2 := []byte{'H','a','n','n','a','h'}
	reverseString(ex2)
	log.Printf("%s", ex2) // hannaH

	ex3 := []byte{'A',' ','m','a','n',',',' ','a',' ','p','l','a','n',',',' ','a',' ','c','a','n','a','l',':',' ','P','a','n','a','m','a'}
	reverseString(ex3)
	log.Printf("%s", ex3) // amanaP :lanac a,nalp a ,nam A
}

func reverseString(s []byte)  {
	if len(s) <= 1 {
		return
	}

	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s) - 1 - i] = s[len(s) - 1 - i], s[i]
	}
}
