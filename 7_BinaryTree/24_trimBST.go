package binarytree

// LeetCode-669:修剪二叉搜索树 [https://leetcode.cn/problems/trim-a-binary-search-tree/description/]
/*
给你二叉搜索树的根节点 root ，同时给定最小边界low 和最大边界 high。
通过修剪二叉搜索树，使得所有节点的值在[low, high]中。
修剪树 不应该 改变保留在树中的元素的相对结构 (即，如果没有被移除，原有的父代子代关系都应当保留)。
可以证明，存在 唯一的答案 。
所以结果应当返回修剪好的二叉搜索树的新的根节点。注意，根节点可能会根据给定的边界发生改变。
*/

func TrimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	// 删除根节点时需要单独处理
	if root.Val > high || root.Val < low {
		dummyP := &TreeNode{
			Val: -1,
			Right: root,
		}
		trimBSTRecursion(root, dummyP, low, high)
		if dummyP.Right == nil {
			return nil
		}
		if dummyP.Right.Val > high || dummyP.Right.Val < low {
			return TrimBST(dummyP.Right, low, high)
		}
		return dummyP.Right
	}

	trimBSTRecursion(root, root, low, high)
    return root
}

func trimBSTRecursion(root, parent *TreeNode, low int, high int) {
	// 思路：画图分析问题
	// 注意点1️⃣：
    if root == nil {
		return
	}
	for root.Val > high || root.Val < low {
		if root.Val < parent.Val {
			if root.Right != nil {
				parent.Left = root.Right
				// 找到 root.right子树 最左边的叶子节点 node-x, node-x.Left = root.Left
				res := &[]*TreeNode{}
				inorderTraversal(root.Right, res)
				// 取 idx = 0
				node := (*res)[0]
				node.Left = root.Left
				root = parent.Left
			} else if root.Right == nil {
				parent.Left = root.Left
				root = parent.Left
			}
		} else if root.Val > parent.Val {
			if root.Left != nil {
				parent.Right = root.Left
				// 找到 root.Left子树 最右边的叶子节点 node-x, node-x.Right = root.Right
				res := &[]*TreeNode{}
				inorderTraversal(root.Left, res)
				// 取 idx = len(*res)-1
				node := (*res)[len(*res)-1]
				node.Right = root.Right
				root = parent.Right
			} else if root.Left == nil {
				parent.Right = root.Right
				root = parent.Right
			}
		}
	} 

	trimBSTRecursion(parent.Right, parent, low, high)
	trimBSTRecursion(parent.Left, parent, low, high)

}