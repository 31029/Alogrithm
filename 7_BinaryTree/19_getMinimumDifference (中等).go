package binarytree

import "math"
// LeetCode-530: 二叉搜索树的最小绝对差. [https://leetcode.cn/problems/minimum-absolute-difference-in-bst/description/]
/* 描述：
给你一个二叉搜索树的根节点 root ，返回 树中任意两不同节点值之间的最小差值 。
差值是一个正数，其数值等于两值之差的绝对值。
*/
func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}


func getMinimumDifferenceRecursion(root *TreeNode) (int, int, int) {
	// 思路同18，但是时间和空间复杂度都很低，需要找寻更好的办法。
	minGap := math.MaxInt
	min := root.Val
	max := root.Val
	if root.Left != nil && root.Right != nil {
		minGapL, minL, maxL := getMinimumDifferenceRecursion(root.Left)
		minGapR, minR, maxR := getMinimumDifferenceRecursion(root.Right)

		min = findMin([]int{min, minL, minR})
		max = findMax([]int{max, maxL, maxR})
		minGap = findMin([]int{abs(root.Val-maxL), abs(root.Val-minR), minGapL, minGapR})
		return minGap, min, max
	}

	if root.Left != nil && root.Right == nil {
		minGapL, minL, maxL := getMinimumDifferenceRecursion(root.Left)

		min = findMin([]int{min, minL})
		max = findMax([]int{max, maxL})
		minGap = findMin([]int{abs(root.Val-maxL), minGapL})
		return minGap, min, max
	}

	if root.Right != nil && root.Left == nil {
		minGapR, minR, maxR := getMinimumDifferenceRecursion(root.Right)

		min = findMin([]int{min, minR})
		max = findMax([]int{max, maxR})
		minGap = findMin([]int{abs(root.Val-minR), minGapR})
		return minGap, min, max
	}

	if root.Right == nil && root.Left == nil {
		return minGap, min, max
	}

	return minGap, min, max
}

func getMinimumDifference(root *TreeNode) int {
    minGap, _, _ := getMinimumDifferenceRecursion(root)
    return minGap
}