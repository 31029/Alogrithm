package binarytree

func findMaxIndex(nums []int) int {
	max := -1
	index := -1
	for k, v := range nums {
		if v > max {
			max = v
			index = k
		}
	}
	return index
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	index := findMaxIndex(nums)
	root := &TreeNode{
		Val: nums[index],
	}
	constructMaximumBinaryTreeRecursion(root, nums[:index], nums[index+1:])
	return root
}

func constructMaximumBinaryTreeRecursion(root *TreeNode, Lnums, Rnums []int){
	if len(Lnums) > 0 {
		Lindex := findMaxIndex(Lnums)
		root.Left = &TreeNode{
			Val: Lnums[Lindex],
		}
		constructMaximumBinaryTreeRecursion(root.Left, Lnums[:Lindex], Lnums[Lindex+1:])
	}
	if len(Rnums) > 0 {
		Rindex := findMaxIndex(Rnums)
		root.Right = &TreeNode{
			Val: Rnums[Rindex],
		}
		constructMaximumBinaryTreeRecursion(root.Right, Rnums[:Rindex], Rnums[Rindex+1:])
	}
}