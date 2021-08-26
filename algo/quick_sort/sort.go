package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

func main() {
	a, b, c := RandomArray(100000)
	aCount := QuickSortFirst(a)
	bCount := QuickSortLast(b)
	cCount := QuickSortMedian(c)
	fmt.Printf("First: %d, Last: %d, Median: %d\n", aCount, bCount, cCount)
	fmt.Printf("%v %v\n", reflect.DeepEqual(a, b), reflect.DeepEqual(b, c))

}

func RandomArray(n int) ([]int, []int, []int) {
	a := make([]int, n)
	b := make([]int, n)
	c := make([]int, n)

	for i := 0; i < n; i++ {
		el := rand.Intn(1000)
		a[i] = el
		b[i] = el
		c[i] = el
	}

	return a, b, c
}

func QuickSortFirst(arr []int) int {
	return quickSort(arr, func(arr []int) int {
		return 0
	})
}

func QuickSortLast(arr []int) int {
	return quickSort(arr, func(arr []int) int {
		return len(arr) - 1
	})
}

func QuickSortMedian(arr []int) int {
	return quickSort(arr, func(arr []int) int {
		size := len(arr)
		first := arr[0]
		last := arr[size-1]
		middleIdx := size / 2
		if size%2 == 1 {
			middleIdx = size/2 + 1
		}
		middle := arr[middleIdx]
		if middle < last && middle > first || middle < first && middle > last {
			return middleIdx
		} else if first > middle && first < last || first > last && first < middle {
			return 0
		} else {
			return size - 1
		}
	})
}

func quickSort(arr []int, choosePivot func([]int) int) int {
	size := len(arr)
	if size < 2 {
		return 0
	} else {
		pivot := choosePivot(arr)
		arr[0], arr[pivot] = arr[pivot], arr[0]
		el := arr[0]

		i := 1

		for j := i; j < size; j++ {
			if arr[j] < el {
				arr[i], arr[j] = arr[j], arr[i]
				i++
			}
		}

		arr[0], arr[i-1] = arr[i-1], arr[0]

		return quickSort(arr[:i-1], choosePivot) + quickSort(arr[i:], choosePivot) + size - 1
	}

}
