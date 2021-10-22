package main

import (
	"container/list"
	"fmt"
)

func main() {
	cache := Constructor(2)

	cache.Put(2, 124)
	fmt.Printf("%v\n", cache.Get(2) == 124)

	cache.Put(1, 123)
	fmt.Printf("%v\n", cache.Get(1) == 123)

	cache.Put(1, 1233)
	fmt.Printf("%v\n", cache.Get(1) == 1233)

	fmt.Printf("%v\n", cache.Get(1) == 1233)
	fmt.Printf("%v\n", cache.Get(2) == 124)

	cache.Put(3, 125)
	fmt.Printf("%v\n", cache.Get(3) == 125)

	fmt.Printf("%v\n", cache.Get(1) == 1233)
	fmt.Printf("%v\n", cache.Get(2) == -1)

}

type LFUCache struct {
	capacity int
	time     uint64
	m        map[int]*list.Element
	l        *list.List
}

type Entry struct {
	key     int
	value   int
	useTime uint64
	seen    int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity,
		0,
		make(map[int]*list.Element),
		list.New(),
	}
}

func (this *LFUCache) Get(key int) int {
	el := this.m[key]

	if el != nil {
		this.updateCounters(el)
		return el.Value.(*Entry).value
	} else {
		return -1
	}
}

func (this *LFUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}
	existingEl := this.m[key]
	if existingEl == nil {
		this.tryEvict()
		this.time++
		newEntry := &Entry{key, value, this.time, 1}
		this.l.PushFront(newEntry)
		this.m[key] = this.l.Front()

	} else {
		this.updateCounters(existingEl)
		existingEl.Value.(*Entry).value = value
	}
}

func (this *LFUCache) updateCounters(el *list.Element) {
	entry := el.Value.(*Entry)
	entry.seen++
	this.time++
	entry.useTime = this.time
	if next := el.Next(); next != nil && next.Value.(*Entry).seen < entry.seen {
		this.l.MoveAfter(el, next)
	}
}

func (this *LFUCache) tryEvict() {
	if this.l.Len() == this.capacity {
		leastFrequentlySeenEl := this.l.Front()
		for e := leastFrequentlySeenEl.Next(); e != nil; e = e.Next() {
			thisValue := leastFrequentlySeenEl.Value.(*Entry)
			nextValue := e.Value.(*Entry)
			if thisValue.seen == nextValue.seen && thisValue.useTime > nextValue.useTime {
				leastFrequentlySeenEl = e
			} else {
				break
			}
		}
		this.l.Remove(leastFrequentlySeenEl)
		delete(this.m, leastFrequentlySeenEl.Value.(*Entry).key)
	}
}
