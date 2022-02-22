package main

import "fmt"

func main() {
	timeMap := Constructor()
	timeMap.Set("1", "1", 1)
	timeMap.Set("1", "2", 5)

	fmt.Printf("%v\n", timeMap.Get("1", 0) == "none")
	fmt.Printf("%v\n", timeMap.Get("1", 1) == "1")
	fmt.Printf("%v\n", timeMap.Get("1", 1) == "1")
	fmt.Printf("%v\n", timeMap.Get("1", 2) == "1")
	fmt.Printf("%v\n", timeMap.Get("1", 5) == "2")
	fmt.Printf("%v\n", timeMap.Get("1", 6) == "2")
}

type Element struct {
	value     string
	timestamp int
}

type TimeMap struct {
	m map[string][]Element
}

func Constructor() TimeMap {
	return TimeMap{
		make(map[string][]Element),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.m[key] = append(this.m[key], Element{value: value, timestamp: timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	result := get(this.m[key], timestamp, nil)
	if result == nil {
		return ""
	} else {
		return result.value
	}
}

func get(arr []Element, timestamp int, previousTimestamp *Element) *Element {
	if len(arr) == 0 {
		return previousTimestamp
	} else {
		middleIdx := len(arr) / 2
		middle := arr[middleIdx]

		if middle.timestamp == timestamp {
			return &middle
		} else if middle.timestamp > timestamp {
			return get(arr[:middleIdx], timestamp, previousTimestamp)
		} else {
			return get(arr[middleIdx+1:], timestamp, &middle)
		}
	}
}
