package binarytree

// 迭代遍历快速解决
func invertTree(root *TreeNode) *TreeNode {
	// 方法： 迭代遍历快速解决
	// 个人觉得层序遍历更好理解。
	if root == nil {
		return nil
	}
	// 注意go的技巧，不需要加入temp临时变量。
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)

	return root
}