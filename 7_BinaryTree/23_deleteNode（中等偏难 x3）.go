package binarytree

// LeetCode-450. 删除二叉搜索树中的节点 [https://leetcode.cn/problems/delete-node-in-a-bst/description/]
/*
给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，并保证二叉搜索树的性质不变。
返回二叉搜索树（有可能被更新）的根节点的引用。注意可能有多个结果
*/
func DeleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	// 删除根节点时需要单独处理
	if root.Val == key {
		dummyP := &TreeNode{
			Val: -1,
			Right: root,
		}
		deleteNodeRecursion(root, dummyP, key)
		return dummyP.Right
	}
	deleteNodeRecursion(root, root, key)
    return root
}

func inorderTraversal(root *TreeNode, res *[]*TreeNode)  {
	if root == nil {
		return
	}
	inorderTraversal(root.Left, res)
	*res = append(*res, root)
	inorderTraversal(root.Right, res)
}

func deleteNodeRecursion(root, parent *TreeNode, key int) {
	// 思路：画图分析问题
	// 注意点1️⃣：该函数集合搜索和删除一体，要注意停止条件
	// 注意点2️⃣: 删除根节点时需要单独处理
    if root == nil {
		return
	}
	if root.Val == key {
		if root.Val < parent.Val {
			if root.Right != nil {
				parent.Left = root.Right
				// 找到 root.right子树 最左边的叶子节点 node-x, node-x.Left = root.Left
				res := &[]*TreeNode{}
				inorderTraversal(root.Right, res)
				// 取 idx = 0
				node := (*res)[0]
				node.Left = root.Left
			} else if root.Right == nil {
				parent.Left = root.Left
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
			} else if root.Left == nil {
				parent.Right = root.Right
			}
		}
		return
	} else if root.Val > key {
		deleteNodeRecursion(root.Left, root, key)
	} else {
		deleteNodeRecursion(root.Right, root, key)
	}

}