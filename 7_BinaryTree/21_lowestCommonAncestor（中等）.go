package binarytree

// LeetCode-236: 二叉树的最近公共祖先. [https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/description/]
/* 描述：
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，
满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

提示：
树中节点数目在范围 [2, 105] 内。
-109 <= Node.val <= 109
所有 Node.val 互不相同 。
p != q
p 和 q 均存在于给定的二叉树中。
*/
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 思路：常规递归思路。
	// 注意点：1、res变量的传递问题。
	res := &[]*TreeNode{}
	lowestCommonAncestorRecursion(root, p.Val, q.Val, res)
	return (*res)[0]
}

func lowestCommonAncestorRecursion(root *TreeNode, p, q int, res *[]*TreeNode) (bool, bool) {
	proot, qroot:= false, false
	if root == nil {
		return false, false
	}
	if root.Val == p {
		proot = true
	} else if root.Val == q {
		qroot = true
	}
	pL, qL := lowestCommonAncestorRecursion(root.Left, p, q, res)
	pR, qR := lowestCommonAncestorRecursion(root.Right, p, q, res)
	proot = proot || pL || pR
	qroot = qroot || qL || qR
	if proot && qroot {
		*res = append(*res, root)
	} 
	return proot, qroot

}