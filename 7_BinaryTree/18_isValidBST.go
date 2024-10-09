package binarytree

// LeetCode-98: 验证二叉搜索树。[https://leetcode.cn/problems/validate-binary-search-tree/description/]
/* 描述：
给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
有效 二叉搜索树定义如下：
节点的左子树只包含 小于 当前节点的数。
节点的右子树只包含 大于 当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
*/
func IsValidBSTRecursion(root *TreeNode) bool {
	if root == nil {
		return false
	}
	if root.Left != nil && root.Right != nil {
		if root.Right.Val > root.Val && root.Left.Val < root.Val {
			return IsValidBSTRecursion(root.Left) && IsValidBSTRecursion(root.Right)
		} else {
			return false
		}
	}

	if root.Left != nil && root.Right == nil {
		if root.Left.Val < root.Val {
			return IsValidBSTRecursion(root.Left)
		} else {
			return false
		}
	}

	if root.Right != nil && root.Left == nil {
		if root.Right.Val > root.Val {
			return IsValidBSTRecursion(root.Right)
		} else {
			return false
		}
	}

	if root.Right == nil && root.Left == nil {
		return true
	}

	return false
}