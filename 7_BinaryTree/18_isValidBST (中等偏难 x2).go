package binarytree

// LeetCode-98: 验证二叉搜索树。[https://leetcode.cn/problems/validate-binary-search-tree/description/]
/* 描述：
给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
有效 二叉搜索树定义如下：
节点的左子树只包含 小于 当前节点的数。
节点的右子树只包含 大于 当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
*/
func isValidBST(root *TreeNode) bool {
	// 思路：多返回值递归模式，需要保证子树递归地满足是BST的同时还要满足 root.Val > maxL && root.Val < minR。
	if root == nil {
		return false
	}
	isValid, _, _ := IsValidBSTRecursion(root)
	return isValid
}

func findMax(nums []int) int {
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}

func findMin(nums []int) int {
	min := nums[0]
	for _, v := range nums {
		if v < min {
			min = v
		}
	}
	return min
}

func IsValidBSTRecursion(root *TreeNode) (bool, int, int) {
	isValid := false
	min := root.Val
	max := root.Val
	if root.Left != nil && root.Right != nil {
		if root.Right.Val > root.Val && root.Left.Val < root.Val {
			isValidL, minL, maxL := IsValidBSTRecursion(root.Left)
			isValidR, minR, maxR := IsValidBSTRecursion(root.Right)

			min = findMin([]int{min, minL, minR})
			max = findMax([]int{max, maxL, maxR})
			isValid = isValidL && isValidR && root.Val > maxL && root.Val < minR
			return isValid, min, max
		}
	}

	if root.Left != nil && root.Right == nil {
		if root.Left.Val < root.Val {
			isValidL, minL, maxL := IsValidBSTRecursion(root.Left)

			min = findMin([]int{min, minL})
			max = findMax([]int{max, maxL})
			isValid = isValidL && root.Val > maxL
			return isValid, min, max
		}
	}

	if root.Right != nil && root.Left == nil {
		if root.Right.Val > root.Val {
			isValidR, minR, maxR := IsValidBSTRecursion(root.Right)

			min = findMin([]int{min, minR})
			max = findMax([]int{max, maxR})
			isValid = isValidR && root.Val < minR
			return isValid, min, max
		}
	}

	if root.Right == nil && root.Left == nil {
		return true, min, max
	}

	return false, min, max
}
