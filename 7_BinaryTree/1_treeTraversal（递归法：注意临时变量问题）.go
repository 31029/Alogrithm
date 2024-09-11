package binarytree

func preorderTraversal(root *TreeNode) []int {
	r := []int{}
	recursionPreorder(root, &r)
	return r
}

// 递归 ①： 确定参数和返回值
func recursionPreorder(root *TreeNode, nums *[]int) {
	// 递归 ②： 确定终止条件
	if root == nil {
		return 
	} 

	// 递归 ③： 确定单层递归逻辑。
	*nums = append(*nums, root.Val)
	recursionPreorder(root.Left, nums)
	recursionPreorder(root.Right, nums)
}

func recursionInorder(root *TreeNode, nums []int) {
	if root == nil {
		return 
	} 
	recursionInorder(root.Left, nums)
	// 此处append分配了新的内存空间，nums变为临时变量
	nums = append(nums, root.Val)
	recursionInorder(root.Right, nums)
}

func recursionPostorder(root *TreeNode, nums []int) {
	if root == nil {
		return 
	} 
	recursionPostorder(root.Left, nums)
	recursionPostorder(root.Right, nums)
	nums = append(nums, root.Val)
}