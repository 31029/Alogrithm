package hashmap

func TwoSum(nums []int, target int) []int {
	result := [2]int{}
	for i := 0; i < len(nums)-1; i++ {
		for j := i+1; j < len(nums); j++ {
			if nums[i] + nums[j] == target {
				result[0] = i
				result[1] = j
			}
		}
	}

	return result[:]
}
