package main

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	orig, dupl := 0, 0
	for dupl < len(nums) {
		if nums[orig] != nums[dupl] {
			orig++
			nums[orig] = nums[dupl]
		}
		dupl++
	}
	return orig + 1
}

func onePass(nums []int) int {
	i := 0
	for _, v := range nums {
		if i < 1 || v > nums[i-1] {
			nums[i] = v
			i++
		}
	}
	return i
}
