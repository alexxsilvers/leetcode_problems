package main

import "log"

// You have a RecentCounter class which counts the number of recent requests within a certain time frame.
// Implement the RecentCounter class:
//
// RecentCounter() Initializes the counter with zero recent requests.
// int ping(int t) Adds a new request at time t, where t represents some time in milliseconds,
// and returns the number of requests that has happened in the past 3000 milliseconds (including the new
// request). Specifically, return the number of requests that have happened in the inclusive range [t - 3000, t].
// It is guaranteed that every call to ping uses a strictly larger value of t than the previous call.
func main() {
	counter := Constructor()
	log.Println(counter.Ping(1))
	log.Println(counter.Ping(100))
	log.Println(counter.Ping(3001))
	log.Println(counter.Ping(3002))
}

type RecentCounter struct {
	head *node
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	if this.head == nil {
		this.head = &node{
			time: t,
		}
	} else {
		temp := this.head
		this.head = &node{
			time: t,
			next: temp,
		}
	}

	latest, cnt, root := t-3000, 0, this.head
	for root != nil {
		if root.time < latest {
			break
		}

		cnt++
		root = root.next
	}

	return cnt
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */

type node struct {
	time int
	next *node
}
