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

	fmt.Printf("%v\n", cache.Get(2) == 124)
	fmt.Printf("%v\n", cache.Get(1) == 1233)

}

type LRUCache struct {
	capacity int
	q        *list.List
	m        map[int]*list.Element
}

type KV struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity,
		list.New(),
		make(map[int]*list.Element),
	}
}

func (this *LRUCache) Get(key int) int {
	v := this.m[key]
	if v == nil {
		return -1
	} else {
		this.m[key] = this.updateKey(v.Value.(KV))
		return v.Value.(KV).value
	}
}

func (this *LRUCache) Put(key int, value int) {
	this.m[key] = this.updateKey(KV{key, value})
	if this.q.Len() > this.capacity {
		last := this.q.Back()
		this.q.Remove(last)
		delete(this.m, last.Value.(KV).key)
	}
}

func (this *LRUCache) updateKey(kv KV) *list.Element {
	v := this.m[kv.key]
	if v != nil {
		this.q.Remove(v)
	}
	this.q.PushFront(kv)
	return this.q.Front()
}
