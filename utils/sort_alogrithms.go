package utils

func BubbleSort(tgt []int) {
	// 优化过后的bubbleSort
	for i, change := 0, true; i < len(tgt) && change; i++ {
		change = false
		for j := 0; j < len(tgt) - 1 - i; j++ {
			if tgt[j] > tgt[j+1] {
				tmp := tgt[j]
				tgt[j] = tgt[j+1]
				tgt[j+1] = tmp
				change = true
			}
		}
	}
}

func partition(low, high int, nums []int) int {
	pivot := nums[low]
	for low < high {
		for nums[high] > pivot && high > low{
			high --
		}
		nums[low] = nums[high]

		for nums[low] < pivot && high > low{
			low ++
		}
		nums[high] = nums[low]
	}
	nums[low] = pivot
	return low
}

func QuickSort(low, high int, nums []int)  {
	if low < high {
		pivot := partition(low, high, nums)
		QuickSort(pivot+1, high, nums)
		QuickSort(0, low, nums)
	}
}