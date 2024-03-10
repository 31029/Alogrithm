package array

import "github.com/h31029/alogrithm/utils"

func SortedSquares(nums []int) []int {
	for k, v := range nums {
		nums[k] = v * v
	}

	utils.BubbleSort(nums)
	return nums
}