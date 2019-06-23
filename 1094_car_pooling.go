package main

import "log"

// You are driving a vehicle that has capacity empty seats initially available for passengers.
// The vehicle only drives east (ie. it cannot turn around and drive west.)
//
// Given a list of trips, trip[i] = [num_passengers, start_location, end_location] contains
// information about the i-th trip: the number of passengers that must be picked up, and the
// locations to pick them up and drop them off.  The locations are given as the number of
// kilometers due east from your vehicle's initial location.
//
// Return true if and only if it is possible to pick up and drop off all passengers for all the given trips.
func main() {
	log.Println(carPooling([][]int{{2, 1, 5}, {3, 3, 7}}, 4))             // false
	log.Println(carPooling([][]int{{2, 1, 5}, {3, 3, 7}}, 5))             // true
	log.Println(carPooling([][]int{{2, 1, 5}, {3, 5, 7}}, 3))             // true
	log.Println(carPooling([][]int{{3, 2, 7}, {3, 7, 9}, {8, 3, 9}}, 11)) // true
	log.Println(carPooling([][]int{{1, 1, 3}, {2, 2, 4}}, 3))             // true
}

// Runtime O(m * n), space O(m * n)
func carPooling(trips [][]int, capacity int) bool {
	kilometersToPassengers := make(map[int]int)
	for _, trip := range trips {
		for i := trip[1]; i < trip[2]; i++ {
			if _, ok := kilometersToPassengers[i]; !ok {
				kilometersToPassengers[i] = 0
			}
			kilometersToPassengers[i] += trip[0]
		}
	}

	for _, passengerCount := range kilometersToPassengers {
		if passengerCount > capacity {
			return false
		}
	}

	return true
}
