package dualPointer

// [1, 2, 3, 2, 5, 4]  k=2
// [1, 4, 3, 2, 5, 2]
// [1, 4, 3, 5, 2, 2]

// [1, 3, 5, 2, 2, 2]  k=2
// [1]  k=1
func RemoveElement(nums []int, val int) int {
	li := 0
	ri := len(nums)-1
	// 注意此处不能为 !=，nums长度为0时出错
	for li <= ri {
		if nums[li] == val {
			// 找到切片右侧第一个不为val的值的index
			for li < ri && nums[ri] == val {
				ri --
			}

			// 此处条件为必要条件，在这里失误了4、5次，
			// 一定要想清楚 li == ri 时的处理逻辑。
			if li == ri {
				return li
			}
			nums[li] = nums[ri]
			nums[ri] = val
		}
		li ++
	}

	return li
}