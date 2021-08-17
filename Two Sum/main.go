package main

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSumUseMap(nums []int, target int) []int {
	mem := make(map[int]int)
	for i, v := range nums {
		if ci, ok := mem[target-v]; ok {
			return []int{ci, i}
		}
		mem[v] = i
	}
	return nil
}

func twoSumSorted(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for right > left {
		sum := nums[left] + nums[right]
		if sum == target {
			return []int{left, right}
		}

		if sum < target {
			left++
		} else {
			right--
		}
	}
	return nil
}
