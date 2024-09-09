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


// 使用map方式解题，降低时间复杂度
func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for index, val := range nums {
        if preIndex, ok := m[target-val]; ok {
            return []int{preIndex, index}
        } else {
            m[val] = index
        }
    }
    return []int{}
}