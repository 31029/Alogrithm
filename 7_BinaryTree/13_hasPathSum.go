package binarytree

func GetAllPathSumRecursion(root *TreeNode, cur int, pSum *[]int)  {
	if root == nil {
		return
	}

	cur += root.Val
	if root.Left == nil && root.Right == nil {
		*pSum = append(*pSum, cur)
		return
	}

	GetAllPathSumRecursion(root.Left, cur, pSum)
	GetAllPathSumRecursion(root.Right, cur, pSum)
}

// leetcode - 112 - 路径总和：
// 描述：二叉树所有路径和中是否存在值为target
func HasPathSum(root *TreeNode, targetSum int) bool {
	// 思路1：通过递归求出所有路径和，再遍历是否存在target值
	pSum := []int{}
	GetAllPathSumRecursion(root, 0, &pSum)
	
	for _, pathSum := range pSum {
		if pathSum == targetSum {
			return true
		}
	}

	return false
}