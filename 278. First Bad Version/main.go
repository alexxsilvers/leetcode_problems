package main

import "log"

/*
You are a product manager and currently leading a team to develop a new product. Unfortunately,
the latest version of your product fails the quality check. Since each version is developed based on
the previous version, all the versions after a bad version are also bad.

Suppose you have n versions [1, 2, ..., n] and you want to find out the first bad one, which causes all the
following ones to be bad.

You are given an API bool isBadVersion(version) which returns whether version is bad. Implement a function
to find the first bad version. You should minimize the number of calls to the API.
*/

func main() {
	//log.Println(firstBadVersion(5)) // 4
	log.Println(firstBadVersion(1)) // 1
}

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	bad, l, r := -1, 1, n

	for l <= r {
		mid := l + (r-l)/2
		isBad := isBadVersion(mid)

		if isBad {
			bad = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return bad
}

func isBadVersion(version int) bool {
	if version == 1 {
		return true
	}
	return false
}
