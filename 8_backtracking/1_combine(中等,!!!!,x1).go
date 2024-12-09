package backtracking

// leetcode-77: 组合 [https://leetcode.cn/problems/combinations/]
/*
	给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
	你可以按 任何顺序 返回答案。
*/
func combineRecursion(n, k int, startIdx int, cpath *[]int, result *[][]int){
	if len(*cpath) == k {
		temp := make([]int, k)
		copy(temp, *cpath)
		*result = append(*result, temp)
		return
	}

	for i := startIdx; i <= n; i++ {
		*cpath = append(*cpath, i)
		combineRecursion(n, k, i+1, cpath, result)
		if len(*cpath) >= 1 {
			*cpath = (*cpath)[:len(*cpath)-1]
		}
	}
}

func Combine(n int, k int) [][]int {
	// 思路：经典回溯解法，用递归来实现多层for循环嵌套
	// 但是时间和内存消耗都较高，考虑不用container包改用Slice模拟栈再实现
	result := [][]int{}
	cpath := []int{}
    combineRecursion(n, k, 1, &cpath, &result)
	return result
}