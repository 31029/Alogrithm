package binarytree

// leetcode-701: 二叉搜索树中的插入操作 [https://leetcode.cn/problems/insert-into-a-binary-search-tree/description/]
/* 描述：
给定二叉搜索树（BST）的根节点 root 和要插入树中的值 value ，将值插入二叉搜索树。 返回插入后二叉搜索树的根节点。
输入数据 保证 ，新值和原始二叉搜索树中的任意节点值都不同。
注意，可能存在多种有效的插入方式，只要树在插入后仍保持为二叉搜索树即可。 你可以返回 任意有效的结果 。
*/
func InsertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: val,
		}
	}

	if val > root.Val {
		insertIntoBSTRecursion(root.Right, root, val)
	} else if val < root.Val {
		insertIntoBSTRecursion(root.Left, root, val)
	}
	return root
}

func insertIntoBSTRecursion(root, parent *TreeNode, val int) {
	if root == nil {
		if val > parent.Val {
			parent.Right = &TreeNode{
				Val: val,
			}
		} else if val < parent.Val {
			parent.Left = &TreeNode{
				Val: val,
			}
		}
		return
	}

	if val > root.Val {
		insertIntoBSTRecursion(root.Right, root, val)
	} else if val < root.Val {
		insertIntoBSTRecursion(root.Left, root, val)
	}
}
