package main

import (
	"fmt"
)

func main() {
	nums := [4]int{2, 7, 11, 15}
	target := 18
	fmt.Println(twoSum1(nums[:], target))
}

// 暴力循环：时间复杂度：O(n2)  空间复杂度：O(1)
func twoSum(nums []int, target int) [2]int {
	var rets = [2]int{-1, -1}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				rets[0] = i
				rets[1] = j
				return rets
			}
		}
	}
	return rets
}

// hashmap 时间：O(n) 空间：O(n)
func twoSum1(nums []int, target int) [2]int {
	numMap := make(map[int]int)
	ret := [2]int{-1, -1}
	for k, v := range nums {
		numMap[v] = k
	}

	for i := 0; i < len(nums); i++ {
		sub := target - nums[i]
		if j, ok := numMap[sub]; ok {
			ret[0] = i
			ret[1] = j
			break
		}
	}

	return ret
}
