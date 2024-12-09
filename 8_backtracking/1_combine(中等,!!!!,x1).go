package backtracking

// leetcode-77: 组合 [https://leetcode.cn/problems/combinations/]
/*
	给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
	你可以按 任何顺序 返回答案。
*/
import "container/list"

var result [][]int
var cpath = list.New()

func combineRecursion(n, k int, startIdx int){
	if cpath.Len() == k {
		tmp := []int{}
		t := cpath.Front()
		for i := 0; i < cpath.Len(); i++ {
			if v, ok := t.Value.(int); ok {
				tmp = append(tmp, v)
				t = t.Next()
			}
		}
		result = append(result, tmp)
		return
	}

	for i := startIdx; i <= n; i++ {
		cpath.PushBack(i)
		combineRecursion(n, k, i+1)
		cpath.Remove(cpath.Back())
	}
}

func Combine(n int, k int) [][]int {
	// 思路：经典回溯解法，用递归来实现多层for循环嵌套
	// 但是时间和内存消耗都较高，考虑不用container包改用Slice模拟栈再实现
	defer func() {
		cpath.Init()
		result = [][]int{}
	}()
    combineRecursion(n, k, 1)
	return result
}