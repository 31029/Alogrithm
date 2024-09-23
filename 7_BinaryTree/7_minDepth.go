package binarytree

func min(a, b int) int {
	if a >= b && b != -1{
		return b
	}
	return a
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	
	mR := minDepth(root.Right)
	mL := minDepth(root.Left)
	
	depth := 1 + min(mR, mL)
	if root.Left == nil && root.Right != nil {
		return mR + 1
	} else if root.Right == nil && root.Left != nil {
		return mL + 1
	}

	return depth
}