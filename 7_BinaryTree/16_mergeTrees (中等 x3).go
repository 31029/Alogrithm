package binarytree

// LeetCode-617: 合并二叉树。【https://leetcode.cn/problems/merge-two-binary-trees/】
/* 描述：
给你两棵二叉树： root1 和 root2 。
想象一下，当你将其中一棵覆盖到另一棵之上时，两棵树上的一些节点将会重叠（而另一些不会）。
你需要将这两棵树合并成一棵新二叉树。合并的规则是：如果两个节点重叠，那么将这两个节点的值相加作为合并后节点的新值；否则，不为 null 的节点将直接作为新二叉树的节点。
返回合并后的二叉树。
*/
func MergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	root := &TreeNode{}
	mergeTreesRecursion(root, root1, root2)
	return root
}

// func mergeTreesRecursion(root *TreeNode, root1 *TreeNode, root2 *TreeNode){
// 	if root1 == nil && root2 == nil {
// 		return
// 	} else if root1 == nil {
// 		root.Val = root2.Val
// 		mergeTreesRecursion(root.Left, nil, root2.Left)
// 		mergeTreesRecursion(root.Right, nil, root2.Right)
// 	} else if root2 == nil {
// 		root.Val = root1.Val
// 		mergeTreesRecursion(root.Left, root1.Left, nil)
// 		mergeTreesRecursion(root.Right, root1.Right, nil)
// 	} else {
// 		root.Val = root1.Val + root2.Val
// 		mergeTreesRecursion(root.Left, root1.Left, root2.Left)
// 		mergeTreesRecursion(root.Right, root1.Right, root2.Right)
// 	}
// }

func mergeTreesRecursion(root *TreeNode, root1 *TreeNode, root2 *TreeNode){
	// 思路：递归构建。 需要反思为啥这么复杂，还错了几次
	if root1 == nil && root2 == nil {
		return
	} else if root1 == nil {
		root.Val = root2.Val
		if root2.Left != nil{
			root.Left = &TreeNode{}
			mergeTreesRecursion(root.Left, nil, root2.Left)
		}
		if root2.Right != nil {
			root.Right = &TreeNode{}
			mergeTreesRecursion(root.Right, nil, root2.Right)
		}
	} else if root2 == nil {
		root.Val = root1.Val
		if root1.Left != nil{
			root.Left = &TreeNode{}
			mergeTreesRecursion(root.Left, nil, root1.Left)
		}
		if root1.Right != nil {
			root.Right = &TreeNode{}
			mergeTreesRecursion(root.Right, nil, root1.Right)
		}
	} else {
		root.Val = root1.Val + root2.Val
		if root1.Left != nil || root2.Left != nil {
			root.Left = &TreeNode{}
			mergeTreesRecursion(root.Left, root1.Left, root2.Left)
		}
		if root1.Right != nil || root2.Right != nil {
			root.Right = &TreeNode{}
			mergeTreesRecursion(root.Right, root1.Right, root2.Right)
		}
	}
}