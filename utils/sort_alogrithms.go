package utils

func BubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j+1] < nums[j] {
				Swap(nums, j+1, j)
			}
		}
	}
}
