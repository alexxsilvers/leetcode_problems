package main

import "log"

// Design and implement a data structure for Least Recently Used (LRU)
// cache. It should support the following operations: get and put.

// get(key) - Get the value (will always be positive) of the key if
// the key exists in the cache, otherwise return -1.
// put(key, value) - Set or insert the value if the key is not already present.
// When the cache reached its capacity, it should invalidate the least
// recently used item before inserting a new item.
// The cache is initialized with a positive capacity.
// Follow up:
// Could you do both operations in O(1) time complexity?
func main() {
	cache := Constructor(2)

	cache.Put(1, 1)
	cache.Put(2, 2)
	log.Println(cache.Get(1)) // returns 1
	cache.Put(3, 3)           // evicts key 2
	log.Println(cache.Get(2)) // returns -1 (not found)
	cache.Put(4, 4)           // evicts key 1
	log.Println(cache.Get(1)) // returns -1 (not found)
	log.Println(cache.Get(3)) // returns 3
	log.Println(cache.Get(4)) // returns 4
}

// LRUCache - is last recently used cache
type LRUCache struct {
	items map[int]*item
	cap   int
	head  *item
	tail  *item
}

type item struct {
	key  int
	val  int
	next *item
	prev *item
}

// Constructor returns new LRUCache
func Constructor(capacity int) LRUCache {
	c := LRUCache{
		items: make(map[int]*item, capacity),
		cap:   capacity,
		head:  &item{},
		tail:  &item{},
	}

	c.head.next = c.tail
	c.tail.prev = c.head

	return c
}

// Get return value if exist or -1
func (c *LRUCache) Get(key int) int {
	item, exist := c.items[key]
	if exist {
		c.remove(item)
		c.add(item)
		return item.val
	}

	return -1
}

// Put set new value
func (c *LRUCache) Put(key int, value int) {
	i, exist := c.items[key]
	if exist {
		c.remove(i)
		c.add(i)
		i.val = value

		return
	}

	if len(c.items) == c.cap {
		i := c.head.next
		c.remove(i)
		delete(c.items, i.key)
	}

	item := &item{
		key: key,
		val: value,
	}
	c.add(item)
	c.items[key] = item
}

func (c *LRUCache) add(i *item) {
	prev := c.tail.prev
	prev.next = i

	c.tail.prev = i
	i.prev = prev
	i.next = c.tail
}

func (c *LRUCache) remove(i *item) {
	prev := i.prev
	next := i.next

	prev.next = next
	next.prev = prev
}
