package main

import "log"

// Design a parking system for a parking lot. The parking lot has three kinds
// of parking spaces: big, medium, and small, with a fixed number of slots for each size.
//Implement the ParkingSystem class:
//ParkingSystem(int big, int medium, int small) Initializes object of the ParkingSystem class.
// The number of slots for each parking space are given as part of the constructor.
// bool addCar(int carType) Checks whether there is a parking space of carType for the
// car that wants to get into the parking lot. carType can be of three kinds: big, medium,
// or small, which are represented by 1, 2, and 3 respectively. A car can only park in
// a parking space of its carType. If there is no space available, return false, else park
// the car in that size space and return true.
func main() {
	p := Constructor(1, 1, 0)
	log.Println(p.AddCar(1)) // true
	log.Println(p.AddCar(2)) // true
	log.Println(p.AddCar(3)) // false
	log.Println(p.AddCar(1)) // false
}

type ParkingSystem struct {
	parking []int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	p := ParkingSystem{
		parking: make([]int, 3),
	}

	p.parking[0] = big
	p.parking[1] = medium
	p.parking[2] = small

	return p
}

func (this *ParkingSystem) AddCar(carType int) bool {
	carType -= 1

	capacity := this.parking[carType]
	if capacity-1 < 0 {
		return false
	}

	this.parking[carType] -= 1
	return true
}
