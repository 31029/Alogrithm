package binarytree

func Min(a, b int) int {
	if a >= b && b != -1{
		return b
	}
	return a
}

func MinDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	
	mR := MinDepth(root.Right)
	mL := MinDepth(root.Left)
	
	depth := 1 + min(mR, mL)
	if root.Left == nil && root.Right != nil {
		return mR + 1
	} else if root.Right == nil && root.Left != nil {
		return mL + 1
	}

	return depth
}