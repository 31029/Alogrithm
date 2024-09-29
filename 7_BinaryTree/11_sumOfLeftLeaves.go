package binarytree

func sumOfLeftLeaves(root *TreeNode) int {
	return sumOfLeftLeavesRecursion(root, false)
}

// leetcode - 404: 左子树之和
func sumOfLeftLeavesRecursion(root *TreeNode, isLeft bool) int {
	// 思路1：递归
	// 思路2： 层序遍历后按规律取值求和
	if root == nil {
		return 0
	}

	sL := sumOfLeftLeavesRecursion(root.Left, true)
	sR := sumOfLeftLeavesRecursion(root.Right, false)

	cur := 0
	if isLeft {
		cur = root.Val
	}
	return cur + sL + sR
}