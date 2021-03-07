package main

import "fmt"

// 暴力枚举
func twoSum1(nums []int, target int) []int {
	for i, x := range nums {
		for j := i + 1; j < len(nums); j++ {
			if x+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// 哈希表
func twoSum2(nums []int, target int) []int {
	sumMap := make(map[int]int)
	for i, x := range nums {
		if p, ok := sumMap[target-x]; ok {
			return []int{p, i}
		}
		sumMap[x] = i
	}
	return nil
}

func main() {
	fmt.Println(twoSum1([]int{3, 2, 4}, 6))
	fmt.Println(twoSum2([]int{3, 2, 4}, 6))
}
