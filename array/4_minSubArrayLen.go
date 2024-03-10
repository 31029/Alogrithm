package array

func MinSubArrayLen(target int, nums []int) int {
	/* 暴力解法，超时 */
	minLen := 0
	for i := 0; i < len(nums); i++ {

		count := 0
		for j := i; j <= len(nums)-1; j++ {
			count += nums[j]
			if count >= target {
				if minLen == 0 {
					minLen = j + 1
				} else if j-i+1 < minLen {
					minLen = j - i + 1
				}
				break
			}
		}
		if minLen == 0 {
			return 0
		}
	}

	return minLen
}
