package dualPointer

func Swap(nums []int, l int, r int) {
	t := nums[l]
	nums[l] = nums[r]
	nums[r] = t
}

func BubbulSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				Swap(nums, j, j+1)
			}
		}
	}
}

// 判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0
// 请你返回所有和为 0 且不重复的三元组。
// [-4,-1,-1,-1,0,1,2,2]
func TreeSum(nums []int) [][]int {
	// 思路：定一、动二，三指针问题。
	// 技巧：剪枝
	BubbulSort(nums)
	result := [][]int{}

	if nums[0] > 0 || nums[len(nums)-1] < 0 {
		return result
	}

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			return result
		}

		if i != 0 && nums[i] == nums[i-1] {
			continue
		}

		l := i + 1
		r := len(nums) - 1
		for l < r {
			if nums[i]+nums[l]+nums[r] == 0 {
				result = append(result, []int{i, l, r})
				for nums[l] == nums[l-1] && l < r {
					l++
				}
				for nums[r] == nums[r-1] && l < r {
					r--
				}
				l++
				r--
			} else if nums[i]+nums[l]+nums[r] < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return result
}
