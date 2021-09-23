package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("%v\n", scheduleCourse([][]int{
		{100, 200},
		{200, 1300},
		{1000, 1250},
		{2000, 3200},
	}) == 3)
	fmt.Printf("%v\n", scheduleCourse([][]int{
		{1, 2},
	}) == 1)
	fmt.Printf("%v\n", scheduleCourse([][]int{
		{3, 2},
		{4, 3},
	}) == 0)
	fmt.Printf("%v\n", scheduleCourse([][]int{
		{1, 2},
		{2, 3},
	}) == 2)
	fmt.Printf("%v\n", scheduleCourse([][]int{
		{5, 15}, {3, 19}, {6, 7}, {2, 10}, {5, 16}, {8, 14}, {10, 11}, {2, 19},
	}) == 5)
	fmt.Printf("%v\n", scheduleCourse([][]int{
		{7, 17}, {3, 12}, {10, 20}, {9, 10}, {5, 20}, {10, 19}, {4, 18},
	}) == 4)
}

type Course struct {
	days     int
	deadline int
	cost     int
}

func scheduleCourse(input [][]int) int {
	courses := make([]Course, 0)
	for _, course := range input {
		c := Course{course[0], course[1], course[1] + course[0]}
		if c.cost > 0 {
			courses = append(courses, c)
		}
	}
	sort.Slice(courses, func(i, j int) bool {
		return courses[i].cost < courses[j].cost
	})

	total := 0
	day := 0

	for _, course := range courses {
		if day+course.days <= course.deadline {
			total += 1
			day += course.days
		}
	}

	return total
}
