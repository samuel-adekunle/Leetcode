package main

import (
	"fmt"
	"sort"
)

func main() {
	candidates, target := []int{2,7,6,3,5,1}, 9
	validCombinations := combinationSum(candidates, target)
	fmt.Println(validCombinations)
}



func combinationSum(candidates []int, target int) [][]int {
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})

	validCombinations := [][]int{}
	keys := map[string]bool{}
	var appendCombination func([]int)
	appendCombination = func (combination []int)  {
		sort.Slice(combination, func(i, j int) bool {
			return combination[i] < combination[j]
		})
		key := ""
		for _,v := range combination {
			key += fmt.Sprintf("%v", v)
		}
		_, ok := keys[key]
		if !ok {
			fmt.Println(key)
			keys[key] = true
			validCombinations = append(validCombinations, combination)
		}
	}

	cache := map[[2]int][]int{} 
	var helper func([]int, int)
	helper = func(curr []int, sum int) {
		if sum == target {
			appendCombination(curr)	
		} else {
			for _,candidate := range candidates {
				cache[[2]int{curr[0], sum}] =  curr
				v, ok := cache[[2]int{candidate, target - sum}]
				if ok {
					appendCombination(append(curr, v...))
				} else if candidate >= curr[0] && (sum + candidate <= target) {
					helper(append(curr, candidate), sum + candidate)
				}
			}
		}
	}
	for _,candidate := range candidates {
		helper([]int{candidate}, candidate)
	}
	return validCombinations
}