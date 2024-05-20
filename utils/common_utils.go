package utils

func Swap(nums []int, left int, right int)  {
	temp := nums[left]
	nums[left] = nums[right]
	nums[right] = temp
}