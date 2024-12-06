package hashmap

import utils "github.com/h31029/alogrithm/utils"

// leetcode-18：四数之和 [https://leetcode.cn/problems/4sum/description/]
/*给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 
找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）
	① 0 <= a, b, c, d < n
	② a、b、c 和 d 互不相同
	③ nums[a] + nums[b] + nums[c] + nums[d] == target
你可以按 任意顺序 返回答案 
*/
func FourSum(nums []int, target int) [][]int {
	// 同三数之和，可以推测出题人的目的，是将O(n^4)复杂度降至O(n^3)，关键的难点在于去重！！！
	// x4: 去重i和j、 right和left； 剪枝条件 target（其实不减枝也可以）
	result := [][]int{}
	utils.BubbleSort(nums)

	for i := 0; i < len(nums)-3; i++ {
		// 剪枝
		if nums[i] > target && nums[i] > 0 {
			break
		}

		// i 去重
		if nums[i] == nums[i-1] && i>0 {
			continue
		}

		// leetcode-15：三数之和
		for j := i+1; j < len(nums)-2; j++ {
			// j 去重， 注意j>i+1这个条件不能漏！！！
			if nums[j] == nums[j-1] && j > i+1{
				continue
			}

			right := len(nums)-1
			left := j+1
			for right > left {
				if nums[i] + nums[j] + nums[left] + nums[right] == target {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})

					// left 和 right去重
					for nums[right] == nums[right-1] && right>left{
						right--
					}
					for nums[left] == nums[left+1] && right>left{
						left++
					}
					right--
					left++
				} else if nums[i] + nums[j] + nums[left] + nums[right] > target {
					right -- 
				} else {
					left ++
				}	
			}

		}
	}
	return result
}