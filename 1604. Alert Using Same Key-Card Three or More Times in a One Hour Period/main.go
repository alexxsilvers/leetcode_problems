package main

import (
	"log"
	"sort"
	"strconv"
	"strings"
)

// Leetcode company workers use key-cards to unlock office doors. Each time a
// worker uses their key-card, the security system saves the worker's name and
// the time when it was used. The system emits an alert if any worker uses the
// key-card three or more times in a one-hour period.
// You are given a list of strings keyName and keyTime where [keyName[i], keyTime[i]]
// corresponds to a person's name and the time when their key-card was used in a single day.
// Access times are given in the 24-hour time format "HH:MM", such as "23:51" and "09:49".
// Return a list of unique worker names who received an alert for frequent keycard
// use. Sort the names in ascending order alphabetically.
// Notice that "10:00" - "11:00" is considered to be within a one-hour period,
// while "23:51" - "00:10" is not considered to be within a one-hour period.
func main() {
	log.Println(alertNames(
		[]string{"daniel", "daniel", "daniel", "luis", "luis", "luis", "luis"},
		[]string{"10:03", "10:40", "11:00", "09:00", "11:00", "13:00", "15:00"},
	)) // ["daniel"]

	log.Println(alertNames(
		[]string{"alice", "alice", "alice", "bob", "bob", "bob", "bob"},
		[]string{"12:01", "12:00", "18:00", "21:00", "21:20", "21:30", "23:00"},
	)) // ["bob"]

	log.Println(alertNames(
		[]string{"john", "john", "john"},
		[]string{"23:58", "23:59", "00:01"},
	)) // []

	log.Println(alertNames(
		[]string{"leslie", "leslie", "leslie", "clare", "clare", "clare", "clare"},
		[]string{"13:00", "13:20", "14:00", "18:00", "18:51", "19:30", "19:49"},
	)) // ["clare", "leslie"]

	log.Println(alertNames(
		[]string{"a", "a", "a", "a", "a", "a", "b", "b", "b", "b", "b"},
		[]string{"23:27", "03:14", "12:57", "13:35", "13:18", "21:58", "22:39", "10:49", "19:37", "14:14", "10:41"},
	)) // ["a"]
}

func alertNames(keyName []string, keyTime []string) []string {
	names := make(map[string][]int)
	for _, name := range keyName {
		if _, ok := names[name]; !ok {
			names[name] = []int{}
		}
	}

	for i := 0; i < len(keyTime); i++ {
		times := strings.Split(keyTime[i], ":")
		h, _ := strconv.Atoi(times[0])
		m, _ := strconv.Atoi(times[1])
		names[keyName[i]] = append(names[keyName[i]], h*60+m)
	}

	var alerted []string
	for name, times := range names {
		if len(times) < 3 {
			continue
		}

		sort.Ints(times)
		needToAlert := false
		for i := 0; i+2 < len(times); i++ {
			if times[i+2]-times[i] <= 60 {
				needToAlert = true
				break
			}
		}

		if needToAlert {
			alerted = append(alerted, name)
		}
	}

	sort.Strings(alerted)
	return alerted
}
