package main

import "log"

// Given a year Y and a month M, return how many days there are in that month.
func main() {
	log.Println(numberOfDays(1992, 7))
	log.Println(numberOfDays(2000, 2))
	log.Println(numberOfDays(1900, 2))
}

func numberOfDays(Y int, M int) int {
	if M == 2 {
		if Y%400 == 0 || (Y%4 == 0 && Y%100 != 0) {
			return 29
		} else {
			return 28
		}
	} else if M == 1 || M == 3 || M == 5 || M == 7 || M == 8 || M == 10 || M == 12 {
		return 31
	} else {
		return 30
	}
}
