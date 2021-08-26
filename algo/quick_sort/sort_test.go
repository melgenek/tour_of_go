package main

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"testing"
)

var tests = []struct {
	arr            []int
	expectedResult []int
}{
	{[]int{1, 2, 3}, []int{1, 2, 3}},
	{[]int{3, 2, 1}, []int{1, 2, 3}},
	{[]int{4, 3, 2, 1, 5}, []int{1, 2, 3, 4, 5}},
	{[]int{4, 3, 2, 3, 1, 5}, []int{1, 2, 3, 3, 4, 5}},
	{[]int{81, 887, 847, 59, 81}, []int{59, 81, 81, 847, 887}},
}

func TestSortFirst(t *testing.T) {
	for idx, test := range tests {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			QuickSortFirst(test.arr)
			if !reflect.DeepEqual(test.arr, test.expectedResult) {
				t.Fatalf("expected: %v, got: %v", test.expectedResult, test.arr)
			}
		})
	}
}

func TestSortLast(t *testing.T) {
	for idx, test := range tests {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			QuickSortLast(test.arr)
			if !reflect.DeepEqual(test.arr, test.expectedResult) {
				t.Fatalf("expected: %v, got: %v", test.expectedResult, test.arr)
			}
		})
	}
}

func TestSortMedian(t *testing.T) {
	for idx, test := range tests {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			QuickSortMedian(test.arr)
			if !reflect.DeepEqual(test.arr, test.expectedResult) {
				t.Fatalf("expected: %v, got: %v", test.expectedResult, test.arr)
			}
		})
	}
}

func TestRandom(t *testing.T) {
	for i := 1; i <= 10; i++ {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			a, b, c := RandomArray(100)
			QuickSortMedian(a)
			sort.Slice(b, func(i, j int) bool {
				return b[i] < b[j]
			})

			if !reflect.DeepEqual(a, b) {
				t.Fatalf("Raw: %v, expected: %v, got: %v", c, b, a)
			}
		})
	}
}
