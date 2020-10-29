package main

import "log"

// Design a data structure that accepts integers of a stream, and checks if it has a pair
// of integers that sum up to a particular value.
//
// Implement a TwoSum class:
//
// TwoSum() Initializes the TwoSum object, with an empty array initially.
// void add(int number) Adds number to the data structure.
// boolean find(int value) Returns true if there exists any pair of numbers whose
// sum is equal to value, otherwise, it returns false.
func main() {
	obj := Constructor()
	obj.Add(1)
	obj.Add(3)
	obj.Add(5)
	log.Println(obj.Find(4)) // true
	log.Println(obj.Find(7)) // false

	obj1 := Constructor()
	obj1.Add(0)
	log.Println(obj1.Find(0)) // false

	obj2 := Constructor()
	obj2.Add(0)
	obj2.Add(0)
	log.Println(obj2.Find(0)) // true

	obj3 := Constructor()
	obj3.Add(0)
	obj3.Add(-1)
	obj3.Add(1)
	log.Println(obj3.Find(0)) // true
}

type TwoSum struct {
	nums map[int]int
}

/** Initialize your data structure here. */
func Constructor() TwoSum {
	return TwoSum{
		nums: make(map[int]int),
	}
}

/** Add the number to an internal data structure.. */
func (this *TwoSum) Add(number int) {
	this.nums[number]++
}

/** Find if there exists any pair of numbers which sum is equal to the value. */
func (this *TwoSum) Find(value int) bool {
	for num, cnt := range this.nums {
		find := value - num

		if find == num {
			if cnt > 1 {
				return true
			}

			continue
		}

		if _, ok := this.nums[find]; ok {
			return true
		}
	}
	return false
}

/**
 * Your TwoSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * param_2 := obj.Find(value);
 */
