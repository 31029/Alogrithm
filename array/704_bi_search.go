package array

func Search(nums []int, target int) (result int) {
	heigh := len(nums) - 1
	low := 0

	for low <= heigh {
		middle := (heigh + low) / 2 
		if target == nums[middle] {
			return middle
		} else if target > nums[middle] {
			low = middle + 1
		} else {
			heigh = middle - 1
		}
	}
	return -1
}
