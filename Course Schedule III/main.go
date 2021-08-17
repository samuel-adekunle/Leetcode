package main

import (
	"fmt"
	"sort"
)

func main() {
	courses := [][]int{
		{3, 2},
		{4, 3},
	}
	fmt.Println("Dynamic =", scheduleCourseDynamic(courses))
	fmt.Println("Iterative =", scheduleCourseIterative(courses))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


func scheduleCourseDynamic(courses [][]int) int {
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})
	var helper func(i, time int) int
	cache := map[string]int{}
	helper = func(i, time int) int {
		id := fmt.Sprintf("%v#%v", i, time)
		if val, ok := cache[id]; ok {
			return val
		}
		if i >= len(courses) {
			return 0
		}
		not_taken := helper(i+1, time)
		if time+courses[i][0] <= courses[i][1] {
			taken := 1 + helper(i+1, time+courses[i][0])
			cache[id] = max(taken, not_taken)
		} else {
			cache[id] = not_taken
		}
		return cache[id]
	}
	return helper(0, 0)
}


func scheduleCourseIterative(courses [][]int) int {
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})
	count, time := 0, 0
	for i, course := range courses {
		
		if time+course[0] <= course[1] {
			
			count++
			time += course[0]
		} else {
			
			maxDuration, maxIndex := course[0], i
			for j := 0; j < i; j++ {
				if courses[j][0] > 0 && courses[j][0] > maxDuration {
					maxIndex = j
					maxDuration = courses[j][0]
				}
			}
			time += course[0] - maxDuration
			courses[maxIndex][0] = -1 
			
		}
	}
	return count
}
