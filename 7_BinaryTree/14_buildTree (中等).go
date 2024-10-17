package binarytree

// LeetCode-106: 从中序与后序遍历构建二叉树
// 描述：给定两个整数数组 inorder 和 postorder ，其中 inorder 是二叉树的中序遍历， postorder 是同一棵树的后序遍历，请你构造并返回这颗 二叉树 。
func buildTree(inorder []int, postorder []int) *TreeNode {
	//思路1： 通过中序后序遍历二叉树的特性求解，画图，注意中序遍历的局限性
	root := &TreeNode{
		Val: postorder[len(postorder)-1],
	}
	rootIndex := 0
	for k, v := range inorder {
		if v == postorder[len(postorder)-1] {
			rootIndex = k
		}
	}
	buildTreeRecursion(root, inorder[:rootIndex], inorder[rootIndex+1:], postorder[:rootIndex], postorder[rootIndex:len(postorder)-1])
	return root
}

// func buildTreeRecursion_Failed(root *TreeNode, LInorder, RInorder []int)  {
// 	if len(LInorder) == 0 && len(RInorder) == 0 {
// 		return
// 	}
// 	if len(LInorder) != 0 {
// 		root.Left = &TreeNode{
// 			Val: LInorder[len(LInorder)/2],
// 		}
// 		buildTreeRecursion(root.Left, LInorder[:len(LInorder)/2], LInorder[len(LInorder)/2 + 1:])
// 	}
// 	if len(RInorder) != 0 {
// 		root.Right = &TreeNode{
// 			Val: RInorder[len(RInorder)/2],
// 		}
// 		buildTreeRecursion(root.Right, RInorder[:len(RInorder)/2], RInorder[len(RInorder)/2 + 1:])
// 	}
// }

func buildTreeRecursion(root *TreeNode, LInorder, RInorder, LPostorder, RPostorder []int)  {
	if len(LInorder) == 0 && len(RInorder) == 0 {
		return
	}
	if len(LInorder) != 0 {
		root.Left = &TreeNode{
			Val: LPostorder[len(LPostorder)-1],
		}
		// find new root index
		nrI := 0 
		for k, v := range LInorder {
			if v == LPostorder[len(LPostorder)-1]{
				nrI = k
			}
		}
		buildTreeRecursion(root.Left, LInorder[:nrI], LInorder[nrI + 1:], LPostorder[:nrI], LPostorder[nrI:len(LPostorder)-1])
	}

	if len(RInorder) != 0 {
		root.Right = &TreeNode{
			Val: RPostorder[len(RPostorder)-1],
		}
		// find new root index
		nrI := 0 
		for k, v := range RInorder {
			if v == RPostorder[len(RPostorder)-1]{
				nrI = k
			}
		}
		buildTreeRecursion(root.Right, RInorder[:nrI], RInorder[nrI + 1:], RPostorder[:nrI], RPostorder[nrI:len(RPostorder)-1])
	}
}