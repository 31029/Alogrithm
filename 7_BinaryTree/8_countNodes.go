package binarytree

func CountNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	nR := CountNodes(root.Right)
	nL := CountNodes(root.Left)

	return nR + nL + 1
}