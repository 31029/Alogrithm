package hashmap

import (
	utils "github.com/h31029/alogrithm/utils"
)

// leetcode-15：三数之和
func ThreeSum(nums []int) [][]int {
	// x3: 去重问题，如何去重
	// 一个关键点和一个难点， 双指针和去重（剪枝）
	result := [][]int{}
	utils.BubbleSort(nums)

	for i := 0; i < len(nums)-2; i++ {
		// 提前结束
		if nums[i] > 0 {
			return result
		}

		// 错误去重a方法，将会漏掉-1,-1,2 这种情况
		/*
		if (nums[i] == nums[i + 1]) {
			continue;
		}
		*/
		// 正确去重a方法
		if (i > 0 && nums[i] == nums[i - 1]) {
			continue;
		}

		left := i + 1 
		right := len(nums) - 1
		for right > left {
			if nums[i] + nums[right] + nums[left] == 0 {
				result = append(result, []int{nums[i], nums[right], nums[left]})
				// 去重 b 和 c
				for right > left && nums[right-1] == nums[right] {
					right --
				}
				for right > left && nums[left] == nums[left+1] {
					left ++
				}
				right --
				left ++
			} else if nums[i] + nums[right] + nums[left] > 0 {
				right --
			} else {
				left ++
			}
		}
	}
	return result
}