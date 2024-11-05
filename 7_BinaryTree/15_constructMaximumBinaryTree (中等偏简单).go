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

func ConstructMaximumBinaryTree(nums []int) *TreeNode {
	index := findMaxIndex(nums)
	root := &TreeNode{
		Val: nums[index],
	}
	ConstructMaximumBinaryTreeRecursion(root, nums[:index], nums[index+1:])
	return root
}

func ConstructMaximumBinaryTreeRecursion(root *TreeNode, Lnums, Rnums []int){
	if len(Lnums) > 0 {
		Lindex := findMaxIndex(Lnums)
		root.Left = &TreeNode{
			Val: Lnums[Lindex],
		}
		ConstructMaximumBinaryTreeRecursion(root.Left, Lnums[:Lindex], Lnums[Lindex+1:])
	}
	if len(Rnums) > 0 {
		Rindex := findMaxIndex(Rnums)
		root.Right = &TreeNode{
			Val: Rnums[Rindex],
		}
		ConstructMaximumBinaryTreeRecursion(root.Right, Rnums[:Rindex], Rnums[Rindex+1:])
	}
}