package array


func RemoveElement(nums []int, val int) int {
	/* 标准快、慢指针方法 */
	slow := 0	//慢指针：指向更新 新数组下标的位置
	fast := 0	//快指针：寻找新数组的元素 ，新数组就是不含有目标元素的数组

	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast ++ 
	}

	return slow
}


func RemoveElement_myself(nums []int, val int) int {
	/* 自己的方法 */
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


/* 失败的快、慢指针方法 */
func swap(nums []int, left int, right int)  {
	temp := nums[left]
	nums[left] = nums[right]
	nums[right] = temp
}


func RemoveElement_failed(nums []int, val int) int {
	/* 失败的快、慢指针方法 */
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