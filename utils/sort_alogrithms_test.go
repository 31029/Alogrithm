package utils

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	t_A := []int{3, 2, 1, 4, 5}

	// 冒泡排序
	BubbleSort(t_A)

	t.Logf("sorted array:%v", t_A)
}


func TestQuickSort(t *testing.T)  {
	nums := []int{5, 7, 2, 1, 9, 8, 10}
	QuickSort(0, len(nums)-1, nums)
	for _, v := range nums {
		println(v)
	}
}