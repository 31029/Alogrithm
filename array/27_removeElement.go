package array


func RemoveElement(nums []int, val int) int {
	count := 0
	moveIndex := 0

	for _, v := range nums {
		if v != val {
			count ++
		}
	}

	if count > 0 {
		for k, v := range nums {
			if v != val && moveIndex <= count{
				nums[moveIndex] = v	
				if moveIndex != k {
					nums[k] = val
				}
				moveIndex ++
			}
		}
	}

	return count
}


func swap(nums []int, left int, right int)  {
	temp := nums[left]
	nums[left] = nums[right]
	nums[right] = temp
}


func RemoveElement1(nums []int, val int) int {
	left := 0
	right := left + 1

	count := 0

	if len(nums) == 1{
		if nums[0] == val {
			return 0
		} else {
			return 1
		}
	}

	for _, v := range nums {
		if v == val {
			count ++
		}
	}
	
	for i := 0; i < count; i++ {
		
	}


	for left < len(nums) && right < len(nums){
		if nums[left] != val {
			left ++
			right = left + 1
		} else {
			for right < len(nums) {
				if nums[right] != val {
					swap(nums, left, right)
					left ++
					right = left + 1
					break
				} else if right==len(nums)-1 {
					return left
				} else {
					right ++
				}
			}
		}
	}

	return count
}