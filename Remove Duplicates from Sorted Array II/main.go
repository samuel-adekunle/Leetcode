package main

func removeDuplicates(nums []int) int {
	i := 0
	for _, v := range nums {
		if i < 2 || v > nums[i-2] {
			nums[i] = v
			i++
		}
	}
	return i
}
