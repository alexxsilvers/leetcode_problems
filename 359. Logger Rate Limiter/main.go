package main

import (
	"log"
	"sync"
)

// Design a logger system that receive stream of messages along with its timestamps,
// each message should be printed if and only if it is not printed in the last 10 seconds.
// Given a message and a timestamp (in seconds granularity), return true if the message
// should be printed in the given timestamp, otherwise returns false.
// It is possible that several messages arrive roughly at the same time.
func main() {
	logger := Constructor()
	log.Println(logger.ShouldPrintMessage(1, "foo")) // true
	log.Println(logger.ShouldPrintMessage(2, "bar")) // true
	log.Println(logger.ShouldPrintMessage(3, "foo")) // false
	log.Println(logger.ShouldPrintMessage(8, "bar")) // false
	log.Println(logger.ShouldPrintMessage(10, "foo")) // false
	log.Println(logger.ShouldPrintMessage(11, "foo")) // true
}

type Logger struct {
	data map[string]int
	m sync.RWMutex
}


/** Initialize your data structure here. */
func Constructor() Logger {
	return Logger{
		data: make(map[string]int),
	}
}


/** Returns true if the message should be printed in the given timestamp, otherwise returns false.
  If this method returns false, the message will not be printed.
  The timestamp is in seconds granularity. */
func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	this.m.RLock()
	ts, exist := this.data[message]
	this.m.RUnlock()
	if !exist {
		this.m.Lock()
		this.data[message] = timestamp
		this.m.Unlock()
		return true
	}

	if timestamp - ts >= 10 {
		this.m.Lock()
		this.data[message] = timestamp
		this.m.Unlock()
		return true
	}

	return false
}


/**
 * Your Logger object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.ShouldPrintMessage(timestamp,message);
 */