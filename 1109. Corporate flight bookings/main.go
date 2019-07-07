package main

import "log"

// There are n flights, and they are labeled from 1 to n.
//
// We have a list of flight bookings.  The i-th booking bookings[i] = [i, j, k]
// means that we booked k seats from flights labeled i to j inclusive.
//
// Return an array answer of length n, representing the number of seats booked
// on each flight in order of their label.
func main() {
	log.Println(corpFlightBookings([][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}, 5)) // [10,55,45,25,25]
}

func corpFlightBookings(bookings [][]int, n int) []int {
	flightsSeats := make([]int, n+1)
	for _, book := range bookings {
		flStart := book[0]
		flEnd := book[1]
		seats := book[2]
		for i := flStart; i <= flEnd; i++ {
			flightsSeats[i] += seats
		}
	}

	return flightsSeats[1:]
}
