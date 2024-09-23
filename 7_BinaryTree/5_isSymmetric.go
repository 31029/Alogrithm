package binarytree

// 树：判断树是否对称
func IsSymmetric(root *TreeNode) bool {
	// 思路1：反转树后判断前后两颗树是否相同。此处不能用前序遍历，不同的树前序遍历的结果相同。
	n1 := LevelOrderWithEmpty(root)
	invertTree(root)
	n2 := LevelOrderWithEmpty(root)

	for i:= 0; i<len(n1); i++ {
		if len(n1[i]) != len (n2[i]) {
			return false
		}
		for j := 0; j < len(n1[i]); j ++ {
			// 处理nil比较情况
			if n1[i][j] == nil && n2[i][j] != nil {
				return false
			} else if n1[i][j] != nil && n2[i][j] == nil {
				return false
			}

			// 处理一般比较情况
			if n1[i][j] != n2[i][j] {
				return false
			}
		}
	}
	return true
}