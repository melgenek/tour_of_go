package main

import (
	"container/list"
	"fmt"
)

func main() {
	m := Constructor()

	m.Put(1, 1)
	m.Put(2, 2)
	fmt.Printf("%v\n", m.Get(1) == 1)
	fmt.Printf("%v\n", m.Get(3) == -1)
	m.Put(2, 1)
	fmt.Printf("%v\n", m.Get(2) == 1)
	m.Remove(2)
	fmt.Printf("%v\n", m.Get(2) == -1)
}

const loadFactor = 0.75

type MyHashMap struct {
	capacity int
	count    int
	storage  []*list.List
}

type Entry struct {
	key   int
	value int
}

func Constructor() MyHashMap {
	return MyHashMap{
		3,
		0,
		createNewStorage(3),
	}
}

func createNewStorage(capacity int) []*list.List {
	arr := make([]*list.List, capacity)
	for i := range arr {
		arr[i] = list.New()
	}
	return arr
}

func (this *MyHashMap) Put(key int, value int) {
	if float64(this.count) > float64(this.capacity)*loadFactor {
		newStorage := createNewStorage(this.capacity * 2)
		newCapacity := this.capacity * 2
		for _, bucket := range this.storage {
			for el := bucket.Front(); el != nil; el = el.Next() {
				v := el.Value.(*Entry)
				newStorage[v.key%newCapacity].PushBack(v)
			}
		}
		this.storage = newStorage
		this.capacity = newCapacity
	}

	this.count++
	bucket := this.storage[this.bucketId(key)]
	for el := bucket.Front(); el != nil; el = el.Next() {
		if el.Value.(*Entry).key == key {
			el.Value.(*Entry).value = value
			return
		}
	}
	bucket.PushBack(&Entry{key, value})
}

func (this *MyHashMap) Get(key int) int {
	bucket := this.storage[this.bucketId(key)]
	for el := bucket.Front(); el != nil; el = el.Next() {
		if el.Value.(*Entry).key == key {
			return el.Value.(*Entry).value
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	bucket := this.storage[this.bucketId(key)]
	for el := bucket.Front(); el != nil; el = el.Next() {
		if el.Value.(*Entry).key == key {
			bucket.Remove(el)
			return
		}
	}
}

func (this *MyHashMap) bucketId(key int) int {
	return key % this.capacity
}
