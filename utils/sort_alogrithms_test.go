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
