package binarytree

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	bL := isBalanced(root.Left)
	bR := isBalanced(root.Right)


	dL := MaxDepth(root.Left)
	dR := MaxDepth(root.Right)

	return (dL - dR <=1 ) && bL && bR
}