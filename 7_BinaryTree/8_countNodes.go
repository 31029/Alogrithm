package binarytree

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	nR := countNodes(root.Right)
	nL := countNodes(root.Left)

	return nR + nL + 1
}