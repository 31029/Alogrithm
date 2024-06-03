package hashmap

func FourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	// 注意题目中的隐藏条件如：只需要返回个数而不需要返回下标，各个数组长度一定
	old_map := make(map[int]int)
	n_list := [][]int{nums1, nums2, nums3, nums4}

	old_map[0] = 1
	for _, nums := range n_list {
		new_map := make(map[int]int)
		for _, num := range nums {
			for k, count := range old_map {
				new_map[num + k] += count
			}
		}
		old_map = new_map
	}

	return old_map[0]
}