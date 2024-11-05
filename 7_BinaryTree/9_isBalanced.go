package binarytree

func IsBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	bL := IsBalanced(root.Left)
	bR := IsBalanced(root.Right)


	dL := MaxDepth(root.Left)
	dR := MaxDepth(root.Right)

	return (dL - dR <=1 ) && bL && bR
}