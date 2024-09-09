package stack_queue

import utils "github.com/h31029/alogrithm/utils"

func TopKFrequent(nums []int, k int) []int {
	countMap := make(map[int]int)
	for _, v := range nums {
		countMap[v]++
	}
	freqList := []int{}
	countMap2 := make(map[int][]int)
	for k, v := range countMap {
		if len(countMap2[v]) == 0 || countMap2[v] == nil {
			freqList = append(freqList, v)
		}
		countMap2[v] = append(countMap2[v], k)
	}

	utils.BubbleSort(freqList)

	c := 0
	result := []int{}
	// freqList: 1c  2c  2c  4c  4c  5c  5c, k=3
	for i := len(freqList)-1; i >= 0 && c < k; i-- {
		freq := freqList[i]
		if len(countMap2[freq]) > k-c {
			result = append(result, countMap2[freq][:k-c]...)
			c += len(countMap2[freq][:k-c])
		} else {
			result = append(result, countMap2[freq]...)
			c += len(countMap2[freq])
		}
	}
	return result
}
