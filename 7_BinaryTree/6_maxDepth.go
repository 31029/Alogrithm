package binarytree

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func MaxDepth(root *TreeNode) int {
	// if root == nil {
	// 	return 0
	// }

	// mR := MaxDepth(root.Right) + 1
	// mL := MaxDepth(root.Left) + 1

	// return max(mR, mL)

	// 思路：后序遍历， 注意返回值
	// 思路2：层序遍历： https://blog.csdn.net/Mamba_Stan/article/details/132729568
	if (root == nil) {
        return 0
    }

    mL := MaxDepth(root.Left)
    mR := MaxDepth(root.Right)
    depth := 1 + max(mL, mR)
    return depth
}