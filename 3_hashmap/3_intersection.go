package hashmap


func Intersection(nums1 []int, nums2 []int) []int {
	i_map := make(map[int]int)
	r_map := make(map[int]int) 
	result := []int{}

	for _, v := range nums1 {
		i_map[v] ++
	}

	for _, v := range nums2 {
		if i_map[v] > 0 {
			r_map[v] ++
		}
	}

	for k, v := range r_map {
		if v > 0 {
			result = append(result, k)
		}
	}

	return result
}